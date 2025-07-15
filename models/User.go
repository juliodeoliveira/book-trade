package models

import "database/sql"

type User struct {
	Id int64
	Name string
	Email string
	Password string
	EmailVerified bool
	VerificationToken sql.NullString
	TokenExpiration sql.NullTime
}