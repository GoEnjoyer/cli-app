package main

import (
	"cli-app/ezcli"
	"cli-app/task"
	"fmt"
	"os"
)

var taskList task.TaskList

func main() {
	const jsonFileName = "todo.json"

	//Load tasks from jsonFile
	taskList.Load(jsonFileName)

	//Flags
	nameFlag := ezcli.NewFlag("-name", true, "task's name")
	categoryFlag := ezcli.NewFlag("-category", true, "task's category")
	idFlag := ezcli.NewFlag("-id", true, "task's id")

	//Set create Command
	ezcli.SetCommand("create",
		"create a new task",
		[]*ezcli.Flag{nameFlag, categoryFlag},
		createHandler,
	)

	//Set delete Command
	ezcli.SetCommand("delete",
		"delete an existing task by id",
		[]*ezcli.Flag{idFlag},
		deleteHandler,
	)

	//Set complete Command
	ezcli.SetCommand("complete",
		"complete task by id",
		[]*ezcli.Flag{idFlag},
		completeHandler,
	)

	//Set uncomplete Command
	ezcli.SetCommand("uncomplete",
		"uncomplete task by id",
		[]*ezcli.Flag{idFlag},
		uncompleteHandler,
	)

	//Set list Command
	ezcli.SetCommand("list",
		"list all tasks",
		nil,
		listHandler,
	)

	err := ezcli.Run(os.Args)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	taskList.Save(jsonFileName)
}
