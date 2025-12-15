package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		}
		command = strings.TrimSpace(command)
		commandExit := "exit"
		if command == commandExit {
			break
		}
		fmt.Println(command[:len(command)-1] + ": command not found")
	}
}
