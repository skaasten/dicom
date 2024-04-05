package main

import (
	"fmt"
	"net/http"

	"github.com/skaasten/dicom/handlers"
)

func main() {

	http.HandleFunc("/dicom/{id}", handlers.GetByIdHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
