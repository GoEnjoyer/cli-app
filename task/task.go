package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

func New(name string, category string, isCompleted bool) *Task {
	return &Task{
		Id:          uuid.New().String(),
		Name:        name,
		Category:    category,
		IsCompleted: isCompleted,
		CreatedAt:   time.Now(),
	}
}

type TaskList struct {
	Tasks []*Task `json:"tasks"`
}

func (taskList *TaskList) Add(task *Task) {
	taskList.Tasks = append(taskList.Tasks, task)
}

func (taskList *TaskList) Remove(id string) error {
	index, err := taskList.CheckId(id)

	if err != nil {
		return err
	}

	taskList.Tasks = slices.Delete(taskList.Tasks, index, index+1)
	return nil
}

func (taskList *TaskList) ListTasks() error {
	if len(taskList.Tasks) == 0 {
		return errors.New("you don't have any tasks yet")
	}

	for i, v := range taskList.Tasks {
		fmt.Printf("%v)%v  %v  %v  %v\n", i+1, v.Name, v.Category, v.IsCompleted, v.Id)
	}

	return nil
}

func (taskList *TaskList) ChangeTaskStatus(id string, status bool) error {
	index, err := taskList.CheckId(id)

	if err != nil {
		return err
	}

	taskList.Tasks[index].IsCompleted = status
	return nil
}

func (taskList *TaskList) CheckId(id string) (index int, err error) {
	isFound := false
	isAmbiguous := false
	lengthId := len(id)

	for i, v := range taskList.Tasks {
		if v.Id[:lengthId] == id {
			if isFound {
				isAmbiguous = true
				break
			}
			index = i
			isFound = true
		}
	}

	if !isFound {
		return -1, errors.New("task is not found")
	}

	if isAmbiguous {
		return -1, errors.New("task id is ambiguous")
	}

	return index, nil
}

func (taskList *TaskList) Load(jsonFileName string) error {
	data, err := os.ReadFile(jsonFileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, taskList)
	if err != nil {
		return err
	}

	return nil
}

func (taskList *TaskList) Save(jsonFileName string) {
	data, _ := json.Marshal(taskList)
	os.WriteFile(jsonFileName, data, 0664)
}
