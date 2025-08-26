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
var Completed_Todo_list = make([]TODO, 0)

func main() {
	fmt.Println("Hello world")
}

func addTodoItem(title string, duration time.Duration, status string, completed bool) []TODO {
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
	for _, TodoItem := range TODO_LIST {
		fmt.Printf("You have \"%v\" on your list. it is %v. you've set %v \n", TodoItem.title, TodoItem.status, TodoItem.duration)
	}
	return TODO_LIST
}

func getTodoItem(id int) TODO {
	if id < len(TODO_LIST) {
		return TODO_LIST[id]
	}

	return TODO{}
}

func moveTodoItemToCompleted(id int) []TODO {
	if id < len(TODO_LIST) {
		// remove from
		TODO_LIST = append(TODO_LIST[:id], TODO_LIST[id+1])
		Completed_Todo_list = append(Completed_Todo_list, TODO_LIST[id])
	}

	return Completed_Todo_list
}
