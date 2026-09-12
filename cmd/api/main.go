package main

import (
	"classroom-management-system/internal/database"
	"classroom-management-system/internal/handlers"
	"log"
	"net/http"
)

func main() {
    // Database ချိတ်ဆက်ခြင်း (Environment variable တွေနဲ့ ချိတ်ရန်)
    database.InitDB("postgres://user:password@localhost/classroom_db?sslmode=disable")

    // Routes များ
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

    log.Println("Classroom Management API running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}