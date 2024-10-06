# Train Ticket Purchase System

## Overview
This project is a train ticket purchase system implemented using Golang and gRPC. It allows users to submit ticket purchases, modify seat assignments, view receipts, and manage user data.


## Features
- Submit a ticket purchase
- Show purchase receipts
- Modify seat assignments
- Remove users from the system
- View users by section

## Requirements
- Go 1.23.2
- gRPC library
- Protocol Buffers compiler (protoc)

## Setup Instructions
```bash
1. Clone the repository:

   git clone <repository-url>
   cd ticket
   
2. go mod tidy

3. protoc --go_out=. --go-grpc_out=. proto/*.proto

4. Go to server directory and run:
    go run server.go

4. Go to client directory and run in a separate tab/ terminal:
    go run client.go



SAMPLE OUTPUT/ INPUT:
go run client.go
Enter a command (submit/show/modify/remove/view/exit): submit
Enter First Name:



