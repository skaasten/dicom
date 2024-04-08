package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/skaasten/dicom/service"
)

type Handlers struct {
	dicom *service.Dicom
}

func New(dicom *service.Dicom) *Handlers {
	return &Handlers{
		dicom,
	}
}

// Define a struct to represent the JSON response
type UploadResponse struct {
	id string `json:"id"`
}

func (h *Handlers) AddHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form containing the uploaded file
	err := r.ParseMultipartForm(10 << 20) // Set maximum file size to 10 MB
	if err != nil {
		http.Error(w, "error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		http.Error(w, "error copying file to buffer: "+err.Error(), http.StatusInternalServerError)
	}

	key := h.dicom.Add(buf.Bytes())
	fmt.Println("key is ", key)
	response := UploadResponse{
		id: key.String(),
	}

	// Marshal the response struct to JSON format
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error encoding JSON: "+err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func (h *Handlers) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "invalid URL", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "UUID: %s\n", id)
	key, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	_, ok := h.dicom.Get(key)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	fmt.Println("success")
	w.WriteHeader(http.StatusOK)
}
