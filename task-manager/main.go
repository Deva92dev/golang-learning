package main

import (
	"fmt"
	"os"
	"strconv"
	"task-manager/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : task-manager [add|list|done|delete] [args]")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usages: task-manager add \"Task Name\" ")
			return
		}
		title := os.Args[2]
		err := task.AddTask(title)
		checkErr(err)
		fmt.Println("✅ Task Added")

	case "list":
		err := task.ListTasks()
		checkErr(err)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task manager done [id]")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := task.MarkDone(id)
		checkErr(err)
		fmt.Println("☑️ Task marked done.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("usage : task-manager delete [id]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		checkErr(err)
		err = task.DeleteTask(id)
		checkErr(err)
		fmt.Println("🗑️ Task deleted.")

	case "default":
		fmt.Println("Unknown command")
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}
