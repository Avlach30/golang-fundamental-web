package main

import (
	"log"
	"net/http"
	"fundamental-webdev/handler"
)

func main() {
	//* Menginisiasikan http handler untuk golang
	mux := http.NewServeMux()

	//* Mendeklarasikan fungsi untuk http request dengan parameter pertama endpoint, parameter kedua handler function
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/api", handler.RootApiHandler)
	mux.HandleFunc("/api/plants", handler.GetPlantsHandler)

	log.Println("Server connected at http://localhost:5000")

	//* Menjalankan server dengan port 5000 dengan http handler nya
	error := http.ListenAndServe(":5000", mux) 

	log.Fatal(error) //* Untuk mencetak error yang terjadi
}