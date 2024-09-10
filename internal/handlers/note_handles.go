package handlers

import (
	"encoding/json"
	"net/http"
	"sports-junior-backend/internal/data"
	"sports-junior-backend/internal/models"
	"strconv"
	"strings"
)

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	notes := data.GetNotes()

	w.Header().Set("Content-Type", "application/json")
	response := models.GetNotesResponse{
		Notes: notes,
	}
	json.NewEncoder(w).Encode(response)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request models.CreateNoteRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil || request.Content == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newNote := data.CreateNote(request.Content)

	w.Header().Set("Content-Type", "application/json")
	response := models.CreateNoteResponse{
		ID: newNote.ID,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	if data.DeleteNoteByID(id) {
		w.Header().Set("Content-Type", "application/json")
		response := models.DeleteNoteResponse{
			Status: "success",
		}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Note with such ID not found", http.StatusNotFound)
	}
}

func HandleNotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNotesHandler(w, r)
	case http.MethodPost:
		createNoteHandler(w, r)
	default:
		http.Error(w, "Method not allowed on this route", http.StatusMethodNotAllowed)
	}
}
