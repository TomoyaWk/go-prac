package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const JSON_PATH = "./json/task.json"

// task　status
type Status int

const (
	Todo Status = iota
	Done
	InProgress
)

// Status文字列変換
func (s Status) toString() string {
	switch s {
	case Todo:
		return "Todo"
	case Done:
		return "Done"
	case InProgress:
		return "InProgress"
	default:
		return "Unknown-status"
	}
}

// status json parse
func (s *Status) UnmarshalJson(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}
	parsed, err := ParseStatus(statusStr)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

func ParseStatus(s string) (Status, error) {
	switch s {
	case "Todo":
		return Todo, nil
	case "Done":
		return Done, nil
	case "InProgress":
		return InProgress, nil
	default:
		return -1, errors.New("invalid status")
	}
}

// task
type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// list
func GetTasks(query string) {
	tasks, err := readJson()
	if err != nil {
		fmt.Printf("cannot read tasks: %q", err)
	}
	var isSearch = query != ""
	var filter Status
	if isSearch {
		s, err := ParseStatus(query)
		if err != nil {
			fmt.Printf("invalid Status query: %s you can filter Todo/Done/InProgress", query)
			return
		}
		filter = s
	}

	for i := 0; i < len(tasks); i++ {
		if !isSearch || tasks[i].Status == filter.toString() {
			fmt.Printf("ID: %d | Desc:%s | Status: %s | CreatedAt: %s | UpdatedAt: %s |\n",
				tasks[i].Id,
				tasks[i].Description,
				tasks[i].Status,
				tasks[i].CreatedAt.Format(time.RFC3339),
				tasks[i].UpdatedAt.Format(time.RFC3339))
		}
	}
}

// create
func createNewTask(description string) (Task, error) {
	read, err := readJson()
	if err != nil {
		fmt.Printf("cannot read tasks: %q", err)
		return Task{}, err
	}
	now := time.Now()
	var newTask = Task{
		Id:          len(read) + 1,
		Description: description,
		Status:      Todo.toString(),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks := append(read, newTask)
	if err = saveJson(tasks); err != nil {
		return newTask, err
	}
	return newTask, nil
}

// update description
func UpdateTaskDesc(id int, description string) (Task, error) {
	read, err := readJson()
	if err != nil {
		fmt.Printf("cannot read tasks: %q", err)
		return Task{}, err
	}
	//filter
	var target Task
	var targetIndex int
	for i := 0; i < len(read); i++ {
		if read[i].Id == id {
			target = read[i]
			targetIndex = i
		}
	}
	if (target == Task{}) {
		//not found
		err := fmt.Errorf("task not found (ID: %d)", id)
		return Task{}, err
	}
	//set
	timeStamp := time.Now()
	read[targetIndex].Description = description
	read[targetIndex].UpdatedAt = timeStamp
	target = read[targetIndex]

	if err = saveJson(read); err != nil {
		return target, err
	}
	return target, nil
}

// update Status
func UpdateTaskStatus(id int, status Status) (Task, error) {
	read, err := readJson()
	if err != nil {
		fmt.Printf("cannot read tasks: %q", err)
		return Task{}, err
	}
	//filter
	var target Task
	var targetIndex int
	for i := 0; i < len(read); i++ {
		if read[i].Id == id {
			target = read[i]
			targetIndex = i
		}
	}
	if (target == Task{}) {
		//not found
		err := fmt.Errorf("task not found (ID: %d)", id)
		return Task{}, err
	}
	//set
	timeStamp := time.Now()
	read[targetIndex].Status = status.toString()
	read[targetIndex].UpdatedAt = timeStamp
	target = read[targetIndex]

	if err = saveJson(read); err != nil {
		return target, err
	}
	return target, nil
}

// delete
func DeleteTask(id int) (Task, error) {
	read, err := readJson()
	if err != nil {
		fmt.Printf("cannot read tasks: %q", err)
		return Task{}, err
	}
	//filter
	var target Task
	var targetIndex int
	for i := 0; i < len(read); i++ {
		if read[i].Id == id {
			target = read[i]
			targetIndex = i
		}
	}
	if (target == Task{}) {
		//not found
		err := fmt.Errorf("task not found (ID: %d)", id)
		return Task{}, err
	}
	//delete実行
	//対象までのスライス＋対象以降
	tasks := append(read[:targetIndex], read[targetIndex+1:]...)
	target = read[targetIndex]

	if err = saveJson(tasks); err != nil {
		return target, err
	}
	return target, nil
}

// jsonファイル読み込み
func readJson() ([]Task, error) {
	file, err := os.ReadFile(JSON_PATH)
	if err != nil {
		panic("file cannot read.")
	}
	tasks := []Task{}
	if err := json.Unmarshal([]byte(file), &tasks); err != nil {
		return tasks, err
	}
	return tasks, nil
}

// jsonファイル書き込み
func saveJson(taskData []Task) error {
	file, _ := os.Create(JSON_PATH)
	defer file.Close()
	err := json.NewEncoder(file).Encode(taskData)
	return err
}
