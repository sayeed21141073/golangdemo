package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var remainingTickets uint = 50
var conferenceName = "AHOM Limited Annual MeetUp"
var bookings = make([]user, 0)

type user struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

var waitgroup = sync.WaitGroup{}

func main() {

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	firstName, lastName, email, noOfTickets := getUserInput()
	isValidEmail, isValidName, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, noOfTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookedTickets(noOfTickets, firstName, lastName, email)

		//using goroutine

		waitgroup.Add(1)
		go sendTicket(noOfTickets, firstName, lastName, email)

		firstNames := getFirstName()
		fmt.Printf("The first names %v\n", firstNames)

		// exit application if no tickets are left
		if remainingTickets == 0 {
			// end program
			fmt.Println("Sorry! No tickets left.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
		//continue
	}
	waitgroup.Wait()

}

func greetUser(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Println("Welcome to", conferenceName, "Don't Forget To Book your Ticket(s)")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func getFirstName() []string {
	// print only first names
	firstNames := []string{}
	for _, bookingElement := range bookings {
		firstNames = append(firstNames, bookingElement.firstName)
	}
	return firstNames
}

// user input validation
// The vaidation logics are: for name, it must be 2 characters, for mail @ is must and for no of tickets
// the number range is 1-50
// we use len built in function to find length
// Contatins() function from strings package can search and find if any character(s)
// is present or not

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var noOfTickets uint

	// asking for user input. we will take 4 inputs
	//firstname, lastname, mail and no of tickets one want to buy.
	//we can take input by using a built in function from fmt package and
	//the function name is Scan(). It takes input in pointer

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email Address: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&noOfTickets)
	return firstName, lastName, email, noOfTickets
}

func bookedTickets(noOfTickets uint, firstName string, lastName string, email string) {
	// book ticket in system
	remainingTickets = remainingTickets - noOfTickets
	var userData = user{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: noOfTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v with payment instruction\n", firstName, noOfTickets, email)
	fmt.Printf("Only %v tickets are available. Hurry if U more tickets.\n", remainingTickets)
}
func sendTicket(noOfTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", noOfTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	waitgroup.Done()

}
