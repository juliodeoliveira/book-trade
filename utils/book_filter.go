package utils

import (
	"book-trade/models"
)


func SubtractBooks(allBooks, userBooks []models.BookView) []models.BookView {
    userBookMap := make(map[int]bool)
    for _, b := range userBooks {
        userBookMap[b.Id] = true
    }

    var result []models.BookView
    for _, b := range allBooks {
        if !userBookMap[b.Id] {
            result = append(result, b)
        }
    }

    return result
}