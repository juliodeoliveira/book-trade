package models

import (
	"book-trade/config"
	"database/sql"
	"fmt"
	"log"
)

type Book struct {
	Id int
	Name string
	Author string
	Publisher sql.NullString // acho melhor trocar para string vazia, aí no front vou ter que validar tambem
	Volume sql.NullInt64
	Year sql.NullInt64
	ImageUrl string
}


// TODO: essas funcoes devem ir para o repositório
func GetAll() ([]Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books");
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book

		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publisher, &book.Volume, &book.Year, &book.ImageUrl)
		if err != nil {
			return nil, err
		}
		
		books = append(books, book)
	}

	return books, nil
}

func AddBook(book Book) {
	query, err := config.DB.Prepare("INSERT INTO books (name, author, publisher, volume, year, image_url) VALUES ($1, $2, $3, $4, $5, $6)")
	if (err != nil) {
		log.Fatal(err)
		fmt.Print(err)
	}
	defer query.Close()

	_, err = query.Exec(book.Name, book.Author, book.Publisher, book.Volume, book.Year, book.ImageUrl)
	if (err != nil) {
		log.Fatal(err)
		fmt.Print(err)
	}
}