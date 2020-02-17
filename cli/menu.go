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

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// Menu Render the cli menu
func Menu() int {
	var title = "CLI - MENU"
	reader := bufio.NewReader(os.Stdin)
	x, _ := terminal.Width()
	var isValid bool = false
	var showError bool = false
	var option string
	for isValid == false {
		Clear()
		// Header
		fmt.Println(strings.Repeat("=", int(x)))
		titlePrefix := strings.Repeat(" ", (int(x)/2)-(len(title)/2))
		fmt.Print(titlePrefix)
		fmt.Println(title)
		fmt.Println(strings.Repeat("=", int(x)))
		// Options
		fmt.Println(" [1] Show the registered students")
		fmt.Println(" [2] Find a student")
		fmt.Println(" [3] Register a student")
		fmt.Println(" [4] Update a student")
		fmt.Println(" [5] Delete a student")
		fmt.Println(" [6] Exit")
		// Warning user message
		if showError {
			fmt.Println("Error: You must choose an option between 1 and 4")
		}
		fmt.Print("Select an option by typing its related number (Num + Enter)> ")
		// Get input
		option, _ = reader.ReadString('\n')
		option = strings.TrimSpace(option)
		isValid, _ = regexp.MatchString("[1-6]", option)
		if !isValid {
			showError = true
		}
	}
	parsedValue, _ := strconv.Atoi(option)
	return parsedValue
}
