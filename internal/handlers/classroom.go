package handlers

import (
	"classroom-management-system/internal/models"
	"encoding/json"
	"net/http"
)

func CreateClassroomHandler(w http.ResponseWriter, r *http.Request) {
    var req models.Classroom
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    classroom, err := models.CreateClassroom(req.Name, req.Capacity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(classroom)
}