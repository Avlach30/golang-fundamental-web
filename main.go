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

	mux.HandleFunc("/api/http-method", handler.FetchHTTPMethod)

	mux.HandleFunc("/add-new-plant", handler.GetAddPlantForm)
	mux.HandleFunc("/post-add-plant", handler.ProcessAddPlantForm)

	fileServer := http.FileServer(http.Dir("assets")) //* Handle static file based on 'assets' directory

	//*handle static directory so can accessed via localhost endpoint
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	log.Println("Server connected at http://localhost:5000")

	//* Menjalankan server dengan port 5000 dengan http handler nya
	error := http.ListenAndServe(":5000", mux) 

	log.Fatal(error) //* Untuk mencetak error yang terjadi
}