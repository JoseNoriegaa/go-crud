package cli

// Author: Jose Noriega
// email: josenoriega723@gmail.com
// Last Update: 2020-02-16

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jinzhu/gorm"
	structs "github.com/josenoriegaa/go-crud/structs"
	"github.com/olekukonko/tablewriter"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// Find Find
func Find(db *gorm.DB) {
	var flag = false
	for !flag {
		var title = "Find Student"
		reader := bufio.NewReader(os.Stdin)
		x, _ := terminal.Width()
		var userError string
		fullname := ""

		for len(fullname) == 0 {
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
			if len(fullname) <= 0 {
				fmt.Println("Enter the name")
				fmt.Print("> ")
				fullname, _ = reader.ReadString('\n')
				fullname = strings.ToLower(fullname)
				fullname = strings.TrimSpace(strings.Title(fullname))
			}
			// Validate entry
			fullnameIsValid, _ := regexp.MatchString("[a-zA-Z ]", fullname)
			if !fullnameIsValid {
				fullname = ""
				userError = "Please, enter a valid name"
			}
		}
		Clear()
		title = "Matches Found"
		var data []structs.Student
		whereField := "%" + fullname + "%"
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Suffix = ": Loading"
		s.Start()
		time.Sleep(time.Second)
		db.Where("first_name LIKE ? OR last_name LIKE ? OR CONCAT(first_name, ' ', last_name) LIKE ?", whereField, whereField, whereField).Find(&data)
		s.Stop()
		// Header
		fmt.Println(strings.Repeat("=", int(x)))
		titlePrefix := strings.Repeat(" ", (int(x)/2)-(len(title)/2))
		fmt.Print(titlePrefix)
		fmt.Println(title)
		fmt.Println(strings.Repeat("=", int(x)))
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "First Name", "Last Name", "Created At"})
		for _, item := range data {
			table.Append([]string{strconv.FormatUint(uint64(item.ID), 10), item.FirstName, item.LastName, string(item.CreatedAt.Format(time.RFC1123))})
		}
		table.Render()
		fmt.Println(strings.Repeat("=", int(x)))
		fmt.Print("To do another search press `r` and `Enter`. Otherwise, press any key and `Enter` to continue:")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		cmd = strings.ToLower(cmd)
		if cmd != "r" {
			flag = true
		}
	}
}
