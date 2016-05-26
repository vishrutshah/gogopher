package world

import (
	"fmt"
	"net/http"
)

// Handle Write comments
func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World is about to rise")
}
