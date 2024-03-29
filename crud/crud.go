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

func AddStudent(student Student) {
	db := connection.GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO students (nim, name, study_program, phone_number, address) VALUES (?, ?, ?, ?, ?)"

	_, err := db.ExecContext(ctx, query, student.NIM, student.Name, student.StudyProgram, student.PhoneNumber, student.Address)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully added student!")
}

func ViewStudents() {
	db := connection.GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT * FROM students"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	index := 1
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.NIM, &student.Name, &student.StudyProgram, &student.PhoneNumber, &student.Address)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("%d. NIM: %s, Name: %s, Study Program: %s, Phone Number: %s, Address: %s\n", index, student.NIM, student.Name, student.StudyProgram, student.PhoneNumber, student.Address)
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func SearchStudent(query string) {
	db := connection.GetConnection()
	defer db.Close()

	ctx := context.Background()
	searchQuery := "SELECT * FROM students WHERE nim LIKE ? OR name LIKE ?"
	rows, err := db.QueryContext(ctx, searchQuery, "%"+query+"%", "%"+query+"%")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	index := 1
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.NIM, &student.Name, &student.StudyProgram, &student.PhoneNumber, &student.Address)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("%d. NIM: %s, Name: %s, Study Program: %s, Phone Number: %s, Address: %s\n", index, student.NIM, student.Name, student.StudyProgram, student.PhoneNumber, student.Address)
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if index == 1 {
		fmt.Println("No student found with NIM or Name:", query)
	}
}

func UpdateStudent(nim string, student Student) {
	db := connection.GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "UPDATE students SET nim=?, name=?, study_program=?, phone_number=?, address=? WHERE nim=?"

	result, err := db.ExecContext(ctx, query, student.NIM, student.Name, student.StudyProgram, student.PhoneNumber, student.Address, nim)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Printf("No student found with NIM: %s. Nothing was updated.\n", student.NIM)
		return
	}

	fmt.Println("Successfully updated student!")
}

func DeleteStudent(nim string) {
	db := connection.GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "DELETE FROM students WHERE nim=?"

	result, err := db.ExecContext(ctx, query, nim)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Printf("No student found with NIM: %s. Nothing was deleted.\n", nim)
		return
	}

	fmt.Println("Successfully deleted student with NIM:", nim)
}
