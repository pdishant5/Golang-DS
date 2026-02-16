package main

import "fmt"

func UpdateSeats(seats *[300]int8, seatNumber int) {
	if seatNumber <= 0 || seatNumber > 300 {
		fmt.Println("Invalid seat number! There are total 300 seats! Please try again!")
		return
	}

	// if (*seats)[seatNumber-1] == 1 {
	if seats[seatNumber-1] == 1 {
		fmt.Printf("Apologies! Seat %d is already booked. Please chooce another seat!\n", seatNumber)
		return
	}

	// We can access the value at the index both way in the same way as we do with the pointer receivers in methods
	// (*seats)[seatNumber-1] = 1 // right way
	seats[seatNumber-1] = 1 // syntactic sugar
	fmt.Printf("Congratulations! Your seat %d has been booked successfully!\n", seatNumber)
}

func ShowAvailableSeats(seats [300]int8) {
	fmt.Print("\n=== Available Seats ===\n\n")

	for i, seatStatus := range seats {
		if seatStatus == 0 {
			fmt.Print(i+1, " ")
		}
	}
	fmt.Println()
}

func main() {
	seats := [300]int8{}

	var choice int
	for {
		fmt.Println("1. See available seats")
		fmt.Println("2. Book a seat")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ShowAvailableSeats(seats)
		case 2:
			fmt.Print("\nYou can book multiple seats in one go! To stop booking, enter -1!\n\n")
			for {
				var seatNumber int
				fmt.Print("Enter the seat number you want to book: ")
				fmt.Scan(&seatNumber)

				if seatNumber == -1 {
					break
				}

				UpdateSeats(&seats, seatNumber)
			}

		case 3:
			return
		default:
			fmt.Println("Invalid choice. Please try again!")
		}
	}
}
