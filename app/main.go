package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
		switch {
		case command[0:4] == "echo":
			fmt.Println(command2[4:])
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
