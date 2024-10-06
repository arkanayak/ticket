package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	"google.golang.org/grpc"
	pb "ticket/proto"
)

// server is used to implement pb.TicketServiceServer.
type server struct {
	pb.UnimplementedTicketServiceServer
	users          map[string][]*pb.UserSeat // Map of user email to slice of UserSeat
	availableSeats map[string][]int          // Map of section to available seat numbers
	bookedSeats    map[string]string         // Map of booked seats formatted as "section_seat_from_to"
	mutex          sync.Mutex                // Mutex to handle concurrent access
}

// NewServer initializes and returns a new server instance.
func NewServer() *server {
	availableSeats := make(map[string][]int)
	availableSeats["A"] = make([]int, 50) // Section A: Seats 1 to 50
	availableSeats["B"] = make([]int, 50) // Section B: Seats 1 to 50
	for i := 0; i < 50; i++ {
		availableSeats["A"][i] = i + 1
		availableSeats["B"][i] = i + 1
	}

	return &server{
		users:          make(map[string][]*pb.UserSeat), // Map to store user bookings
		availableSeats: availableSeats,
		bookedSeats:    make(map[string]string), // Initialize the booked seats map
	}
}

// SubmitPurchase handles the purchase of a ticket.
func (s *server) SubmitPurchase(ctx context.Context, req *pb.PurchaseRequest) (*pb.Receipt, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if there are available seats in the requested section
	if len(s.availableSeats[req.Section]) == 0 {
		return nil, fmt.Errorf("no available seats in section %s", req.Section)
	}

	// Find the first available seat in the requested section
	var seatNumber int
	for _, seat := range s.availableSeats[req.Section] {
		seatNumber = seat
		break
	}

	// Remove the assigned seat from the available seats
	s.availableSeats[req.Section] = removeSeat(s.availableSeats[req.Section], seatNumber)

	// Create a new UserSeat using the Protobuf generated type
	userSeat := &pb.UserSeat{
		User:       req.User,
		SeatNumber: fmt.Sprintf("%d", seatNumber), // Convert int to string
		Section:    req.Section,
	}

	// Append the new booking to the user's slice
	s.users[req.User.Email] = append(s.users[req.User.Email], userSeat)

	// Create and return the receipt
	return &pb.Receipt{
		From:       req.From,
		To:         req.To,
		User:       req.User,
		PricePaid:  req.PricePaid,
		SeatNumber: fmt.Sprintf("%d", seatNumber), // Convert int to string
		Section:    req.Section,
	}, nil
}

// ShowReceipt retrieves and displays the receipt details for a user.
func (s *server) ShowReceipt(ctx context.Context, req *pb.UserRequest) (*pb.UserList, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if the user has made any bookings
	userSeats, exists := s.users[req.Email]
	if !exists {
		return nil, fmt.Errorf("no bookings found for user: %s", req.Email)
	}

	// Create a UserList response containing all the user's booked seats
	userList := &pb.UserList{Seats: userSeats}

	// Display the user's first and last name
	for _, seat := range userSeats {
		fmt.Printf("User: %s %s, Seat: %s, Section: %s\n", seat.User.FirstName, seat.User.LastName, seat.SeatNumber, seat.Section)
	}

	return userList, nil
}

// ModifySeat modifies the user's seat based on the request.
func (s *server) ModifySeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if the user exists and has the old seat assigned
	userSeats, exists := s.users[req.Email]
	if !exists {
		return nil, fmt.Errorf("no bookings found for user: %s", req.Email)
	}

	var found bool
	var oldSeatIndex int
	for i, userSeat := range userSeats {
		if userSeat.SeatNumber == req.OldSeatNumber && userSeat.Section == req.Section {
			found = true
			oldSeatIndex = i
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("user does not have seat number %s in section %s", req.OldSeatNumber, req.Section)
	}

	// Check if the new seat is available
	newSeatAvailable := false
	for _, seat := range s.availableSeats[req.Section] {
		if fmt.Sprintf("%d", seat) == req.NewSeatNumber {
			newSeatAvailable = true
			break
		}
	}

	if !newSeatAvailable {
		return nil, fmt.Errorf("the seat number %s is not available", req.NewSeatNumber)
	}

	// Restore the old seat
	oldSeatNum, _ := strconv.Atoi(req.OldSeatNumber)
	s.availableSeats[req.Section] = append(s.availableSeats[req.Section], oldSeatNum)

	// Then proceed to update the user's seat
	s.users[req.Email][oldSeatIndex].SeatNumber = req.NewSeatNumber

	// Remove the new seat from the available seats
	newSeatNum, _ := strconv.Atoi(req.NewSeatNumber)
	s.availableSeats[req.Section] = removeSeat(s.availableSeats[req.Section], newSeatNum)

	return &pb.Response{Message: "Seat modified successfully."}, nil
}

// ViewUsersBySection retrieves the users and their allocated seats for a specified section.
func (s *server) ViewUsersBySection(ctx context.Context, req *pb.ViewUsersBySectionRequest) (*pb.ViewUsersBySectionResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if the requested section exists
	if _, exists := s.availableSeats[req.Section]; !exists {
		return nil, fmt.Errorf("section %s does not exist", req.Section)
	}

	// Prepare the response
	response := &pb.ViewUsersBySectionResponse{}

	// Iterate through the users and find those who have booked seats in the requested section
	for _, userSeats := range s.users {
		for _, userSeat := range userSeats {
			if userSeat.Section == req.Section {
				// Add the user-seat information to the response
				response.UserSeats = append(response.UserSeats, &pb.UserSeatInfo{
					User:       userSeat.User,
					SeatNumber: userSeat.SeatNumber,
				})
			}
		}
	}

	return response, nil
}

// RemoveUser removes a user's booking based on the email provided.
func (s *server) RemoveUser(ctx context.Context, req *pb.UserRequest) (*pb.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if the user exists
	userSeats, exists := s.users[req.Email]
	if !exists {
		return nil, fmt.Errorf("no bookings found for user: %s", req.Email)
	}

	// Restore the booked seats back to available
	for _, userSeat := range userSeats {
		// Convert the seat number from string to int
		seatNum, err := strconv.Atoi(userSeat.SeatNumber)
		if err != nil {
			return nil, fmt.Errorf("error converting seat number %s to int: %v", userSeat.SeatNumber, err)
		}

		// Append the seat number back to the available seats
		s.availableSeats[userSeat.Section] = append(s.availableSeats[userSeat.Section], seatNum)
	}

	// Remove user bookings
	delete(s.users, req.Email)

	return &pb.Response{Message: "User removed successfully."}, nil
}

// Helper function to remove a seat from a slice
func removeSeat(seats []int, seatNumber int) []int {
	for i, seat := range seats {
		if seat == seatNumber {
			return append(seats[:i], seats[i+1:]...)
		}
	}
	return seats
}

func main() {
	s := NewServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTicketServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
