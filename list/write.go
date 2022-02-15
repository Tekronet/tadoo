package list

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func NewList(name string) {
	f, err := os.Create("./data/lists/" + name)
	if err != nil {
		fmt.Println("List already exists or tadoo catalog is broken")
	}
	defer f.Close()
}
func DeleteList(name string) {
	//checking if file is not existing
	if _, err := os.Stat("./data/lists/" + name); errors.Is(err, os.ErrNotExist) {
		fmt.Println("List " + name + " does not exist!")
		os.Exit(0)
	}
	fmt.Print("Do you really want to delete it? This action cannot be undone. [y/n]: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	input = strings.TrimSuffix(input, "\n")

	switch input {
	case "y":
		os.Remove("./data/lists/" + name)
		fmt.Println("List deleted!")
	default:
		os.Exit(0)
		fmt.Println("List not deleted.")
	}
}
