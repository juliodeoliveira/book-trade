package repositories

import (
	"book-trade/config"
	"book-trade/models"
	"fmt"
	"log"
)

// TODO: essas funcoes devem ir para o repositório
func GetAll() ([]models.Book, error) {
	rows, err := config.DB.Query("SELECT id, name, author, publisher, volume, year, image_url FROM books");
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book

		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publisher, &book.Volume, &book.Year, &book.ImageUrl, )
		if err != nil {
			return nil, err
		}
		
		books = append(books, book)
	}

	return books, nil
}

func AddBook(book models.Book) {
	query, err := config.DB.Prepare("INSERT INTO books (name, author, publisher, volume, year, image_url, user_id, delete_hash_imgur) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if (err != nil) {
		log.Fatal(err)
		fmt.Print(err)
	}
	defer query.Close()

	_, err = query.Exec(book.Name, book.Author, book.Publisher, book.Volume, book.Year, book.ImageUrl, book.UserId, book.DeleteHash)
	if (err != nil) {
		log.Fatal(err)
		fmt.Print(err)
	}
}

func GetBooksByUserId(userId int) ([]models.Book, error) {
	rows, err := config.DB.Query(`SELECT id, name, author, publisher, volume, year, image_url FROM books WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publisher, &book.Volume, &book.Year, &book.ImageUrl)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func UserHasThisBook(bookId int64, userId int) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = $1 AND user_id = $2)", bookId, userId).Scan(&exists)
	if err != nil || !exists {
		return false, err
	}

	return exists, nil
}

func DeleteUserBook(bookId int64) error {
	_, err := config.DB.Exec("DELETE FROM books WHERE id = $1", bookId)
	if err != nil {
		// http.Error(w, "Erro ao deletar livro", http.StatusInternalServerError)
		return err
	}

	return nil
}

func GetImageDeleteHash(bookId int64) (string, error) {
	var deleteHash string
	query := `SELECT delete_hash_imgur FROM books WHERE id = $1`
	err := config.DB.QueryRow(query, bookId).Scan(&deleteHash)
	if err != nil {
		return "", err
	}

	return deleteHash, nil
}