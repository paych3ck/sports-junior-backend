package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sports-junior-backend/internal/data"
	"sports-junior-backend/internal/handlers"
	"sports-junior-backend/internal/models"
	"strings"
	"testing"
)

func mockNotesData() {
	data.CreateNote("Note with ID 0")
	data.CreateNote("Note with ID 1")
	data.CreateNote("Note with ID 2")
}

func TestGetNotesHandler(t *testing.T) {
	data.ResetNotes()
	mockNotesData()

	request, err := http.NewRequest(http.MethodGet, "/notes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleNotes)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}

	var response models.GetNotesResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Errorf("Error decoding response: %v", err)
	}

	if len(response.Notes) != 3 {
		t.Errorf("Got %d notes, expected 3", len(response.Notes))
	}
}

func TestCreateNoteHandler(t *testing.T) {
	data.ResetNotes()
	newNoteContent := `{"content":"Another created note for testing"}`

	request, err := http.NewRequest(http.MethodPost, "/notes", strings.NewReader(newNoteContent))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleNotes)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}

	var response models.CreateNoteResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Errorf("Error decoding response: %v", err)
	}

	if response.ID != 0 {
		t.Errorf("Expected note with ID 0, got ID %d", response.ID)
	}
}

func TestDeleteNoteHandler(t *testing.T) {
	data.ResetNotes()
	mockNotesData()

	request, err := http.NewRequest(http.MethodDelete, "/notes/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteNote)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}

	var response models.DeleteNoteResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Errorf("Error decoding response: %v", err)
	}

	if response.Status != "success" {
		t.Errorf("Expected success, got %v", response.Status)
	}
}

func TestDeleteNoteHandlerInvalidID(t *testing.T) {
	data.ResetNotes()
	mockNotesData()

	request, err := http.NewRequest(http.MethodDelete, "/notes/111", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteNote)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v expected %v", status, http.StatusNotFound)
	}

	expected := "Note with such ID not found\n"
	if rr.Body.String() != expected {
		t.Errorf("Expected %v, got %v", expected, rr.Body.String())
	}
}
