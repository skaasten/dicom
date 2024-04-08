package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/skaasten/dicom/processor"
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
	ID string `json:"id"`
}

type GetResponse struct {
	ID    string                      `json:"id"`
	Attrs []processor.HeaderAttribute `json:"attrs"`
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
	response := UploadResponse{
		ID: key.String(),
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
	tags := []processor.Tag{}
	queryParams := r.URL.Query()
	for _, qt := range queryParams["tag"] {
		tag, err := paramToTag(qt)
		if err != nil {
			http.Error(w, "bad query param", http.StatusBadRequest)
			return
		}
		tags = append(tags, *tag)
	}
	attrs, err := h.dicom.HeaderAttributes(key, tags)
	if err != nil {
		fmt.Println("error getting attrs: %s", err)
		http.Error(w, "error getting attrs", http.StatusInternalServerError)
		return
	}

	response := GetResponse{
		ID:    key.String(),
		Attrs: attrs,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error encoding JSON: "+err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func paramToTag(s string) (*processor.Tag, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("bad tag format")
	}
	// Convert hexadecimal group and element to numerical values
	group, err := strconv.ParseUint(parts[0], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("bad tag format: %w", err)
	}
	element, err := strconv.ParseUint(parts[1], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("bad tag format: %w", err)
	}
	return &processor.Tag{uint16(group), uint16(element)}, nil
}
