package main

import (
	"basics/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var RemainingTickets = 50
var bookings = make([]UserData, 0) // empty slice with userData struct

type UserData struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	// we can return multiple values and use like this
	isValidName, isValidEmail, isValidTicketCount := helper.ValidateUserInput(firstName, lastName, email, userTickets, RemainingTickets)

	if isValidName && isValidEmail && isValidTicketCount {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1) // cause there is only one other thread with main thread, that is why we put 1.
		go sendTickets(userTickets, firstName, lastName, email)

		// when we return something we have to use that return value in a variable
		firstNames := getFirstNames()
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)

		if RemainingTickets == 0 {
			fmt.Println("All tickets are booked. Come back next year")
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Your email must contain @")
		}
		if !isValidTicketCount {
			fmt.Println("You have entered invalid ticket count")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets.And remaining tickets are %v:\n", conferenceTickets, RemainingTickets)
	fmt.Println("At", conferenceName, "you can book your tickets now")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	RemainingTickets = RemainingTickets - userTickets

	// create a struct for user
	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		numOfTickets: userTickets,
	}

	// create a map for user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numOfTickets"] = strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining.\n", RemainingTickets)
}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", tickets, email)
	fmt.Println("############")

	wg.Done()
}
