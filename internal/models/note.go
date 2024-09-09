package models

type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type CreateNoteRequest struct {
	Content string `json:"content"`
}

type CreateNoteResponse struct {
	ID int `json:"id"`
}

type GetNotesResponse struct {
	Notes []Note `json:"notes"`
}

type DeleteNoteResponse struct {
	Status string `json:"status"`
}
