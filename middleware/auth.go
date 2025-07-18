package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateJWT(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := &jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	t, err := jwt.Parse(tokenString, func(tokenn *jwt.Token) (interface{}, error) {
		if _, ok := tokenn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing error: %v", tokenn.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authheader := r.Header.Get("Authorization")
		if authheader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authheader, "Bearer ")

		_, err := VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}


 func ExtractUserID(r *http.Request) (string, error) {
    authHeader := r.Header.Get("Authorization")
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")

    token, err := VerifyJWT(tokenString)
    if err != nil {
        return "", errors.New("missing token")
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("invalid token claims")
    }

    userID, ok := claims["userID"].(string)
    if !ok {
        return "", errors.New("userID not found")
    }
    return userID, nil
} 