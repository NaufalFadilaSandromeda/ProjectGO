package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Task struct {
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Deadline    time.Time `json:"deadline"`
}

var FilePath = "data/tasks.json"

func LoadTasks() ([]Task, error) {
	file, err := os.Open(FilePath)
	if err != nil {
		return []Task{}, err
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	var tasks []Task
	err = json.Unmarshal(bytes, &tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(FilePath, bytes, 0644)
}
