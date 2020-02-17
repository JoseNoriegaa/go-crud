package cli

// Author: Jose Noriega
// email: josenoriega723
// Last Update: 2020-02-16

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	structs "github.com/josenoriegaa/go-crud/structs"
	"github.com/olekukonko/tablewriter"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// List Returns all the records into the student's table
func List(db *gorm.DB) {
	var title = "List"
	x, _ := terminal.Width()
	Clear()
	// Header
	fmt.Println(strings.Repeat("=", int(x)))
	titlePrefix := strings.Repeat(" ", (int(x)/2)-(len(title)/2))
	fmt.Print(titlePrefix)
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", int(x)))
	var data []structs.Student
	db.Find(&data)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "First Name", "Last Name", "Created At"})
	for _, item := range data {
		table.Append([]string{strconv.FormatUint(uint64(item.ID), 10), item.FirstName, item.LastName, string(item.CreatedAt.Format(time.RFC1123))})
	}
	table.Render()
	fmt.Println(strings.Repeat("=", int(x)))
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to continue: ")
	reader.ReadString('\n')
}
