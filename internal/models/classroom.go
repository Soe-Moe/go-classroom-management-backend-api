package models

import "classroom-management-system/internal/database"

type Classroom struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Capacity int    `json:"capacity"`
}

func CreateClassroom(name string, capacity int) (Classroom, error) {
    var c Classroom
    err := database.DB.QueryRow(
        "INSERT INTO classrooms(name, capacity) VALUES($1, $2) RETURNING id, name, capacity",
        name, capacity,
    ).Scan(&c.ID, &c.Name, &c.Capacity)
    return c, err
}

func GetAllClassrooms() ([]Classroom, error) {
    rows, err := database.DB.Query("SELECT id, name, capacity FROM classrooms")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var classrooms []Classroom
    for rows.Next() {
        var c Classroom
        if err := rows.Scan(&c.ID, &c.Name, &c.Capacity); err != nil {
            return nil, err
        }
        classrooms = append(classrooms, c)
    }
    return classrooms, nil
}