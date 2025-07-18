package repositories

import (
	"book-trade/config"
	"book-trade/models"
	"book-trade/utils"
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) {
	token, _ := utils.GenerateToken(16)
	expiration := time.Now().Add(1 * time.Hour)

	tx, err := config.DB.Begin()
	if err != nil {
		log.Println("Erro ao iniciar transação:", err)
		return
	}

	var execErr error
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if execErr != nil {
			tx.Rollback()
		} else {
			execErr = tx.Commit()
			if execErr != nil {
				log.Println("Erro ao commitar transação:", execErr)
			}
		}
	}()

	userQuery := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var userID int
	err = tx.QueryRow(userQuery, user.Name, user.Email, user.Password).Scan(&userID)
	if err != nil {
		log.Println("Erro ao inserir usuário:", err)
		return
	}


	verifyQuery := `
		INSERT INTO email_verifications (user_id, token, expiration, verified)
		VALUES ($1, $2, $3, $4)
	`
	_, err = tx.Exec(verifyQuery, userID, token, expiration, false)
	if err != nil {
		log.Println("Erro ao inserir verificação:", err)
		return
	}

	// Isso na verdade vai ser enviado para o email que o usuario colocou no formulario
	fmt.Println("http://localhost:8080/user/verify?token="+token)
}

func CheckEmail(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := config.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return true, err
	}

	return exists, nil
}

// TODO: aqui precisa retornar erros se nao der certo
func VerifyToken(token string) {
	var userId int
	query := `SELECT user_id FROM email_verifications
				WHERE token = $1 AND verified = false 
				AND expiration > NOW()`
	err := config.DB.QueryRow(query, token).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Consertar isso aqui, tem q disparar erro para o usuario
			fmt.Printf("Token inválido ")
		}
		return
	}

	updateVerification := `UPDATE email_verifications SET verified = true WHERE user_id = $1`
	_, err = config.DB.Exec(updateVerification, userId)
	if err != nil {
		log.Println("Token inválido ", err)
		return
	}

	fmt.Println("Token Validado!")
}

func Login(email string, password string) (int, error) {
	var user models.User
	err := config.DB.QueryRow(`
		SELECT u.id, u.name, u.email, u.password
		FROM users u
		JOIN email_verifications ev ON ev.user_id = u.id
		WHERE u.email = $1 AND ev.verified = true
		LIMIT 1
	`, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		// Email não verificado ou usuário não existe
		return 0, err
	} else if err != nil {
		// Qualquer outro erro
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// Senha incorreta
		return 0, err
	}

	// Gera token
	// middleware.GenerateJWT(int(user.Id))

	return int(user.Id), nil
}

func GetUsername(userId int) (string, error) {
	var username string
	query := `SELECT name FROM users WHERE id = $1`
	err := config.DB.QueryRow(query, userId).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}

