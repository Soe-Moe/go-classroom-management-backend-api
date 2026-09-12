package handlers

import (
	"classroom-management-system/internal/models"
	"encoding/json"
	"net/http"
)

func EnrollStudentHandler(w http.ResponseWriter, r *http.Request) {
    var req models.Student
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    student, err := models.EnrollStudent(req.Name, req.Email, req.ClassroomID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(student)
}