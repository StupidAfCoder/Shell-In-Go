package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(command[:len(command)-1] + ": command not found")
}
