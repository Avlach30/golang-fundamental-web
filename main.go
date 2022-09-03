package main

import (
	"log"
	"net/http"
)

func main() {
	//* Menginisiasikan http handler untuk golang
	mux := http.NewServeMux()

	//* Mendeklarasikan fungsi untuk http request dengan parameter pertama endpoint, parameter kedua handler function
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if(req.URL.Path != "/") { //* Jika request endpoint berbeda dengan endpoint yang telah dideklarasikan
			errorResponse(res, req, http.StatusNotFound)
			return
		}

		res.Write([]byte("Hello world!"))
	})

	mux.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		if(req.URL.Path != "/api") {
			errorResponse(res, req, http.StatusNotFound)
			return
		}

		res.Write([]byte("Hello! this is an implement of golang for web development"))
	})

	mux.HandleFunc("/api/plants", func(res http.ResponseWriter, req *http.Request) {
		if(req.URL.Path != "/api/plants") {
			errorResponse(res, req, http.StatusNotFound)
			return
		}

		res.Write([]byte("Plant list:\n1.Apple\n2.Strawberry"))
	})

	log.Println("Server connected at http://localhost:5000")

	//* Menjalankan server dengan port 5000 dengan http handler nya
	error := http.ListenAndServe(":5000", mux) 
	
	log.Fatal(error) //* Untuk mencetak error yang terjadi
}

func errorResponse(res http.ResponseWriter, req *http.Request, status int){
	res.WriteHeader(status)

	if status == http.StatusNotFound {
		res.Write([]byte("Data not found"))
	}
}