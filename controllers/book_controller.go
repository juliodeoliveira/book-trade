package controllers
//! Esse arquivo esta muito grande, refatorar ele da melhor maneira possível
import (
	"book-trade/models"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"book-trade/services"
	"book-trade/utils"
)

type BookView struct {
	Id        int
	Name      string
	Author    string
	Publisher string
	Volume    string
	Year      string
	ImageUrl  string
}

func convertBookToView(book models.Book) BookView {

    return BookView{
        Id:        book.Id,
        Name:      book.Name,
        Author:    book.Author,
        Publisher: utils.NullableToString(book.Publisher, "Não informado"),
        Volume:    utils.NullableIntToString(book.Volume, "N/A"),
        Year:      utils.NullableIntToString(book.Year, "Desconhecido"),
		ImageUrl:  book.ImageUrl,
    }

}

func GetBooks(w http.ResponseWriter, r *http.Request) ([]BookView, error) {
    books, err := models.GetAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar os livros %v", http.StatusInternalServerError)
	}

	var booksView []BookView
	for _, book := range books{
		booksView = append(booksView, convertBookToView(book))
	}

	return booksView, nil	
}

func SendBook(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Arquivo muito grande", http.StatusRequestEntityTooLarge)
	}

	file, _, err := r.FormFile("photo")
	if err != nil {
		fmt.Println("Erro ao enviar arquivo: ", err)
		return
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Erro ao ler arquivo", nil)
		return
	}

	base64Image := base64.StdEncoding.EncodeToString(fileBytes)
	imgurUrl, err := services.UploadToImgur(base64Image)
	if err != nil {
		fmt.Println("erro na chamada da funcao para enviar para a API: ", err)
	}

	bookModel := parseBookForm(r, imgurUrl)
	models.AddBook(bookModel)

	http.Redirect(w, r, "/books/add", http.StatusSeeOther)
}

func parseBookForm(r *http.Request, imgurUrl string) models.Book {
	r.ParseForm()
	
	publisher := sql.NullString{
		String: strings.TrimSpace(r.FormValue("publisher")),
		Valid: true,
	}
	if (strings.TrimSpace(r.FormValue("publisher")) == "") {
		publisher = sql.NullString{
			String: "",
			Valid: false,
		}
	}

	bookVolume, _ := strconv.ParseInt(r.FormValue("volume"), 10, 64)
	volume := sql.NullInt64{
		Int64: bookVolume,
		Valid: true,
	}
	if (strings.TrimSpace(r.FormValue("volume")) == "") {
		volume = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	}

	bookYear, _ := strconv.ParseInt(r.FormValue("year"), 10, 64)
	year := sql.NullInt64{
		Int64: bookYear,
		Valid: true,
	}
	if (strings.TrimSpace(r.FormValue("year")) == "") {
		year = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	}


	return models.Book{
		Id: 0,
		Name: strings.TrimSpace(r.FormValue("bookName")),
		Author: strings.TrimSpace(r.FormValue("author")),
		Publisher: publisher,
		Volume: volume,
		Year: year,
		ImageUrl: imgurUrl,
	}
}
