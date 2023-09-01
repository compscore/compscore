package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UserConfirm(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	defer reader.Reset(os.Stdin)

	fmt.Printf("%s [y/n]: ", prompt)

	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))

	if response == "y" || response == "yes" {
		return true
	} else if response == "n" || response == "no" {
		return false
	} else {
		fmt.Println("Invalid response; interpreting as 'no'")
		return false
	}
}
