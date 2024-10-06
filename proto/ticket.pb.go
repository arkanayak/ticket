// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: proto/ticket.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// User message definition
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"` // First Name
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`    // Last Name
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`                          // Email
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// UserSeat message definition
type UserSeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`                               // User
	SeatNumber string `protobuf:"bytes,2,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"` // Seat Number
	Section    string `protobuf:"bytes,3,opt,name=section,proto3" json:"section,omitempty"`                         // Section
}

func (x *UserSeat) Reset() {
	*x = UserSeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSeat) ProtoMessage() {}

func (x *UserSeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSeat.ProtoReflect.Descriptor instead.
func (*UserSeat) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *UserSeat) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserSeat) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *UserSeat) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

// PurchaseRequest message definition
type PurchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`                              // Departure location
	To        string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`                                  // Destination location
	User      *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`                              // User making the purchase
	Section   string  `protobuf:"bytes,4,opt,name=section,proto3" json:"section,omitempty"`                        // Section of the seat
	PricePaid float32 `protobuf:"fixed32,5,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"` // Price paid for the ticket
}

func (x *PurchaseRequest) Reset() {
	*x = PurchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseRequest) ProtoMessage() {}

func (x *PurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseRequest.ProtoReflect.Descriptor instead.
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PurchaseRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *PurchaseRequest) GetPricePaid() float32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

// Receipt message definition
type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From       string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`                               // Departure location
	To         string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`                                   // Destination location
	User       *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`                               // User receiving the receipt
	SeatNumber string  `protobuf:"bytes,4,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"` // Assigned seat number
	Section    string  `protobuf:"bytes,5,opt,name=section,proto3" json:"section,omitempty"`                         // Section of the seat
	PricePaid  float32 `protobuf:"fixed32,6,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`  // Price paid for the ticket
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *Receipt) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Receipt) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Receipt) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Receipt) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *Receipt) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Receipt) GetPricePaid() float32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

// UserRequest message definition
type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // Email of the user for retrieving their booking
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// UserList message definition
type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seats []*UserSeat `protobuf:"bytes,1,rep,name=seats,proto3" json:"seats,omitempty"` // List of seats booked by the user
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *UserList) GetSeats() []*UserSeat {
	if x != nil {
		return x.Seats
	}
	return nil
}

// ModifySeatRequest message definition
type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email         string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	OldSeatNumber string `protobuf:"bytes,2,opt,name=oldSeatNumber,proto3" json:"oldSeatNumber,omitempty"` // Old seat number the user wants to modify
	NewSeatNumber string `protobuf:"bytes,3,opt,name=newSeatNumber,proto3" json:"newSeatNumber,omitempty"` // New seat number to assign
	Section       string `protobuf:"bytes,4,opt,name=section,proto3" json:"section,omitempty"`             // Section where the seat belongs
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *ModifySeatRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ModifySeatRequest) GetOldSeatNumber() string {
	if x != nil {
		return x.OldSeatNumber
	}
	return ""
}

func (x *ModifySeatRequest) GetNewSeatNumber() string {
	if x != nil {
		return x.NewSeatNumber
	}
	return ""
}

func (x *ModifySeatRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

// Response message definition
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Response message for operations
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request message for viewing users and seats by section
type ViewUsersBySectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"` // The section to query
}

func (x *ViewUsersBySectionRequest) Reset() {
	*x = ViewUsersBySectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewUsersBySectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewUsersBySectionRequest) ProtoMessage() {}

func (x *ViewUsersBySectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewUsersBySectionRequest.ProtoReflect.Descriptor instead.
func (*ViewUsersBySectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *ViewUsersBySectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

// Response message containing user-seat information
type UserSeatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`                               // The user information
	SeatNumber string `protobuf:"bytes,2,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"` // The assigned seat number
}

func (x *UserSeatInfo) Reset() {
	*x = UserSeatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSeatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSeatInfo) ProtoMessage() {}

func (x *UserSeatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSeatInfo.ProtoReflect.Descriptor instead.
func (*UserSeatInfo) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *UserSeatInfo) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserSeatInfo) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

type ViewUsersBySectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSeats []*UserSeatInfo `protobuf:"bytes,1,rep,name=user_seats,json=userSeats,proto3" json:"user_seats,omitempty"` // List of user-seat information
}

func (x *ViewUsersBySectionResponse) Reset() {
	*x = ViewUsersBySectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewUsersBySectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewUsersBySectionResponse) ProtoMessage() {}

func (x *ViewUsersBySectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewUsersBySectionResponse.ProtoReflect.Descriptor instead.
func (*ViewUsersBySectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{10}
}

func (x *ViewUsersBySectionResponse) GetUserSeats() []*UserSeatInfo {
	if x != nil {
		return x.UserSeats
	}
	return nil
}

var File_proto_ticket_proto protoreflect.FileDescriptor

var file_proto_ticket_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x58, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x67, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x90, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61,
	0x69, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x22, 0x23,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x32, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74,
	0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x6c, 0x64, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x6c, 0x64, 0x53,
	0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x65, 0x77,
	0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x35, 0x0a, 0x19, 0x56, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x1a, 0x56, 0x69, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x73, 0x32, 0xd6, 0x02, 0x0a,
	0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x12, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b,
	0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x13, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65,
	0x61, 0x74, 0x12, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x13, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x12, 0x56, 0x69, 0x65, 0x77,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ticket_proto_rawDescOnce sync.Once
	file_proto_ticket_proto_rawDescData = file_proto_ticket_proto_rawDesc
)

func file_proto_ticket_proto_rawDescGZIP() []byte {
	file_proto_ticket_proto_rawDescOnce.Do(func() {
		file_proto_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ticket_proto_rawDescData)
	})
	return file_proto_ticket_proto_rawDescData
}

var file_proto_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_ticket_proto_goTypes = []any{
	(*User)(nil),                       // 0: ticket.User
	(*UserSeat)(nil),                   // 1: ticket.UserSeat
	(*PurchaseRequest)(nil),            // 2: ticket.PurchaseRequest
	(*Receipt)(nil),                    // 3: ticket.Receipt
	(*UserRequest)(nil),                // 4: ticket.UserRequest
	(*UserList)(nil),                   // 5: ticket.UserList
	(*ModifySeatRequest)(nil),          // 6: ticket.ModifySeatRequest
	(*Response)(nil),                   // 7: ticket.Response
	(*ViewUsersBySectionRequest)(nil),  // 8: ticket.ViewUsersBySectionRequest
	(*UserSeatInfo)(nil),               // 9: ticket.UserSeatInfo
	(*ViewUsersBySectionResponse)(nil), // 10: ticket.ViewUsersBySectionResponse
}
var file_proto_ticket_proto_depIdxs = []int32{
	0,  // 0: ticket.UserSeat.user:type_name -> ticket.User
	0,  // 1: ticket.PurchaseRequest.user:type_name -> ticket.User
	0,  // 2: ticket.Receipt.user:type_name -> ticket.User
	1,  // 3: ticket.UserList.seats:type_name -> ticket.UserSeat
	0,  // 4: ticket.UserSeatInfo.user:type_name -> ticket.User
	9,  // 5: ticket.ViewUsersBySectionResponse.user_seats:type_name -> ticket.UserSeatInfo
	2,  // 6: ticket.TicketService.SubmitPurchase:input_type -> ticket.PurchaseRequest
	4,  // 7: ticket.TicketService.ShowReceipt:input_type -> ticket.UserRequest
	6,  // 8: ticket.TicketService.ModifySeat:input_type -> ticket.ModifySeatRequest
	4,  // 9: ticket.TicketService.RemoveUser:input_type -> ticket.UserRequest
	8,  // 10: ticket.TicketService.ViewUsersBySection:input_type -> ticket.ViewUsersBySectionRequest
	3,  // 11: ticket.TicketService.SubmitPurchase:output_type -> ticket.Receipt
	5,  // 12: ticket.TicketService.ShowReceipt:output_type -> ticket.UserList
	7,  // 13: ticket.TicketService.ModifySeat:output_type -> ticket.Response
	7,  // 14: ticket.TicketService.RemoveUser:output_type -> ticket.Response
	10, // 15: ticket.TicketService.ViewUsersBySection:output_type -> ticket.ViewUsersBySectionResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_ticket_proto_init() }
func file_proto_ticket_proto_init() {
	if File_proto_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ticket_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserSeat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PurchaseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Receipt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UserList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySeatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ViewUsersBySectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UserSeatInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ticket_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ViewUsersBySectionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ticket_proto_goTypes,
		DependencyIndexes: file_proto_ticket_proto_depIdxs,
		MessageInfos:      file_proto_ticket_proto_msgTypes,
	}.Build()
	File_proto_ticket_proto = out.File
	file_proto_ticket_proto_rawDesc = nil
	file_proto_ticket_proto_goTypes = nil
	file_proto_ticket_proto_depIdxs = nil
}