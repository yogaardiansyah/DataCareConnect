// controller/router.go
package controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// SetRouter menetapkan semua rute untuk aplikasi
func SetRouter(r *mux.Router) {
	// Menangani GET request ke halaman utama
	r.HandleFunc("/", halamanUtamaHandler)
	r.HandleFunc("/index", halamanUtamaHandler)
}

// halamanUtamaHandler menangani permintaan ke halaman utama
func halamanUtamaHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		http.Error(w, "Gagal parsing template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}
