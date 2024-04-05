package handlers

import (
	"fmt"
	"net/http"
)

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "UUID: %s\n", id)
}
