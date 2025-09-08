package controllers

import (
	"book-trade/middleware"
	"book-trade/models"
	"book-trade/repositories"
	"book-trade/utils"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func Home(w http.ResponseWriter, r *http.Request) {
	books, err := GetBooks(w, r)

	if (err != err) {
		http.Error(w, "Erro ao buscar os livros ", http.StatusInternalServerError)
		return
	}

	var username string
	var userBooks []models.BookView
	user, err := middleware.GetUserFromToken(r)
	if err == nil {
		username, _ = repositories.GetUsername(user.UserID)	
		userBooks = GetUserBooks(user.UserID)
	} else {
		username = "Convidado"
		userBooks = []models.BookView{}
	}

	// TODO: Proxima funcionalidade é entrar em contato para trocar

	filteredBooks := utils.SubtractBooks(books, userBooks)

	CSSFiles := utils.BuildStaticURLs([]string {
		"/static/css/indexPage/bodyConfig.css",
		"/static/css/indexPage/mainCards.css",
		"/static/css/indexPage/tradingPopup.css", 
		"/static/css/indexPage/mainLogo.css",
		"/static/css/indexPage/hamburguerMenu.css",
		"/static/css/indexPage/warning.css",
	})

	JSFiles := utils.BuildStaticURLs([]string {
		"/static/js/tradePopup.js",
		"/static/js/sidebar.js",
		"/static/js/scrollRevealConfig.js",
	})

	data := map[string]interface{}{
		"CSSFiles": CSSFiles,
		"JSFiles": JSFiles,
		"logoPath": utils.BuildSingleStaticURL("/static/no-background.png"),

		"books": filteredBooks,
		"isAuthenticated": username != "Convidado",
		"loggedUser": username,
		"userBooks" : userBooks,
	}
	
	render(w, "views/index.html", data)
}

func AddBookPage(w http.ResponseWriter, r *http.Request) {
	var username string
	user, err := middleware.GetUserFromToken(r)
	if err == nil {
		username, _ = repositories.GetUsername(user.UserID)	
	} else {
		username = "Convidado"
	}

	CSSFiles := utils.BuildStaticURLs([]string{
		"/static/css/sendBookPage/bodyConfig.css",
		"/static/css/sendBookPage/formConfig.css",
	})

	data := map[string]interface{}{
		"CSSFiles": CSSFiles,
		"isAuthenticated": username != "Convidado",
		"loggedUser": username,
	}

	render(w, "views/sendUserBook.html", data)
}

func UserBooksPage(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromToken(r)

	userBooks := GetUserBooks(user.UserID)
	username, err := repositories.GetUsername(user.UserID)
	if err != nil {
		log.Println("Erro ao buscar os livros do usuário ", err)
		return
	}

	CSSFiles := utils.BuildStaticURLs([]string {
		"/static/css/indexPage/bodyConfig.css",
		"/static/css/indexPage/mainCards.css",
		"/static/css/indexPage/tradingPopup.css", 
		"/static/css/indexPage/hamburguerMenu.css",
		"/static/css/indexPage/warning.css",
	})

	JSFiles := utils.BuildStaticURLs([]string{
		"/static/js/deleteBookConfirmation.js",
		"/static/js/sidebar.js",
		"/static/js/scrollRevealConfig.js",
	})

	data := map[string]interface{}{
		"CSSFiles": CSSFiles,
		"JSFiles": JSFiles,

		"userBooks": userBooks,
		"isAuthenticated": username != "Convidado",
		"loggedUser": username,
	}

	render(w, "views/userBooks.html", data)
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

	CSSFiles := utils.BuildStaticURLs([]string{
		"/static/css/sendBookPage/bodyConfig.css",
		"/static/css/sendBookPage/formConfig.css",
	})

	data := map[string]interface{}{
		"CSSFiles": CSSFiles,

		"Errors": errorMessage,
	}

	render(w, "views/login.html", data)
}

func Register(w http.ResponseWriter, r *http.Request) {
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

	CSSFiles := utils.BuildStaticURLs([]string{
		"/static/css/sendBookPage/bodyConfig.css",
		"/static/css/sendBookPage/formConfig.css",
	})

	data := map[string]interface{}{
		"CSSFiles": CSSFiles,

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


