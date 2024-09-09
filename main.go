package main

import (
	"fmt"
)

var todo = []string{}

func commands() {
	fmt.Println("Commands Available:")
	fmt.Println("Help: Show The Commands Available")
	fmt.Println("List: Show The To-Do List")
	fmt.Println("Add: Add a To-Do to The List")
	fmt.Println("Edit: Edit a To-Do on The List")
	fmt.Println("Delete: Delete a To-Do on The List")
	fmt.Println("More: More Projects From Bagi763")
}

func add() {
	var todoName string
	fmt.Print("\nName of The To-Do: ")
	fmt.Scan(&todoName)
	todo = append(todo, todoName)
	fmt.Printf("\nTo-Do %v is Now on The List\n\n", todoName)
}

func edit() {
	var todoIndex uint
	var todoName string
	fmt.Println(todo)
	fmt.Print("\nIndex of The To-Do: ")
	fmt.Scan(&todoIndex)
	fmt.Print("\nNew Name of The To-Do: ")
	fmt.Scan(&todoName)
	todo = append(todo[:todoIndex], todoName)
	todoIndex += 1
	fmt.Print("\nTo-Do Edited\n")
}

func delete()  {
	var todoIndex uint
	fmt.Println(todo)
	fmt.Print("\nIndex of The To-Do: ")
	fmt.Scan(&todoIndex)
	todo = append(todo[:todoIndex], todo[todoIndex+1:]... )
	fmt.Println("To-Do Deleted")
}

func main() {
	var command string

	fmt.Print("--------------------------------------------------------------------\n\n")
	fmt.Print("Welcome to Go To-Do APP")
	fmt.Print("\n\n--------------------------------------------------------------------\n\n")
	commands()
	for {
		fmt.Print("> ")
		fmt.Scan(&command)

		switch command {
			case "help", "Help":
				commands()
			case "list", "List":
				if len(todo) == 0{
					fmt.Println("The List is Empty")
				}
				//* i tried using else but gave me erros, so i hard coded it
				if len(todo) != 0{
					fmt.Printf("\n %v \n\n", todo)
				}
				
			case "add", "Add":
				add()
			case "edit", "Edit":
				edit()
			case "delete", "Delete":
				delete()
			case "more", "More":
				fmt.Print("--------------------------------------------------------------------\n\n")
				fmt.Println("Github: https://github.com/Bagi763")
				fmt.Print("Linktree: https://linktr.ee/Bagi763")
				fmt.Print("\n\n--------------------------------------------------------------------\n\n")
			default:
				fmt.Println("Error")
	}
	
	}
}