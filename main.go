package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type TODO struct {
	todoID   string
	todoName string
}

var todoList []TODO

var todo TODO

func main() {
	welcome()
	menu()
}

func welcome() {
	fmt.Println("******* Welcome to my ToDo CLI App *******")
}

func menu() {
	newline(1)
	fmt.Println("Enter operation:")
	fmt.Println("1 Create new ToDo")
	fmt.Println("2 List ToDo")
	fmt.Println("3 Edit/Update ToDo")
	fmt.Println("4 Delete a ToDo")
	fmt.Println("0 to exit")

	newline(1)
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: Invalid input format")
		menu()
	}

	switch input {
	case 0:
		exit()
	case 1:
		createTodo()
	case 2:
		listTodo()
	case 3:
		updateTodo()
	case 4:
		deleteTodo()
	default:
		fmt.Println("Enter a valid operation")
		menu()
	}
}

func newline(noOfLines int) {

	for i := 0; i < noOfLines; i++ {
		fmt.Println("\n")
	}
}

func listTodo() {
	newline(1)

	if todoList == nil {
		fmt.Println("There are no TODOs. Create one below..")
	} else {
		fmt.Println("Here's a list of all your TODOs...")
		fmt.Println("*todo-ID*\t*to-do*")
	}
	newline(1)
	for _, v := range todoList {
		fmt.Println(v.todoID, "\t\t", v.todoName)
		newline(1)
	}
	menu()
}

func createTodo() {
	var todoTitle string
	newline(1)
	fmt.Println("Enter todo title:")

	_, err := fmt.Scan(&todoTitle)
	if err != nil {
		fmt.Println("Enter a valid todo title:")
	}

	todo.todoName = todoTitle
	todo.todoID = strconv.Itoa(rand.Intn(10000))

	todoList = append(todoList, todo)
	newline(1)
	fmt.Println("**Todo item added**")
	menu()
}

func updateTodo() {
	var newName string
	var update string

	fmt.Println("Enter the todo ID to edit:")
	_, err := fmt.Scan(&update)
	if err != nil {
		fmt.Println("Error: Invalid ID")
	}
	newline(1)

	for i, v := range todoList {
		if update == v.todoID {
			todoList = append(todoList[:i], todoList[i+1:]...)
			var newTodo TODO

			fmt.Printf("Todo found! Your todo is: %v\n", v.todoName)
			newline(1)
			fmt.Println("Enter new todo name:")

			_, err := fmt.Scan(&newName)
			if err != nil {
				fmt.Println("error!")
			}
			newline(1)

			newTodo.todoID = update
			newTodo.todoName = newName

			todoList = append(todoList, newTodo)
			fmt.Printf("New todo name '%v' saved\n", newName)
			menu()
		}
	}

	fmt.Println("No Todo found with that ID!")
	menu()
}

func deleteTodo() {
	var deleteInput string

	fmt.Println("Enter the Todo ID to delete:")
	_, err := fmt.Scan(&deleteInput)
	if err != nil {
		fmt.Println("Error: Invalid input")
	}
	newline(1)

	for i, v := range todoList {
		if deleteInput == v.todoID {
			todoList = append(todoList[:i], todoList[i+1:]...)
			fmt.Println("Todo item deleted")
			menu()
		}
	}
	fmt.Println("No Todo found with that ID.")
	menu()
}

func exit() {
	fmt.Println("** Goodbye! **")
	os.Exit(0)
}
