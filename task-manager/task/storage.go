package task

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const filePath = "data/tasks.json"

func LoadTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil // return empty list
		}
		return nil, err
	}

	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func SaveTask(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
