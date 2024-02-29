package main

import (
	"bufio"
	"fmt"
	"golang-mysql-crud/crud"
	"os"
	"strings"
)

func main() {
	repeat := true

	for repeat {
		showMenu()
		fmt.Print("Enter your choice: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number 0-5.")
			continue
		}

		fmt.Scanln()

		switch choice {
		case 0:
			repeat = false
			fmt.Println("Exiting program. Goodbye!")
		case 1:
			fmt.Println("===== Add Student =====")
			student := getStudentDetails()
			crud.AddStudent(student)
		case 2:
			fmt.Println("===== View List of Students =====")
			crud.ViewStudents()
		case 3:
			fmt.Println("===== Search for Student by NIM or Name =====")
			query := getSearchQuery()
			crud.SearchStudent(query)
		case 4:
			fmt.Println("===== Update Student =====")
			updateNIM := getUpdateNIM()
			student := getStudentDetails()
			student.NIM = updateNIM
			crud.UpdateStudent(student)
		default:
			fmt.Println("Invalid choice. Please choose a valid option (number 0-5).")
		}

		isContinue := shouldContinue()
		if !isContinue {
			repeat = false
		}
	}
}

func showMenu() {
	fmt.Println("===== STUDENT CRUD =====")
	fmt.Println("1. Add Student")
	fmt.Println("2. View List of Students")
	fmt.Println("3. Search for Student by NIM or Name")
	fmt.Println("4. Update Student")
	fmt.Println("5. Delete Student")
	fmt.Println("0. Exit")
}

func shouldContinue() bool {
	for {
		fmt.Print("Do you want to continue? (y/n): ")
		var continueChoice string
		_, err := fmt.Scan(&continueChoice)
		if err != nil {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}

		switch strings.ToLower(continueChoice) {
		case "y":
			return true
		case "n":
			fmt.Println("Exiting program. Goodbye!")
			return false
		default:
			fmt.Println("Invalid choice. Please enter 'y' or 'n'.")
		}
	}
}

func getStudentDetails() crud.Student {
	var student crud.Student
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter NIM: ")
	scanner.Scan()
	student.NIM = scanner.Text()

	fmt.Print("Enter Name: ")
	scanner.Scan()
	student.Name = scanner.Text()

	fmt.Print("Enter Study Program: ")
	scanner.Scan()
	student.StudyProgram = scanner.Text()

	fmt.Print("Enter Phone Number: ")
	scanner.Scan()
	student.PhoneNumber = scanner.Text()

	fmt.Print("Enter Address: ")
	scanner.Scan()
	student.Address = scanner.Text()

	return student
}

func getSearchQuery() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter NIM or Name to search: ")
	scanner.Scan()
	query := scanner.Text()

	return query
}

func getUpdateNIM() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter NIM of the student to update: ")
	scanner.Scan()
	nim := scanner.Text()

	return nim
}
