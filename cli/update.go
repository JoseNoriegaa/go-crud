package cli

// Author: Jose Noriega
// email: josenoriega723
// Last Update: 2020-02-16

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	structs "github.com/josenoriegaa/go-crud/structs"

	"github.com/briandowns/spinner"
	"github.com/jinzhu/gorm"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// Update Update an student record
func Update(db *gorm.DB) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	var student structs.Student
	var title = "Update a student"
	reader := bufio.NewReader(os.Stdin)
	x, _ := terminal.Width()
	var userError string
	id := ""
	flag := false
	for !flag {
		Clear()
		flag = true
		// Header
		fmt.Println(strings.Repeat("=", int(x)))
		titlePrefix := strings.Repeat(" ", (int(x)/2)-(len(title)/2))
		fmt.Print(titlePrefix)
		fmt.Println(title)
		fmt.Println(strings.Repeat("=", int(x)))
		// User error
		if len(userError) > 0 {
			fmt.Println(userError)
		}
		// Inputs
		fmt.Println("Enter the student id")
		fmt.Print("> ")
		id, _ = reader.ReadString('\n')
		id = strings.ToLower(id)
		id = strings.TrimSpace(strings.Title(id))
		// Validate entry
		idIsValid, _ := regexp.MatchString("[0-9]", id)
		if !idIsValid {
			flag = false
			userError = "Please, enter a numeric id"
			continue
		}
		s.Suffix = ": Loading"
		s.Start()
		db.Where("id = ?", id).First(&student)
		s.Stop()
		if (structs.Student{}) == student {
			flag = false
			userError = "This student is not registered. Please, try it again."
		}
	}
	flag = false
	userError = ""
	firstName := ""
	lastName := ""
	title = "Update: " + student.FirstName + " " + student.LastName
	for len(firstName) == 0 && len(lastName) == 0 {
		Clear()
		// Header
		fmt.Println(strings.Repeat("=", int(x)))
		titlePrefix := strings.Repeat(" ", (int(x)/2)-(len(title)/2))
		fmt.Print(titlePrefix)
		fmt.Println(title)
		fmt.Println(strings.Repeat("=", int(x)))
		// User error
		if len(userError) > 0 {
			fmt.Println(userError)
		}
		// Inputs
		if len(firstName) <= 0 {
			fmt.Println("Enter the first name")
			fmt.Print("> ")
			firstName, _ = reader.ReadString('\n')
			firstName = strings.ToLower(firstName)
			firstName = strings.TrimSpace(strings.Title(firstName))
		}
		// Validate entry
		firstNameIsValid, _ := regexp.MatchString("[a-zA-Z]", firstName)
		if !firstNameIsValid {
			firstName = ""
			userError = "Please, enter a valid name"
			continue
		}

		if len(lastName) <= 0 {
			fmt.Println("Enter the last name")
			fmt.Print("> ")
			lastName, _ = reader.ReadString('\n')
			lastName = strings.ToLower(lastName)
			lastName = strings.TrimSpace(strings.Title(lastName))
		}
		// Validate entry
		lastNameIsValid, _ := regexp.MatchString("[a-zA-Z]", lastName)
		if !lastNameIsValid {
			lastName = ""
			userError = "Please, enter a valid last name"
		}
	}
	student.FirstName = firstName
	student.LastName = lastName
	// Confirm
	fmt.Println("- Press Y and Enter to save the changes. Otherwise, press any key and Enter to cancel.")
	fmt.Print(">")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(strings.ToLower(option))
	Clear()
	if option == "y" {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Suffix = ": Loading"
		s.Start()
		time.Sleep(time.Second)
		db.Save(&student)
		s.Stop()
	} else {
		fmt.Println("Changes has been discarted")
	}
}
