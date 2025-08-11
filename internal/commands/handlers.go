package commands

import (
	"fmt"

	"github.com/EugenVolosciuc/shellbee/internal/storage"
)

func Save(args ...string) error {
	key := args[0]
	command := args[1]

	err := storage.WriteKey(key, command)

	if err != nil {
		return err
	}

	fmt.Println("Command saved.")
	return nil
}

func Run(args ...string) error {
	fmt.Println("Run handler")
	return nil
}

func List(args ...string) error {
	fmt.Println("List handler")
	return nil
}

func Find(args ...string) error {
	fmt.Println("Find handler")
	return nil
}

func Delete(args ...string) error {
	fmt.Println("Delete handler")
	return nil
}

func Help(args ...string) error {
	fmt.Println("Help handler")
	return nil
}
