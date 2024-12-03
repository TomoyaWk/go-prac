package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// task
type Task struct {
	Id     int
	Title  string
	IsDone bool
}

func main() {
	file, err := os.ReadFile("./json/task.json")
	if err != nil {
		panic("file cannot read.")
	}
	tasks := []Task{}
	if err := json.Unmarshal([]byte(file), &tasks); err != nil {
		panic(err)
	}

	fmt.Println(tasks[0])
	fmt.Println(tasks[1])
}
