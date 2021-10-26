package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func MakeMenu(options ...string) int {
	var menu string
	for index, value := range options {
		menu += fmt.Sprintf("%d) %s\n", index+1, value)
	}
	menu += "0) EXIT\n\n"
	menu += "Your selection: "

	return promptForUserSelection(menu, 0, len(options))
}

func promptForUserSelection(menu string, min, max int) int {
	fmt.Print(menu)
	if scanner.Scan() {
		userSelection, err := strconv.Atoi(scanner.Text())
		isValid := err != nil || userSelection > max || userSelection < min
		for isValid {
			fmt.Print("Please enter in a number value in range: ")
			if scanner.Scan() {
				userSelection, err = strconv.Atoi(scanner.Text())
				isValid = err != nil || userSelection > max || userSelection < min
			}
		}
		return userSelection
	}
	return 0
}
