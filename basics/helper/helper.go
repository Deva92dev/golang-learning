package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, RemainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= RemainingTickets

	return isValidName, isValidEmail, isValidTicketCount
}
