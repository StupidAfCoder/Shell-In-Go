package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var allCommands = []string{"echo", "exit", "type"}

func checkCommand(command string) (string, string) {
	cmd, param, found := strings.Cut(command, " ")
	if !found {
		return command[:len(command)-1] + ": command not found", ""
	}
	switch cmd {
	case "echo":
		return cmd, param
	case "type":
		return cmd, param
	default:
		return command[:len(command)-1] + ": command not found", ""
	}
}

func executeCommand(command string, param string) {
	if param == "" {
		fmt.Println(command)
	}
	switch command {
	case "echo":
		fmt.Print(param)
	case "type":
		actualCommand := false
		for _, v := range allCommands {
			if param == v {
				actualCommand = true
				break
			}
		}
		if actualCommand {
			fmt.Printf("%s is a shell builtin\n", param)
		} else {
			fmt.Printf("%s: not found\n", param)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		}
		command2 := strings.TrimSpace(command)
		commandExit := "exit"
		if command2 == commandExit {
			break
		}
		checkedCommand, param := checkCommand(command2)
		executeCommand(checkedCommand, param)
	}
}
