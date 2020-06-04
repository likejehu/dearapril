package handlers

import (
	"fmt"
	"net/http"
)

// Handler is struct for handlers
type Handler struct {
}

// Hello just says hello
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello!")
}
