package webhost

import (
	"encoding/json"
	"net/http"
	"notes_server/database"
	"notes_server/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	user := extractUser(w, r)
	if user == nil {
		return
	}

	var note models.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid input data.")
		return
	}

	if err := database.DB.Model(user).Association("Notes").Append(&note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to create note.")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note.ID)
}

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	user := extractUser(w, r)
	if user == nil {
		return
	}

	database.DB.Preload("Notes").Where("id = ?", user.ID).First(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user.Notes)
}

func getNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	user := extractUser(w, r)
	if user == nil {
		return
	}

	noteIDStr := mux.Vars(r)["id"]
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid note ID.")
		return
	}

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, user.ID).First(&note).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Note not found.")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Failed to retrieve note.")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	user := extractUser(w, r)
	if user == nil {
		return
	}

	noteIDStr := mux.Vars(r)["id"]
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid note ID.")
		return
	}

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, user.ID).First(&note).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Note not found.")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Failed to retrieve note.")
		}
		return
	}
	var updatedNoteData models.Note

	if err := json.NewDecoder(r.Body).Decode(&updatedNoteData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid input data.")
		return
	}

	note.Name = updatedNoteData.Name
	note.Content = updatedNoteData.Content

	if err := database.DB.Save(&note).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed to update note.")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}
