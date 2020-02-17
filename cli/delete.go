package cli

// Author: Jose Noriega
// email: josenoriega723@gmail.com
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

// Delete Delete an student record
func Delete(db *gorm.DB) {
	tx := db.Begin()
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	var student structs.Student
	var title = "Delete student"
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
		tx.Where("id = ?", id).First(&student)
		s.Stop()
		if (structs.Student{}) == student {
			flag = false
			userError = "This student is not registered. Please, try it again."
		}
	}
	// Confirm
	fmt.Println(strings.Repeat("=", int(x)))
	fmt.Printf("Student: %s %s\n", student.FirstName, student.LastName)
	fmt.Println("- Press Y and Enter to save the changes. Otherwise, press any key and Enter to cancel.")
	fmt.Print(">")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(strings.ToLower(option))
	Clear()
	if option == "y" {
		s.Start()
		time.Sleep(time.Second)
		tx.Delete(&student)
		s.Stop()
		tx.Commit()
	} else {
		fmt.Println("Changes has been discarted")
		tx.Rollback()
	}
}
