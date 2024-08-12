package webhost

import (
	"encoding/json"
	"net/http"
	"notes_server/config"
	"notes_server/database"
	"notes_server/models"

	"github.com/golang-jwt/jwt/v5"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Could not find JWT cookie.")
		return
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT.Secret), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Could not parse JWT.")
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Invalid token")
		return
	}

	if claims, err := token.Claims.(jwt.MapClaims); err {
		var user models.User
		database.DB.Where("username = ?", claims["sub"]).First(&user)
		user.Password = ""

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Invalid token")
		return
	}

}
