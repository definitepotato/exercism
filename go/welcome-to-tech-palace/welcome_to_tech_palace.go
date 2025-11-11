package techpalace

import (
	"bufio"
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customer = strings.ToUpper(customer)
	return fmt.Sprintf("Welcome to the Tech Palace, %s", customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := ""
	for i := 0; i < numStarsPerLine; i++ {
		border += "*"
	}

	return fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := []string{}
	scanner := bufio.NewScanner(strings.NewReader(oldMsg))
	for scanner.Scan() {
		msg = append(msg, scanner.Text())
	}

	noStars := strings.Trim(msg[1], "*")
	cleanedMsg := strings.TrimSpace(noStars)

	return cleanedMsg
}
