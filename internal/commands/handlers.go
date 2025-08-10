package commands

import "fmt"

func Save(args ...string) {
	fmt.Println("Save handler")
}

func Run(args ...string) {
	fmt.Println("Run handler")
}

func List(args ...string) {
	fmt.Println("List handler")
}

func Find(args ...string) {
	fmt.Println("Find handler")
}

func Delete(args ...string) {
	fmt.Println("Delete handler")
}

func Help(args ...string) {
	fmt.Println("Help handler")
}
