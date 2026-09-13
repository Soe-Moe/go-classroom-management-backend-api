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

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
    classroomID := r.URL.Query().Get("classroom_id")
    
    var students []models.Student
    var err error

    if classroomID != "" {
        // classroom_id ပါလာရင် အဲဒီအတန်းထဲက ကျောင်းသားတွေကိုပဲ ယူမယ်
        students, err = models.GetStudentsByClassroomID(classroomID)
    } else {
        // မပါလာရင် ကျောင်းသားအားလုံးကို ယူမယ် (လိုအပ်ရင် ထည့်လို့ရပါတယ်)
        http.Error(w, "classroom_id query parameter is required", http.StatusBadRequest)
        return
    }

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(students)
}