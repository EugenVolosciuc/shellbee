package commands

import (
	"fmt"
	"os/exec"

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
	key := args[0]

	command, err := storage.ReadKey(key)

	if err != nil {
		return err
	}

	cmd := exec.Command(command)

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("error running the '%s' command: %w", key, err)
	}

	return nil
}

func List(args ...string) error {
	fmt.Println("List handler")
	return nil
}

func Search(args ...string) error {
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
