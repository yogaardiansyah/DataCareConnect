// controller/server.go
package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// InitServer menginisialisasi server HTTP
func InitServer() {
	// Inisialisasi router
	r := mux.NewRouter()

	// Mengatur semua rute
	SetRouter(r)

	// Mulai server
	log.Fatal(http.ListenAndServe(":8080", r))
}
