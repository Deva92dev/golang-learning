package note

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const filePath = "data/notes.json"

func LoadNotes() ([]Note, error) {
	var notes []Note

	data, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return notes, nil
	}

	// '&notes'	Passes a reference (pointer) so the function can modify the actual notes slice, Unmarshal()	Needs to write decoded data into a real memory location, not a temporary copy
	err = json.Unmarshal(data, &notes)
	return notes, err
}

func SaveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
