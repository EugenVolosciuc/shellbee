package commands

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/EugenVolosciuc/shellbee/internal/storage"
)

func Save(args ...string) {
	key := args[0]
	command := args[1]

	err := storage.WriteKey(key, command)

	if err != nil {
		fmt.Printf("Could not write the provided command alias, error: %v", err)
		return
	}

	fmt.Println("Command alias saved.")
}

func Run(args ...string) {
	key := args[0]

	command, err := storage.ReadKey(key)

	if err != nil {
		fmt.Printf("Could not read the provided command alias, error: %v", err)
		return
	}

	cmd := exec.Command("sh", "-c", command)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Printf("Could not run the command of the selected command alias, error: %v", err)
		return
	}

	fmt.Print(string(stdout))
}

func List(args ...string) {
	keys, err := storage.ListKeys()

	if err != nil {
		fmt.Printf("Could not list command aliases, error: %v", err)
		return
	}

	displayedList := "Total aliases: " + strconv.Itoa(len(keys)) + "\n"

	for _, key := range keys {
		displayedList += "  - " + "\"" + key + "\""
	}

	fmt.Printf("%s\n", displayedList)
}

func Search(args ...string) {
	fmt.Println("Find handler")
}

func Delete(args ...string) {
	key := args[0]

	if err := storage.DeleteKey(key); err != nil {
		fmt.Printf("Could not delete the provided command, error: %v", err)
		return
	}

	fmt.Println("Command alias deleted.")
}

func Help(args ...string) {
	helpText := `
shellbee — Your busy little helper in the shell 🐝

Usage:
  shb <command> [arguments]

Commands:
  save (s)    <name> <command>   Save a new command alias
  run  (r)    <name>             Run a saved command
  list (l)                       List all saved command aliases
  find (f)    <query>            Search saved commands by a partial name or content
  delete (d)  <name>             Delete a saved command alias
  help (h)                       Show this help message

Examples:
  shb save greet "echo 'Hello World'"
  shb run greet
  shb list
  shb find hello
  shb delete greet

More info: https://github.com/EugenVolosciuc/shellbee`
	fmt.Println(helpText)
}
