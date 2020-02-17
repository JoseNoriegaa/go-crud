package cli

// Clear Clear the terminal screen
func Clear() {
	print("\033[H\033[2J")
}
