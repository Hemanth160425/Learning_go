package main

import (
	"encoding/json"
	"os"
)

const fileName = "task.json"

func loadTasks() ([]Task, error) {

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.WriteFile(fileName, []byte("[]"), 0644)
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
