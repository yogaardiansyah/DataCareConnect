// controllers/db.go
package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL (ganti sesuai dengan jenis database yang Anda gunakan)
)

var db *sql.DB

// InitDB menginisialisasi koneksi ke database
func InitDB() error {
	// Ganti dengan konfigurasi database Anda
	connectionString := "username:password@tcp(127.0.0.1:3306)/dbname"

	// Membuka koneksi ke database
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		return fmt.Errorf("Gagal membuka koneksi ke database: %v", err)
	}

	// Memeriksa ketersediaan koneksi ke database
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Gagal melakukan ping ke database: %v", err)
	}

	fmt.Println("Koneksi ke database berhasil!")

	// Pastikan untuk menutup koneksi ke database saat aplikasi berhenti
	// defer db.Close()

	return nil
}
