package main

import (
	"fmt"
	"os"

	"github.com/ContinuumLLC/GO/Golden-Crown/kingdom"
	"github.com/ContinuumLLC/GO/Golden-Crown/messages"
	"github.com/ContinuumLLC/GO/Golden-Crown/processor"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "" {
		fmt.Println("No input file path is given")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Failed to open file")
		return
	}

	inputs, err := processor.ProcessInput(file)
	if err != nil {
		fmt.Println("Failed to process inputs from file")
	}

	kingdom.Setup()

	fmt.Println(messages.ProcessMessages(inputs))

}
