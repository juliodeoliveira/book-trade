package controllers

import (
	"book-trade/middleware"
	"book-trade/models"
	"book-trade/repositories"
	"book-trade/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t.Execute(w, data)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if (err != nil) {
		http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
		return
	}

	username := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	var errors []string
	if password != confirmPassword {
		errors = append(errors, "As senhas não coincidem!")
	}

	isRegistered, _ := repositories.CheckEmail(email)
	if isRegistered {
		errors = append(errors, "E-mail já cadastrado!")
	}

	if len(errors) > 0 {
		jsonErrors, err := json.Marshal(errors)
		if err != nil {
			http.Error(w, "Erro interno ao processar erros", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name: "flash",
			Value: url.QueryEscape(string(jsonErrors)),
			Path: "/",
		})
		http.Redirect(w, r, "/user/register", http.StatusSeeOther)
		return
	}

	fmt.Println(username)
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println(confirmPassword)
	fmt.Println("=========================")

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Erro ao hashear a senha", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name: username,
		Email: email,
		Password: hashPassword,
	}

	repositories.CreateUser(user)

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token não fornecido", http.StatusBadRequest)
		return
	}
	
	repositories.VerifyToken(token)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	userId, errorr := repositories.Login(email, password)
	if errorr != nil {
		http.SetCookie(w, &http.Cookie{
			Name: "flash",
			Value: "Credenciais invalidas",
			Path: "/",
		})

		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}

	tokenStr, _ := middleware.GenerateJWT(userId)
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenStr,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // Remove o cookie imediatamente
		HttpOnly: true,
	})

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}