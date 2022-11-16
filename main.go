package main

import (
	"fmt"
	"net/http"

	"github.com/EtheriousKelv/go-simple-web/handlers"
)

func main() {
	fmt.Println("Server is running with post 8000...")

	// Serve static file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Membuat router dan atur path dan Method
	http.HandleFunc("/", handlers.Get)

	// Tambah Data
	http.HandleFunc("/add", handlers.Post)

	// Edit Data
	// /edit?id=123
	http.HandleFunc("/edit", handlers.Edit)

	// Delete Data
	// /delete?id=1233123
	http.HandleFunc("/delete", handlers.Delete)

	// Serve Web dengan port 8000
	http.ListenAndServe(":8000", nil)
}
