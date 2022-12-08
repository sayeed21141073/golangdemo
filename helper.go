package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, noOfTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := noOfTickets > 0 && noOfTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicketNumber
}
