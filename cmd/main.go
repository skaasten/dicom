package main

import (
	"fmt"
	"net/http"

	"github.com/skaasten/dicom/handlers"
	"github.com/skaasten/dicom/repository"
	"github.com/skaasten/dicom/service"
)

func main() {
	repo := repository.New()
	dicom := service.New(repo)
	h := handlers.New(dicom)

	http.HandleFunc("GET /dicom/{id}", h.GetByIdHandler)
	http.HandleFunc("POST /dicom", h.AddHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
