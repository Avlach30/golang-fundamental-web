package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func errorResponse(res http.ResponseWriter, req *http.Request, status int, message string){
	res.WriteHeader(status)

	if status == http.StatusNotFound { //* Handler for not found error
		res.Write([]byte(message))
	}

	if status == http.StatusBadRequest { //* Handler for bad request error 
		res.Write([]byte(message))
	}

	if status == http.StatusInternalServerError { //* Handler for internal server error
		res.Write([]byte(message))
	}
}

func RootHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" { //* Jika request endpoint berbeda dengan endpoint yang telah dideklarasikan
		errorResponse(res, req, http.StatusNotFound, "Data not found")
		return
	}

	//* Memparsing file index.html sebagai template views
	temp, err := template.ParseFiles(path.Join("views", "index.html"))
	if (err != nil) {
		errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
		return
	}

	err = temp.Execute(res, nil) //* Mengeksekusi temp variabel supaya bisa dijadikan response render
	if (err != nil) {
		errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
		return
	}
	// res.Write([]byte("Hello world!"))
}

func RootApiHandler(res http.ResponseWriter, req *http.Request) {
	if(req.URL.Path != "/api") {
		errorResponse(res, req, http.StatusNotFound, "Data not found")
		return
	}

	res.Write([]byte("Hello! this is an implement of golang for web development"))
}

func GetPlantsHandler(res http.ResponseWriter, req *http.Request) {
	if(req.URL.Path != "/api/plants") {
		errorResponse(res, req, http.StatusNotFound, "Data not found")
		return
	}

	id := req.URL.Query().Get("id") //* Get query parameter 'id'

	//* If id query parameter is undefined
	if (id == "") {
		res.Write([]byte("Plant list:\n1.Apple\n2.Strawberry"))
		return
	} 

	idNum, error := strconv.Atoi(id) //*Convert string to integers

	if (error != nil || idNum < 1) {
		errorResponse(res, req, http.StatusBadRequest, "Sorry, only integers can be accepted as query params (min 1)")
		return
	}

	//* If id query parameter is defined
	if (idNum >= 1) {
		fmt.Fprintf(res, "Plants id: %d", idNum)
		return
	}

}