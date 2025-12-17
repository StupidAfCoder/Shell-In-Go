package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

var allCommands = []string{"echo", "exit", "type"}
var PATH = os.Getenv("PATH")

func runCommand(path string, args []string) error {

	if args == nil {
		cmd := exec.Command(path)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		return cmd.Run()
	}

	cmd := exec.Command(path, args[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func checkIfExecutableExists(exec string) (string, bool) {
	pathSlice := strings.Split(PATH, ":")
	for _, path := range pathSlice {
		entries, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.Name() == exec {
				filePath := filepath.Join(path, exec)
				info, err := os.Stat(filePath)
				if err != nil {
					continue
				}
				isExecutable := info.Mode()&0111 != 0
				if isExecutable {
					return filePath, true
				}
			}
		}
	}
	return "", false
}

func checkCommand(command string) (string, string) {
	cmd, param, found := strings.Cut(command, " ")
	if !found {
		return command, ""
	}
	switch cmd {
	case "echo":
		return cmd, param
	case "type":
		return cmd, param
	default:
		return cmd, param
	}
}

func executeCommand(command string, param string) {
	switch command {
	case "echo":
		fmt.Println(param)
	case "type":
		actualCommand := slices.Contains(allCommands, param)
		if actualCommand {
			fmt.Printf("%s is a shell builtin\n", param)
			break
		}
		foundPath, directoryCommand := checkIfExecutableExists(param)
		if directoryCommand {
			fmt.Printf("%s is %s\n", param, foundPath)
			break
		}
		fmt.Printf("%s: not found\n", param)
	default:
		foundPath, directoryCommand := checkIfExecutableExists(command)
		if directoryCommand {
			param = strings.TrimSpace(param)
			if len(param) == 0 {
				err := runCommand(foundPath, nil)
				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				splitParam := strings.Split(param, " ")
				err := runCommand(foundPath, splitParam)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		} else {
			fmt.Println(command[:] + ": command not found" + foundPath)
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
