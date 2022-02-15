package main

import (
	"fmt"
	"os"
	"tadoo/list"
)

const (
	PROGRAM_NAME      = "Tadoo"
	PROGRAM_VERSION   = "1.0"
	PROGRAM_COPYRIGHT = "Tekronet 2022"
)

func main() {
	//checking if command is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid command. Type 'tadoo --help' for more information.")
	} else {
		switch os.Args[1] {
		//creating list
		case "create":
			if len(os.Args) > 2 {
				list.NewList(os.Args[2])
			} else {
				fmt.Println("Please provide a list name to create.")
			}
		//deleting list
		case "delete":
			if len(os.Args) > 2 {
				list.DeleteList(os.Args[2])
			} else {
				fmt.Println("Please provide a list name to delete.")
			}
		//selecting list
		case "select":
			if len(os.Args) > 2 {
				list.SelectList(os.Args[2])
			} else {
				fmt.Println("Please provide a list name to select.")
			}
		case "current":
			list.GetSelectedList()
		}
	}
}
