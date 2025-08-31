package main

import (
	"fmt"
	"time"
)

// BE ABLE TO SET A TODO ITEM
// BE ABLE TO ADD A TIMER
// BE ABLE TO LIST ALL YOUR TODO ITEM
// BE ABLE TO MOVE TODO ITEMS TO INPRGRSEE

type TODO struct {
	title     string
	duration  time.Duration
	status    string
	completed bool
}

var TODO_LIST = make([]TODO, 0)

func main() {
	fmt.Println("Hello world")

	// CREATE A TODO
	TODO_LIST = addTodoItem("Wash my clothes", TODO_LIST, (30 * time.Minute), "New", false)
	TODO_LIST = addTodoItem("Have my bath", TODO_LIST, (30 * time.Minute), "New", false)
	TODO_LIST = addTodoItem("Teach the scriptures", TODO_LIST, (30 * time.Minute), "New", false)
	TODO_LIST = addTodoItem("Study the bible", TODO_LIST, (30 * time.Minute), "New", false)
	TODO_LIST = addTodoItem("Code a game", TODO_LIST, (30 * time.Minute), "New", false)
	TODO_LIST = addTodoItem("CODE IN GOOO", TODO_LIST, (30 * time.Minute), "New", false)

	// MOVE WASH MY CLOTHES TO COMPLETION
	TODO_LIST = moveTodoItemToCompleted(0)
	TODO_LIST = moveTodoItemToCompleted(1)

	// GET ALL ITEMS
	getAllTodoItems()

	// SEE ALL COMPLETED ITEMS
	getAllCompletedItems()

	// SEE ITEM STATUS
	getItemStatus(0)
}

func addTodoItem(title string, TODO_LIST []TODO, duration time.Duration, status string, completed bool) []TODO {
	var TodoItem = TODO{
		title:     title,
		duration:  duration,
		status:    status,
		completed: completed,
	}

	TODO_LIST = append(TODO_LIST, TodoItem)
	fmt.Println("YOUR TODO ITEM HAS BEEN CREATED THANK YOU!")

	return TODO_LIST
}

func getAllTodoItems() []TODO {
	// var result = []string{}
	fmt.Println("See your Todos for the day bellow")
	for index, TodoItem := range TODO_LIST {
		fmt.Printf("%d.) You have \"%v\" on your list. it is %v. \n", index+1, TodoItem.title, TodoItem.status)
	}
	fmt.Print("\n\n\n")
	return TODO_LIST
}

func getTodoItem(id int) TODO {
	if id < len(TODO_LIST) {
		return TODO_LIST[id]
	}
	fmt.Print("\n\n\n")
	return TODO{}
}

func moveTodoItemToCompleted(id int) []TODO {
	if id < len(TODO_LIST) {
		fmt.Printf("You have completed \"%v\" \n", TODO_LIST[id])
		// remove from
		TODO_LIST[id].completed = true
		TODO_LIST[id].status = "Completed"
	}
	fmt.Print("\n\n\n")
	return TODO_LIST
}

func moveTodoItemToInProgress(id int) {
	if id < len(TODO_LIST) {
		TODO_LIST[id].status = "In progress"
	}
}

func getAllCompletedItems() {
	// var result = []string{}
	fmt.Println("See your completed Todos for the day bellow")
	for _, TodoItem := range TODO_LIST {
		if TodoItem.status == "Completed" {
			fmt.Printf("You have \"%v\" on your list. it is %v. \n", TodoItem.title, TodoItem.status)
		}
	}
	fmt.Print("\n")
}

func startATodoItem(id int) []TODO {
	if id < len(TODO_LIST) {
		TODO_LIST[id].status = "In progress"
	}
	return TODO_LIST
}

func getItemStatus(id int) {
	if id < len(TODO_LIST) {

		fmt.Printf("%v", TODO_LIST[id].status)
	}
	fmt.Println("Not Found")
}
