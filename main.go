package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// fmt.Println("Welcome to", conferenceName, "booking application.")
	greetUser()

	//get user input
	firstName, lastName, email, userTickets := getUserInput()

	//validate user input
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
	if isValidName && isValidEmail && isValidTicketNumber {

		//book tickets
		bookTickets(firstName, lastName, userTickets, email)

		//send tickets
		//go keyword make goroutine
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		//print fnames
		firstNames := getFirstNames()
		fmt.Printf("These are all first names of bookings %v\n", firstNames)

		if remainingTickets == 0 {
			//end program

			fmt.Printf("%v is booked out. Come back next year.\n", conferenceName)

		}
	} else {

		if !isValidName {
			fmt.Println("First name or last name entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Invalid email id.")
		}
		if !isValidTicketNumber {
			fmt.Println("Invalid Ticket Number")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets = remainingTickets - userTickets

	//create map for users

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Println("List of bookings is", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(50 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v.", userTickets, firstName, lastName)

	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n %v \nto email %v.\n ", ticket, email)
	fmt.Println("###########")
	wg.Done()
}
