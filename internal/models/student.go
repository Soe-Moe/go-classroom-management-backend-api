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

func GetStudentsByClassroomID(classroomID string) ([]Student, error) {
    rows, err := database.DB.Query("SELECT id, name, email, classroom_id FROM students WHERE classroom_id = $1", classroomID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []Student
    for rows.Next() {
        var s Student
        if err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.ClassroomID); err != nil {
            return nil, err
        }
        students = append(students, s)
    }
    return students, nil
}