package note

import (
	"fmt"
	"strings"
)

func AddNotes(title, content, category string) error {
	notes, _ := LoadNotes()
	newId := 1
	if len(notes) > 0 {
		newId = notes[len(notes)-1].ID + 1
	}

	notes = append(notes, Note{
		ID:       newId,
		Title:    title,
		Content:  content,
		Category: category,
	})

	return SaveNotes(notes)
}

func ListNotes() error {
	notes, err := LoadNotes()
	if err != nil {
		return err
	}

	for _, note := range notes {
		fmt.Printf("%d: %s [%s]\n", note.ID, note.Title, note.Category)
	}

	return nil
}

func SearchNotes(keyword string) error {
	notes, err := LoadNotes()
	if err != nil {
		return err
	}
	keyword = strings.ToLower(keyword)
	for _, note := range notes {
		if strings.Contains(strings.ToLower(note.Title), keyword) || strings.Contains(strings.ToLower(note.Content), keyword) {
			fmt.Printf("%d: %s - %s [%s]\n", note.ID, note.Title, note.Content, note.Category)
		}
	}
	return nil
}

func FilterByCategory(category string) error {
	notes, err := LoadNotes()
	if err != nil {
		return err
	}

	for _, note := range notes {
		if strings.EqualFold(note.Category, category) {
			fmt.Printf("%d: %s - %s\n", note.ID, note.Title, note.Content)
		}
	}

	return nil
}

func EditNote(id int, newTitle, newContent string) error {
	notes, err := LoadNotes()
	if err != nil {
		return err
	}

	for i, note := range notes {
		if note.ID == id {
			notes[i].Title = newTitle
			notes[i].Content = newContent
			return SaveNotes(notes)
		}
	}
	return fmt.Errorf("Note with ID %d not found", id)
}

func DeleteNote(id int) error {
	notes, err := LoadNotes()
	if err != nil {
		return err
	}
	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return SaveNotes(notes)
		}
	}
	return fmt.Errorf("Note with ID %d not found", id)
}
