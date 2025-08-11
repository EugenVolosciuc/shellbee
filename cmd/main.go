package main

import (
	"fmt"
	"os"

	"github.com/EugenVolosciuc/shellbee/internal/commands"
	"github.com/EugenVolosciuc/shellbee/internal/storage"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please provide a valid command")
		return
	}

	action, ok := commands.ActionsMap[args[0]]

	if !ok {
		action, ok = commands.ActionsMap[string([]byte(args[0]))]

		if !ok {
			fmt.Println("Invalid command provided")
			return
		}
	}

	err := storage.CheckAndCreateStorageFile()

	if err != nil {
		fmt.Printf("Could not find storage file, error: %v\n", err)
		return
	}

	handler, ok := commands.ActionHandlersMap[action]

	if !ok {
		fmt.Printf("'%v' action not handled!\n", action)
		return
	}

	handler(args[1:]...)
}
