package webhost

import (
	"encoding/json"
	"net/http"
	"notes_server/config"
	"notes_server/database"
	"notes_server/models"

	"github.com/golang-jwt/jwt/v5"
)

func extractUser(w http.ResponseWriter, r *http.Request) *models.User {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Could not find JWT cookie.")
		return nil
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT.Secret), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Could not parse JWT.")
		return nil
	}

	if !token.Valid {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Invalid token")
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		var user models.User
		database.DB.Where("username = ?", claims["sub"]).First(&user)

		return &user

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Invalid token")
		return nil
	}
}
