package controllers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	books, err := GetBooks(w, r)
	if (err != err) {
		http.Error(w, "Erro ao buscar os livros ", http.StatusInternalServerError)
		return
	}
	render(w, "views/index.html", books)
}

func AddBookPage(w http.ResponseWriter, r *http.Request) {
	render(w, "views/sendUserBook.html", nil)
}

func render(w http.ResponseWriter, filePath string, data interface{}) {
	templates, err := template.ParseFiles(filePath)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.Execute(w, data)
}

