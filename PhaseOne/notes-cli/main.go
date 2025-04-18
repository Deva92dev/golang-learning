package main

import (
	"fmt"
	"notes-cli/note"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		// total 5 commands
		fmt.Println("Usage: notes-cli [add|list|search|filter|edit|delete]")
		return
	}

	// os.Args[] It is a slice of strings ([]string) containing command-line arguments passed to the program,
	cmd := os.Args[1]
	switch cmd {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: notes-cli add <title> <content> <category>")
			return
		}
		title := os.Args[2]
		content := os.Args[3]
		category := os.Args[4]

		err := note.AddNotes(title, content, category)
		checkError(err)
		fmt.Println("📝 Note added.")

	case "list":
		err := note.ListNotes()
		checkError(err)

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: notes-cli search <keyword>")
			return
		}
		keyword := os.Args[2]
		err := note.SearchNotes(keyword)
		checkError(err)

	case "filter":
		if len(os.Args) < 3 {
			fmt.Println("Usage: notes-cli filter <category>")
			return
		}
		category := os.Args[2]
		err := note.FilterByCategory(category)
		checkError(err)

	case "edit":
		if len(os.Args) < 5 {
			fmt.Println("Usage: notes-cli edit <id> <new title> <new content>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		newTitle := os.Args[3]
		newContent := os.Args[4]
		err := note.EditNote(id, newTitle, newContent)
		checkError(err)
		fmt.Println("✏️ Note updated.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: notes-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := note.DeleteNote(id)
		checkError(err)
		fmt.Println("🗑️ Note deleted.")
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}
