package data

import (
	"sports-junior-backend/internal/models"
	"sync"
)

var (
	notes  []models.Note
	nextID int
	mu     sync.Mutex
)

func ResetNotes() {
	mu.Lock()
	defer mu.Unlock()
	notes = []models.Note{}
	nextID = 0
}

func GetNotes() []models.Note {
	mu.Lock()
	defer mu.Unlock()

	return notes
}

func CreateNote(content string) models.Note {
	mu.Lock()
	defer mu.Unlock()

	newNote := models.Note{
		ID:      nextID,
		Content: content,
	}
	notes = append(notes, newNote)
	nextID++
	return newNote
}

func DeleteNoteByID(id int) bool {
	mu.Lock()
	defer mu.Unlock()

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return true
		}
	}
	return false
}
