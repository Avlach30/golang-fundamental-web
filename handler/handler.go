package handler

import (
	"html/template"
	"net/http"
	"path"
)

func errorResponse(res http.ResponseWriter, req *http.Request, status int, message string) {
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
	if req.URL.Path == "/" {
		errorResponse(res, req, http.StatusNotFound, "Data not found")
		return
	}
}

func RootApiHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api" { //* Jika request endpoint berbeda dengan endpoint yang telah dideklarasikan
		errorResponse(res, req, http.StatusNotFound, "Data not found")
		return
	}

	//* Memparsing file index.html sebagai template views
	temp, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
		return
	}

	err = temp.Execute(res, nil) //* Mengeksekusi temp variabel supaya bisa dijadikan response render
	if err != nil {
		errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
		return
	}
}

func FetchHTTPMethod(res http.ResponseWriter, req *http.Request) {
	METHOD := req.Method //* Fetch HTTP request method from client

	switch METHOD {
	case "GET":
		res.Write([]byte("GET method requested by client"))
	case "POST":
		res.Write([]byte("POST method requested by client"))
	default:
		errorResponse(res, req, http.StatusBadRequest, "Sorry! only GET and POST method allowed by server!")
	}
}

func GetAddPlantForm(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		temp, err := template.ParseFiles(path.Join("views", "form-add-plant.html"))
		if err != nil {
			errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
			return
		}

		err = temp.Execute(res, nil)
		if err != nil {
			errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
			return
		}
		return
	} else {
		errorResponse(res, req, http.StatusBadRequest, "Sorry! only GET method allowed by server!")
		return
	}
}

func ProcessAddPlantForm(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		err := req.ParseForm() //* Parsing form value from request client
		if err != nil {
			errorResponse(res, req, http.StatusInternalServerError, "Error occured! keep calm")
			return
		} else {
			plantName := req.Form.Get("name")
			plantDesc := req.Form.Get("description") //* Get value from input description by form

			temp, err := template.ParseFiles(path.Join("views", "plant.html"))
			if err != nil {
				errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
				return
			}

			data := map[string]interface{}{
				"ID": "dwdin93932in4239",
				"Name": plantName,
				"Description": plantDesc,
			}

			err = temp.Execute(res, data)
			if err != nil {
				errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
				return
			}
			return
		}
	} else {
		errorResponse(res, req, http.StatusBadRequest, "Sorry! only POST method allowed by server!")
		return
	}
}
