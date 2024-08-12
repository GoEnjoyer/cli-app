package main

import (
	"cli-app/ezcli"
	"cli-app/task"
	"fmt"
)

// Create handler
func createHandler(data ezcli.CmdData, err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	taskList.Add(task.New(data["-name"], data["-category"], false))
	fmt.Println("Task is succefully added!")
}

// Delete handler
func deleteHandler(data ezcli.CmdData, err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = taskList.Remove(data["-id"])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Task is succesfully removed!")
}

// Complete handler
func completeHandler(data ezcli.CmdData, err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = taskList.ChangeTaskStatus(data["-id"], true)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Task is succesfully completed!")
}

// Uncomplete handler
func uncompleteHandler(data ezcli.CmdData, err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = taskList.ChangeTaskStatus(data["-id"], false)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Task is succesfully uncompleted!")
}

// List handler
func listHandler(data ezcli.CmdData, err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = taskList.ListTasks()

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
