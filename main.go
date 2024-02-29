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

		switch choice {
		case 0:
			repeat = false
			fmt.Println("Exiting program. Goodbye!")
		case 1:
			fmt.Println("===== Add Student =====")
			student := getStudentDetails()
			success := crud.AddStudent(student)

			if success {
				fmt.Println("Successfully added student!")
			} else {
				fmt.Println("Failed to add student. Please try again.")
			}
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
	fmt.Println("3. Search for Student")
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
	fmt.Scanln()
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
