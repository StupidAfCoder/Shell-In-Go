package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var allCommands = []string{"echo", "exit", "type"}
var PATH = "/usr/bin:/usr/local/bin:" + os.Getenv("PATH")

func checkCommand(command string) (string, string) {
	cmd, param, found := strings.Cut(command, " ")
	if !found {
		return command[:] + ": command not found", ""
	}
	switch cmd {
	case "echo":
		return cmd, param
	case "type":
		return cmd, param
	default:
		return command[:] + ": command not found", ""
	}
}

func executeCommand(command string, param string) {
	if param == "" {
		fmt.Println(command)
	}
	switch command {
	case "echo":
		fmt.Println(param)
	case "type":
		actualCommand := slices.Contains(allCommands, param)
		if actualCommand {
			fmt.Printf("%s is a shell builtin\n", param)
			break
		}
		directoryCommand := false
		foundPath := ""
		pathSlice := strings.Split(PATH, ":")
		for _, path := range pathSlice {
			entries, err := os.ReadDir(path)
			if err != nil {
				log.Fatal(err.Error())
			}
			for _, entry := range entries {
				if entry.Name() == param {
					info, err := os.Stat(path)
					if err != nil {
						log.Fatal(err.Error())
					}
					isExecutable := info.Mode()&0111 != 0
					if isExecutable {
						foundPath = path
						directoryCommand = true
					}
				}
			}
		}
		if directoryCommand {
			fmt.Printf("%s is %s\n", param, foundPath)
			break
		}
		fmt.Printf("%s: not found\n", param)
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
