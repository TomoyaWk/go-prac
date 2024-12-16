package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

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
	Status      string    `json:status`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetTasks() {
	file, err := os.ReadFile("./json/task.json")
	if err != nil {
		panic("file cannot read.")
	}
	tasks := []Task{}
	if err := json.Unmarshal([]byte(file), &tasks); err != nil {
		panic(err)
	}

	for i := 0; i < len(tasks); i++ {
		fmt.Printf("ID: %d | Desc:%s | Status: %s | CreatedAt: %s | UpdatedAt: %s |\n",
			tasks[i].Id,
			tasks[i].Description,
			tasks[i].Status,
			tasks[i].CreatedAt.Format(time.RFC3339),
			tasks[i].UpdatedAt.Format(time.RFC3339))
	}
}
