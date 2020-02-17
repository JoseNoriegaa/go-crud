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

	"github.com/jinzhu/gorm"
	structs "github.com/josenoriegaa/go-crud/structs"

	"github.com/briandowns/spinner"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// CaptureStudent Shows an interface to register a new student
func CaptureStudent(db *gorm.DB) {
	var title = "Student Registration Form"
	reader := bufio.NewReader(os.Stdin)
	x, _ := terminal.Width()
	var userError string
	firstName := ""
	lastName := ""

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
	// Confirm
	fmt.Println(strings.Repeat("=", int(x)))
	fmt.Printf("Capture: %s %s\n", firstName, lastName)
	fmt.Println("- Press Y and Enter to save the new student. Otherwise, press any key and Enter to cancel.")
	fmt.Print(">")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(strings.ToLower(option))
	Clear()
	if option == "y" {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Suffix = ": Loading"
		s.Start()
		time.Sleep(time.Second)
		student := structs.Student{
			FirstName: firstName,
			LastName:  lastName,
		}
		db.Create(&student)
		s.Stop()
	} else {
		fmt.Println("Changes has been discarted")
	}
}
