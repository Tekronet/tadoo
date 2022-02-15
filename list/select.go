package list

import (
	"errors"
	"fmt"
	"os"
)

func SelectList(name string) {
	file, err := os.Open("./data/tadoo/selected")
	if err != nil {
		fmt.Println("ERROR: Cannot find data files. Make sure tadoo is in the right directory and data folder is not modified.")
		os.Exit(0)
	}
	file.Close()

	if _, err := os.Stat("./data/lists/" + name); errors.Is(err, os.ErrNotExist) {
		fmt.Println("List " + name + " does not exist!")
		os.Exit(0)
	}

	data := []byte(name)
	err2 := os.WriteFile("./data/tadoo/selected", data, 0644)
	if err2 != nil {
		fmt.Println("ERROR: Cannot write to data file. Make sue tadoo is in right directory and data folder is not modified.")
		os.Exit(0)
	}
	fmt.Println("Selected list: " + name)
}

func GetSelectedList() {
	file, err := os.ReadFile("./data/tadoo/selected")
	if err != nil {
		fmt.Println("ERROR: Cannot read data files. Make sure tadoo is in the right directory and data folder is not modified.")
	}

	content := string(file)
	if content == "" {
		fmt.Println("No list selected")
		os.Exit(0)
	}
	fmt.Println("Selected list: " + content)
}
