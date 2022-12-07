package main

import (
	"fmt"
	"booking-app/helper"
	"bufio"
	"os"
	"strconv"
	"time"
	"sync"
)

var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainningTickets uint64 = 50
var userTickets uint64
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint64
}

var wg = sync.WaitGroup {}

func main()  {

	greetUsers()

	fmt.Printf("conferenceName is %T\n", conferenceName)
	fmt.Printf("userTickets is %T\n", userTickets)
	fmt.Printf("conferenceTickets is %T\n", conferenceTickets)
	fmt.Printf("bookings is %T\n", bookings)	

	for {
		fmt.Printf("bookings size is %v\n", len(bookings))
		userFirstName, userLastName, userEmail := getUserInput()

		isValidName, isValidEmail := helper.ValidateData(userFirstName, userLastName, userEmail)
		
		if (isValidName && isValidEmail) {

			if (remainningTickets > 1) {
				fmt.Printf("Enter how many tickets you want: ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				ticketString := scanner.Text()
				userTicketsParse, err := strconv.ParseUint(ticketString, 10, 0)
				userTickets = userTicketsParse
				if (err!=nil) {
					fmt.Printf("Enter a valid number less or equal than %v remaining tickets, try again\n", remainningTickets)
					continue
				}

				isValidTicketNumber := userTickets > 0 

				if (!isValidTicketNumber) {
					fmt.Printf("Your ticket number to purchase value is invalid, try again\n")
					continue
				}
			
				if (userTickets > remainningTickets) {
					fmt.Printf("We only have %v tickets left, so you can't book %v tickets.\n", remainningTickets, userTickets)
					fmt.Printf("Would you like to book all %v of the remainning tickets?..(y or n): ", remainningTickets)
					var choice string
					fmt.Scan(&choice)
					if (choice == "y") {
						userTickets = remainningTickets
					} else {
						continue
					}
				}
			}

			if (remainningTickets == 1) {
				fmt.Printf("Only 1 ticket left\n")
				fmt.Printf("Would you like to prchase it, enter y or n: \n")
				var choice string
				fmt.Scan(&choice)
				if (choice == "y") {
					userTickets = 1
				} else {
					continue
				}
			}
		
			completeBooking(userFirstName, userLastName, userEmail)

			if (remainningTickets == 0) {
				wg.Wait()
				fmt.Printf("Our conference is booked out, come back next year!\n")
				break
			}
		} else {
			fmt.Printf("Your data input is invalid, try again\n")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have %v total tickets and %v still avaliable tickets\n", conferenceTickets, remainningTickets)
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking:= range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string) {

	var userFirstName string
	var userLastName string
	var userEmail string

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&userFirstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&userLastName)
	fmt.Printf("Enter your email: ")
	fmt.Scan(&userEmail)

	return userFirstName, userLastName, userEmail

}

func completeBooking(userFirstName string, userLastName string, userEmail string) {
	var userData = UserData {
		firstName: userFirstName,
		lastName: userLastName,
		email: userEmail,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	remainningTickets = remainningTickets - userTickets
		
	fmt.Printf("Thank you %v for purchasing %v tickets. You will recieve a confirmation email at %v soon\n", userFirstName, userTickets, userEmail)
	fmt.Printf("These are all out bookings %v\n", bookings)
	fmt.Printf("Remainning tickets are %v\n", remainningTickets)

	wg.Add(1)
	go sendTicket(userData.firstName, userData.lastName, userData.email, userData.numberOfTickets)
	fmt.Printf("These are all the firstNames: %v\n\n", getFirstNames())
}

func sendTicket(userFirstName string, userLastName string, userEmail string, userTickets uint64) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, userFirstName, userLastName)
	fmt.Printf("\n##############################################\n")
	fmt.Printf("Sending ticket: %v\nTo email address: %v\n", ticket, userEmail)
	fmt.Printf("##############################################\n")
	wg.Done()
}
