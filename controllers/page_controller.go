package controllers

import (
	"book-trade/middleware"
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	books, err := GetBooks(w, r)
	if (err != err) {
		http.Error(w, "Erro ao buscar os livros ", http.StatusInternalServerError)
		return
	}

	var loggedUser string
	user, err := middleware.GetUserFromToken(r)
	if err != nil {
		loggedUser = ""
	} else {
		loggedUser = strconv.Itoa(user.UserID)
	}

	data := map[string]interface{}{
		"books": books,
		"loggedUser": loggedUser,
	}
	
	render(w, "views/index.html", data)
}

func AddBookPage(w http.ResponseWriter, r *http.Request) {
	render(w, "views/sendUserBook.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Na verdade aqui vai ter um valor que seria m erro, se senha e/ou usuario/email estiver errado ele nao deixa entrar e reotrna um erro no lugar do nil
	errorMessage := ""
	cookie, err := r.Cookie("flash")
	if err == nil {
		errorMessage = cookie.Value

		http.SetCookie(w, &http.Cookie{
			Name: "flash",
			Value: "",
			Path: "/",
			MaxAge: -1,
		})
	}

	data := map[string]interface{}{
		"Errors": errorMessage,
	}

	render(w, "views/login.html", data)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: Na verdade aqui vai ter um valor que seria m erro, se senha e/ou usuario/email estiver errado ele nao deixa entrar e reotrna um erro no lugar do nil
	cookie, err := r.Cookie("flash")
	var errorMessages []string
	if err == nil {
		decodedValue, _ := url.QueryUnescape(cookie.Value)
		json.Unmarshal([]byte(decodedValue), &errorMessages)

		http.SetCookie(w, &http.Cookie{
			Name: "flash",
			Value: "",
			Path: "/",
			MaxAge: -1,
		})
	}

	data := map[string]interface{}{
		"Errors": errorMessages,
	}

	render(w, "views/register.html", data)
}

func render(w http.ResponseWriter, filePath string, data interface{}) {
	templates, err := template.ParseFiles(filePath)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.Execute(w, data)
}

