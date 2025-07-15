package middleware

import (
	"net/http"
	"time"
	"book-trade/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", ""))

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// Gera o token JWT
func GenerateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

// Extrai e valida o token do cookie
func GetUserFromToken(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}

	tokenStr := cookie.Value

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// Middleware que exige login
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := GetUserFromToken(r)
		if err != nil {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		// Você pode adicionar as claims no contexto se quiser
		next.ServeHTTP(w, r)
	})
}
