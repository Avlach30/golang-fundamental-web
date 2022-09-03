package main

import (
	"log"
	"net/http"
)

func main() {
	//* Menginisiasikan http handler untuk golang
	mux := http.NewServeMux()

	//* Mendeklarasikan fungsi untuk http request dengan parameter pertama endpoint, parameter kedua handler function
	mux.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello! this is an implement of golang for web development"))
	})

	log.Println("Server connected at http://localhost:5000")

	//* Menjalankan server dengan port 5000 dengan http handler nya
	error := http.ListenAndServe(":5000", mux) 
	
	log.Fatal(error) //* Untuk mencetak error yang terjadi
}