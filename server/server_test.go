package main

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	pb "ticket/proto"
)

var listener *bufconn.Listener

// Create a buffer connection for the tests
func init() {
	listener = bufconn.Listen(1024 * 1024)
	go func() {
		s := NewServer()
		grpcServer := grpc.NewServer()
		pb.RegisterTicketServiceServer(grpcServer, s)
		if err := grpcServer.Serve(listener); err != nil {
			panic(err)
		}
	}()
}

// Dial connects to the in-memory gRPC server for testing
func dialer() (*grpc.ClientConn, error) {
	return grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(
		func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}), grpc.WithInsecure())
}

// Unit tests for the TicketService
func TestSubmitPurchase(t *testing.T) {
	conn, err := dialer()
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTicketServiceClient(conn)

	user := &pb.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}

	req := &pb.PurchaseRequest{
		User:      user,
		From:      "Station A",
		To:        "Station B",
		Section:   "A",
		PricePaid: 50.0,
	}

	_, err = client.SubmitPurchase(context.Background(), req)
	if err != nil {
		t.Fatalf("SubmitPurchase failed: %v", err)
	}
}

func TestRemoveUser(t *testing.T) {
	conn, err := dialer()
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTicketServiceClient(conn)

	// First submit a purchase to create a user
	user := &pb.User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@example.com",
	}

	req := &pb.PurchaseRequest{
		User:      user,
		From:      "Station A",
		To:        "Station B",
		Section:   "A",
		PricePaid: 75.0,
	}

	_, err = client.SubmitPurchase(context.Background(), req)
	if err != nil {
		t.Fatalf("SubmitPurchase failed: %v", err)
	}

	// Now remove the user
	removeReq := &pb.UserRequest{Email: user.Email}
	res, err := client.RemoveUser(context.Background(), removeReq)
	if err != nil {
		t.Fatalf("RemoveUser failed: %v", err)
	}

	if res.Message != "User removed successfully." {
		t.Fatalf("Expected success message, got: %s", res.Message)
	}
}

func TestShowReceipt(t *testing.T) {
	conn, err := dialer()
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTicketServiceClient(conn)

	// First submit a purchase to create a user
	user := &pb.User{
		FirstName: "Alice",
		LastName:  "Smith",
		Email:     "alice.smith@example.com",
	}

	req := &pb.PurchaseRequest{
		User:      user,
		From:      "Station A",
		To:        "Station B",
		Section:   "B",
		PricePaid: 100.0,
	}

	_, err = client.SubmitPurchase(context.Background(), req)
	if err != nil {
		t.Fatalf("SubmitPurchase failed: %v", err)
	}

	// Now show the receipt
	showReq := &pb.UserRequest{Email: user.Email}
	res, err := client.ShowReceipt(context.Background(), showReq)
	if err != nil {
		t.Fatalf("ShowReceipt failed: %v", err)
	}

	if len(res.Seats) == 0 {
		t.Fatalf("Expected at least one seat, got none")
	}
}
