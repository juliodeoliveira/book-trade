package models

import (
	"database/sql"
)

type Book struct {
	Id int
	Name string
	Author string
	Publisher sql.NullString // acho melhor trocar para string vazia, aí no front vou ter que validar tambem
	Volume sql.NullInt64
	Year sql.NullInt64
	ImageUrl string
	UserId int64
	DeleteHash string
}

