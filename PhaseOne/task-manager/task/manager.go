package task

import (
	"fmt"
	"time"
)

func AddTask(title string) error {
	tasks, _ := LoadTasks()

	newId := 1
	if len(tasks) > 0 {
		newId = tasks[len(tasks)-1].ID + 1
	}

	tasks = append(tasks, Task{
		ID:        newId,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	})

	return SaveTask(tasks)
}

func ListTasks() error {
	tasks, err := LoadTasks()

	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No Tasks Found")
		return nil
	}

	// %d	base 10, %s	the uninterpreted bytes of the string or slice
	for _, t := range tasks {
		status := "❌"
		if t.Completed {
			status = "✅"
			duration := t.DoneAt.Sub(t.CreatedAt)
			fmt.Printf("[%s] %d. %s (Completed in %s)\n", status, t.ID, t.Title, duration.Round(time.Second))
		} else {
			fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Title)
		}
	}

	return nil
}

// the error return type signifies that this function may return an error to the caller if something goes wrong — it’s part of Go’s explicit error handling philosophy.
func MarkDone(id int) error {
	tasks, _ := LoadTasks()

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			tasks[i].CreatedAt = time.Now()
			return SaveTask(tasks)
		}
	}

	return fmt.Errorf("Task #%d not found", id)
}

func DeleteTask(id int) error {
	tasks, _ := LoadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return SaveTask(tasks)
		}
	}
	return fmt.Errorf("Task #%d not found", id)
}
