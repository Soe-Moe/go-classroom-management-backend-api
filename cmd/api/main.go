package main

import (
	"classroom-management-system/internal/database"
	"classroom-management-system/internal/handlers"
	"log"
	"net/http"
)

func main() {
    // Database
    database.InitDB("postgres://user:password@localhost/classroom_db?sslmode=disable")

    // Routes
    http.HandleFunc("/api/classrooms", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            handlers.CreateClassroomHandler(w, r)
        }
    })

    http.HandleFunc("/api/students", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            handlers.EnrollStudentHandler(w, r)
        }
    })


// Classroom Routes
    http.HandleFunc("/api/classrooms", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            handlers.GetClassroomsHandler(w, r)
        case http.MethodPost:
            handlers.CreateClassroomHandler(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    // Student Routes
    http.HandleFunc("/api/students", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            handlers.GetStudentsHandler(w, r)
        case http.MethodPost:
            handlers.EnrollStudentHandler(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Classroom Management API running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}