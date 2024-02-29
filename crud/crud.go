package crud

import (
	"context"
	"fmt"
	"golang-mysql-crud/connection"
)

type Student struct {
	NIM          string
	Name         string
	StudyProgram string
	PhoneNumber  string
	Address      string
}

func AddStudent(student Student) bool {
	db := connection.GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO students (nim, name, study_program, phone_number, address) VALUES (?, ?, ?, ?, ?)"

	_, err := db.ExecContext(ctx, query, student.NIM, student.Name, student.StudyProgram, student.PhoneNumber, student.Address)

	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}