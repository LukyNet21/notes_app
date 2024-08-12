package webhost

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes_server/config"
	"notes_server/database"
	"notes_server/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Unable to parse form")
		return
	}

	user := models.User{
		Username:  r.PostFormValue("username"),
		Email:     r.PostFormValue("email"),
		FirstName: r.PostFormValue("firstName"),
		LastName:  r.PostFormValue("lsatName"),
		Password:  r.PostFormValue("password"),
	}

	if user.Username == "" || user.Email == "" || user.FirstName == "" || user.LastName == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Some form fields are missing.")
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error while creating user")
		return
	}

	user.Password = string(password)
	database.DB.Create(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User successfully created")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Unable to parse form")
		return
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Some form fields are missing.")
		return
	}

	var user models.User

	database.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not fount")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid password")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"iss": "notes-app",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now(),
	})

	token, err := claims.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Could not login")
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HttpOnly: true,
		SameSite: 3, //1 - None, 2 - Lax, 3 - Strict
		Secure:   true,
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Successfully logged in")
}

func isJWTVaidHandler(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
}
