package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "ticket/proto"
)

func main() {
	// Establish a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // Ensure connection is closed at the end

	client := pb.NewTicketServiceClient(conn) // Create a new client for the TicketService

	for {
		var cmd string
		fmt.Print("Enter a command (submit/show/modify/remove/view/exit): ")
		fmt.Scanln(&cmd) // Get user command input

		// Handle user commands
		switch cmd {
		case "submit":
			submitPurchase(client) // Submit a new ticket purchase
		case "show":
			showReceipt(client) // Show receipt for a user
		case "modify":
			modifySeat(client) // Modify an existing seat assignment
		case "remove":
			removeUser(client) // Remove a user from the system
		case "view":
			viewUsersBySection(client) // View users in a specific section
		case "exit":
			return // Exit the application
		default:
			fmt.Println("Invalid command. Try again.") // Handle invalid commands
		}
	}
}

func modifySeat(client pb.TicketServiceClient) {
	var email, oldSeatNumber, newSeatNumber, section string

	fmt.Print("Enter Email: ")
	fmt.Scanln(&email) // Get the user's email
	fmt.Print("Enter Old Seat Number: ")
	fmt.Scanln(&oldSeatNumber) // Get the old seat number
	fmt.Print("Enter New Seat Number: ")
	fmt.Scanln(&newSeatNumber) // Get the new seat number
	fmt.Print("Enter Section: ")
	fmt.Scanln(&section) // Get the section

	// Create a request to modify the seat
	req := &pb.ModifySeatRequest{
		Email:         email,
		OldSeatNumber: oldSeatNumber,
		NewSeatNumber: newSeatNumber,
		Section:       section,
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ModifySeat(ctx, req) // Call the ModifySeat RPC
	if err != nil {
		fmt.Printf("Error when modifying seat: %v\n", err)
		return
	}

	fmt.Println("Modify Seat Response:", res.Message) // Output the response message
}

func submitPurchase(client pb.TicketServiceClient) {
	var firstName, lastName, email, from, to, section string
	var pricePaid float32

	// Get user input for ticket purchase details
	fmt.Println("Enter First Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter Last Name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter Email:")
	fmt.Scan(&email)
	fmt.Println("Enter From Location:")
	fmt.Scan(&from)
	fmt.Println("Enter To Location:")
	fmt.Scan(&to)
	fmt.Println("Enter Section:")
	fmt.Scan(&section)
	fmt.Println("Enter Price Paid:")
	fmt.Scan(&pricePaid)

	// Create a User object from input data
	user := &pb.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	// Create a PurchaseRequest object
	req := &pb.PurchaseRequest{
		User:      user,
		From:      from,
		To:        to,
		Section:   section,
		PricePaid: pricePaid,
	}

	res, err := client.SubmitPurchase(context.Background(), req) // Call the SubmitPurchase RPC
	if err != nil {
		fmt.Printf("Error while submitting purchase: %v\n", err)
		return
	}

	// Output the purchase confirmation
	fmt.Printf("Purchase successful: User %s %s, Seat: %s, Section: %s\n", firstName, lastName, res.SeatNumber, res.Section)
}

func showReceipt(client pb.TicketServiceClient) {
	var email string
	fmt.Print("Enter Email: ")
	fmt.Scanln(&email) // Get the user's email

	req := &pb.UserRequest{Email: email} // Create a request for showing receipt

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ShowReceipt(ctx, req) // Call the ShowReceipt RPC
	if err != nil {
		fmt.Printf("Error when showing receipt: %v\n", err)
		return
	}

	// Output the seat details from the receipt
	for _, seat := range res.Seats {
		fmt.Printf("User: %s, Seat: %s, Section: %s\n", seat.User, seat.SeatNumber, seat.Section)
	}
}

func removeUser(client pb.TicketServiceClient) {
	var email string
	fmt.Print("Enter Email: ")
	fmt.Scanln(&email) // Get the user's email

	req := &pb.UserRequest{Email: email} // Create a request to remove user

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.RemoveUser(ctx, req) // Call the RemoveUser RPC
	if err != nil {
		fmt.Printf("Error when removing user: %v\n", err)
		return
	}

	// Output the response message from user removal
	fmt.Println("Remove User Response:", res.Message)
}

func viewUsersBySection(client pb.TicketServiceClient) {
	var section string
	fmt.Print("Enter Section: ")
	fmt.Scanln(&section) // Get the section to view users

	req := &pb.ViewUsersBySectionRequest{Section: section} // Create a request to view users by section

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ViewUsersBySection(ctx, req) // Call the ViewUsersBySection RPC
	if err != nil {
		fmt.Printf("Error when viewing users by section: %v\n", err)
		return
	}

	// Output user details from the specified section
	for _, userSeat := range res.UserSeats {
		fmt.Printf("User: %s %s, Seat: %s\n",
			userSeat.User.FirstName, userSeat.User.LastName, userSeat.SeatNumber)
	}
}
