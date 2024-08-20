package webhost

import "github.com/gorilla/mux"

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/getUser", getUserHandler).Methods("GET")
	r.HandleFunc("/isTokenValid", isJWTVaidHandler).Methods("GET")
	r.HandleFunc("/createNote", createNoteHandler).Methods("POST")
	r.HandleFunc("/getNotes", getNotesHandler).Methods("GET")
	r.HandleFunc("/getNoteByID/{id}", getNoteByIDHandler).Methods("GET")
	r.HandleFunc("/updateNote/{id}", updateNoteHandler).Methods("PUT")
	r.HandleFunc("/deleteNote/{id}", deleteNoteHandler).Methods("DELETE")

}
