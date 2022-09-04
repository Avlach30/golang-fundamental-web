package handler

import (
	"fundamental-webdev/entity"
	"html/template"
	"net/http"
	"path"
	"strconv"
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

func GetPlantsHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api/plants" {
		errorResponse(res, req, http.StatusNotFound, "Data not found")
		return
	}

	id := req.URL.Query().Get("id") //* Get query parameter 'id'

	temp, err := template.ParseFiles(path.Join("views", "plants.html"))
	if err != nil {
		errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
		return
	}

	//* If id query parameter is undefined
	if id == "" {

		// renderData := map[string]interface{}{
		// 	"title": "Plants page",
		// 	"content": "Plant list:\n1.Apple\n2.Strawberry",
		// }

		//* Mendeklarasikan data dengan tipe data array of object untuk ditampilkan ke template
		renderData := []entity.Plant{
			{ID: 1, Name: "padi", Description: "penghasil beras"},
			{ID: 2, Name: "cabai", Description: "penghasil rempah terbaik"},
			{ID: 3, Name: "kentang", Description: "penghasil karbohidrat"},
		}

		//* Mengeksekusi temp variabel supaya bisa dijadikan response render dengan nilai dinamis
		err = temp.Execute(res, renderData)
		if err != nil {
			errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
			return
		}
		return
	}

	idNum, error := strconv.Atoi(id) //*Convert string to integers

	if error != nil || idNum < 1 {
		errorResponse(res, req, http.StatusBadRequest, "Sorry, only integers can be accepted as query params (min 1)")
		return
	}

	//* If id query parameter is defined
	if idNum >= 1 {
		renderData := entity.Plant{
			ID:          idNum,
			Name:        "padi",
			Description: "penghasil beras",
		}

		// renderData := map[string]interface{}{
		// 	"title": "Plants page",
		// 	"content": fmt.Sprintf("Plants id: %d", idNum),
		// }

		temp, err := template.ParseFiles(path.Join("views", "plant.html"))
		if err != nil {
			errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
			return
		}

		err = temp.Execute(res, renderData)
		if err != nil {
			errorResponse(res, req, http.StatusInternalServerError, "Render failed! keep calm")
			return
		}

		return
	}

}
