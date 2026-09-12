package models

import "classroom-management-system/internal/database"

type Student struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    ClassroomID int    `json:"classroom_id"`
}

func EnrollStudent(name, email string, classroomID int) (Student, error) {
    var s Student
    err := database.DB.QueryRow(
        "INSERT INTO students(name, email, classroom_id) VALUES($1, $2, $3) RETURNING id, name, email, classroom_id",
        name, email, classroomID,
    ).Scan(&s.ID, &s.Name, &s.Email, &s.ClassroomID)
    return s, err
}