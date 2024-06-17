package main

import (
	"fmt"
	"time"
	"sync"
)
	
	const conferenceTickets int = 50
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = make([]map[string]string, 0)

	type UserData struct{
		firstName string
		lastName string
		email string
		numberofTickets unit
	}
	var wg = sync.WaitGroup()


func main() {

	
	greetUsers()

	for {

		firstName, lastName, userTickets := getUserInput()
		isvalidName, isvalidemail,isvalidTicketsNumber, := helper.ValidateUserInput (firstName, lastName, email, userTickets, remainingTickets)

		

		if isValidName && isValidEmail && isValidTickerNumber {

			bookTicket( userTickets, firstName, lastName,email)
			go sentTickets( userTickets, firstName, lastName,email)
			//call fucn print first name
			firstNames := getFirstNames()
			wg.Add(1)
			fmt.Println("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			} else {
				if !isValidName {
					fmt.Println("first name or last name you entered is too short")
				} else if !isValidEmail {
					fmt.Print("email address you entered doesn't contain @ sign")
				} else if !isValidTickerNumber {
					fmt.Println("number of tickets you entered id invaliid")
				}
			}
		}
		wg.Wait()
	}
}
func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v our conference \n ", confName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() [] string{
	firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, booking.firstName)
			}
			 return firstNames
}



func getuserInput(string, string, string, uint){
	
	var firstName string
	var lastName string
	var email string
	var userTickets uint


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
func bookTicket(userTickets uint,firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets
	
	//create a map for a user

	var userData = UserData {
		firstName : firstName
		lastName: lastName
		email : email
		numberofTickets : userTickets,
	}

	var myslice []string
	var mymap map[string]string

	var userData = make(map[string]string)
	userData["firstName"]=firstName
	userData["lastName"]=lastName
	userData["email"]=email
	userData["numberofTickets"]=strconv.FormatUnit(unit64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Print("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for bookings %v tickets. you will receive a confirmation email at %v \n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}

func sentTickets(userTickets unit, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf(" %v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("##################################")
	fmt.Println("Sending ticket: \n %v to email address: %v \n". tickets, email)
	fmt.Println("##################################")
	wg.Done()
}