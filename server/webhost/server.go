package webhost

import "github.com/gorilla/mux"

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/getUser", getUserHandler).Methods("POST")
	r.HandleFunc("/isTokenValid", isJWTVaidHandler).Methods("POST")
}
