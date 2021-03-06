// Copyright 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protobuf/smallbank_pb2/smallbank.proto

package smallbank_pb2

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

type SmallbankTransactionPayload_PayloadType int32

const (
	SmallbankTransactionPayload_PAYLOAD_TYPE_UNSET SmallbankTransactionPayload_PayloadType = 0
	SmallbankTransactionPayload_CREATE_ACCOUNT     SmallbankTransactionPayload_PayloadType = 1
	SmallbankTransactionPayload_DEPOSIT_CHECKING   SmallbankTransactionPayload_PayloadType = 2
	SmallbankTransactionPayload_WRITE_CHECK        SmallbankTransactionPayload_PayloadType = 3
	SmallbankTransactionPayload_TRANSACT_SAVINGS   SmallbankTransactionPayload_PayloadType = 4
	SmallbankTransactionPayload_SEND_PAYMENT       SmallbankTransactionPayload_PayloadType = 5
	SmallbankTransactionPayload_AMALGAMATE         SmallbankTransactionPayload_PayloadType = 6
)

// Enum value maps for SmallbankTransactionPayload_PayloadType.
var (
	SmallbankTransactionPayload_PayloadType_name = map[int32]string{
		0: "PAYLOAD_TYPE_UNSET",
		1: "CREATE_ACCOUNT",
		2: "DEPOSIT_CHECKING",
		3: "WRITE_CHECK",
		4: "TRANSACT_SAVINGS",
		5: "SEND_PAYMENT",
		6: "AMALGAMATE",
	}
	SmallbankTransactionPayload_PayloadType_value = map[string]int32{
		"PAYLOAD_TYPE_UNSET": 0,
		"CREATE_ACCOUNT":     1,
		"DEPOSIT_CHECKING":   2,
		"WRITE_CHECK":        3,
		"TRANSACT_SAVINGS":   4,
		"SEND_PAYMENT":       5,
		"AMALGAMATE":         6,
	}
)

func (x SmallbankTransactionPayload_PayloadType) Enum() *SmallbankTransactionPayload_PayloadType {
	p := new(SmallbankTransactionPayload_PayloadType)
	*p = x
	return p
}

func (x SmallbankTransactionPayload_PayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SmallbankTransactionPayload_PayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_smallbank_pb2_smallbank_proto_enumTypes[0].Descriptor()
}

func (SmallbankTransactionPayload_PayloadType) Type() protoreflect.EnumType {
	return &file_protobuf_smallbank_pb2_smallbank_proto_enumTypes[0]
}

func (x SmallbankTransactionPayload_PayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SmallbankTransactionPayload_PayloadType.Descriptor instead.
func (SmallbankTransactionPayload_PayloadType) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{1, 0}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Customer ID
	CustomerId uint32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Customer Name
	CustomerName string `protobuf:"bytes,2,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	// Savings Balance (in cents to avoid float)
	SavingsBalance uint32 `protobuf:"varint,3,opt,name=savings_balance,json=savingsBalance,proto3" json:"savings_balance,omitempty"`
	// Checking Balance (in cents to avoid float)
	CheckingBalance uint32 `protobuf:"varint,4,opt,name=checking_balance,json=checkingBalance,proto3" json:"checking_balance,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Account) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *Account) GetSavingsBalance() uint32 {
	if x != nil {
		return x.SavingsBalance
	}
	return 0
}

func (x *Account) GetCheckingBalance() uint32 {
	if x != nil {
		return x.CheckingBalance
	}
	return 0
}

type SmallbankTransactionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadType     SmallbankTransactionPayload_PayloadType                     `protobuf:"varint,1,opt,name=payload_type,json=payloadType,proto3,enum=SmallbankTransactionPayload_PayloadType" json:"payload_type,omitempty"`
	CreateAccount   *SmallbankTransactionPayload_CreateAccountTransactionData   `protobuf:"bytes,2,opt,name=create_account,json=createAccount,proto3" json:"create_account,omitempty"`
	DepositChecking *SmallbankTransactionPayload_DepositCheckingTransactionData `protobuf:"bytes,3,opt,name=deposit_checking,json=depositChecking,proto3" json:"deposit_checking,omitempty"`
	WriteCheck      *SmallbankTransactionPayload_WriteCheckTransactionData      `protobuf:"bytes,4,opt,name=write_check,json=writeCheck,proto3" json:"write_check,omitempty"`
	TransactSavings *SmallbankTransactionPayload_TransactSavingsTransactionData `protobuf:"bytes,5,opt,name=transact_savings,json=transactSavings,proto3" json:"transact_savings,omitempty"`
	SendPayment     *SmallbankTransactionPayload_SendPaymentTransactionData     `protobuf:"bytes,6,opt,name=send_payment,json=sendPayment,proto3" json:"send_payment,omitempty"`
	Amalgamate      *SmallbankTransactionPayload_AmalgamateTransactionData      `protobuf:"bytes,7,opt,name=amalgamate,proto3" json:"amalgamate,omitempty"`
}

func (x *SmallbankTransactionPayload) Reset() {
	*x = SmallbankTransactionPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmallbankTransactionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmallbankTransactionPayload) ProtoMessage() {}

func (x *SmallbankTransactionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmallbankTransactionPayload.ProtoReflect.Descriptor instead.
func (*SmallbankTransactionPayload) Descriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{1}
}

func (x *SmallbankTransactionPayload) GetPayloadType() SmallbankTransactionPayload_PayloadType {
	if x != nil {
		return x.PayloadType
	}
	return SmallbankTransactionPayload_PAYLOAD_TYPE_UNSET
}

func (x *SmallbankTransactionPayload) GetCreateAccount() *SmallbankTransactionPayload_CreateAccountTransactionData {
	if x != nil {
		return x.CreateAccount
	}
	return nil
}

func (x *SmallbankTransactionPayload) GetDepositChecking() *SmallbankTransactionPayload_DepositCheckingTransactionData {
	if x != nil {
		return x.DepositChecking
	}
	return nil
}

func (x *SmallbankTransactionPayload) GetWriteCheck() *SmallbankTransactionPayload_WriteCheckTransactionData {
	if x != nil {
		return x.WriteCheck
	}
	return nil
}

func (x *SmallbankTransactionPayload) GetTransactSavings() *SmallbankTransactionPayload_TransactSavingsTransactionData {
	if x != nil {
		return x.TransactSavings
	}
	return nil
}

func (x *SmallbankTransactionPayload) GetSendPayment() *SmallbankTransactionPayload_SendPaymentTransactionData {
	if x != nil {
		return x.SendPayment
	}
	return nil
}

func (x *SmallbankTransactionPayload) GetAmalgamate() *SmallbankTransactionPayload_AmalgamateTransactionData {
	if x != nil {
		return x.Amalgamate
	}
	return nil
}

type SmallbankTransactionPayload_CreateAccountTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Customer ID
	CustomerId uint32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Customer Name
	CustomerName string `protobuf:"bytes,2,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	// Initial Savings Balance (in cents to avoid float)
	InitialSavingsBalance uint32 `protobuf:"varint,3,opt,name=initial_savings_balance,json=initialSavingsBalance,proto3" json:"initial_savings_balance,omitempty"`
	// Initial Checking Balance (in cents to avoid float)
	InitialCheckingBalance uint32 `protobuf:"varint,4,opt,name=initial_checking_balance,json=initialCheckingBalance,proto3" json:"initial_checking_balance,omitempty"`
}

func (x *SmallbankTransactionPayload_CreateAccountTransactionData) Reset() {
	*x = SmallbankTransactionPayload_CreateAccountTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmallbankTransactionPayload_CreateAccountTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmallbankTransactionPayload_CreateAccountTransactionData) ProtoMessage() {}

func (x *SmallbankTransactionPayload_CreateAccountTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmallbankTransactionPayload_CreateAccountTransactionData.ProtoReflect.Descriptor instead.
func (*SmallbankTransactionPayload_CreateAccountTransactionData) Descriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SmallbankTransactionPayload_CreateAccountTransactionData) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SmallbankTransactionPayload_CreateAccountTransactionData) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *SmallbankTransactionPayload_CreateAccountTransactionData) GetInitialSavingsBalance() uint32 {
	if x != nil {
		return x.InitialSavingsBalance
	}
	return 0
}

func (x *SmallbankTransactionPayload_CreateAccountTransactionData) GetInitialCheckingBalance() uint32 {
	if x != nil {
		return x.InitialCheckingBalance
	}
	return 0
}

type SmallbankTransactionPayload_DepositCheckingTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Customer ID
	CustomerId uint32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Amount
	Amount uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SmallbankTransactionPayload_DepositCheckingTransactionData) Reset() {
	*x = SmallbankTransactionPayload_DepositCheckingTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmallbankTransactionPayload_DepositCheckingTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmallbankTransactionPayload_DepositCheckingTransactionData) ProtoMessage() {}

func (x *SmallbankTransactionPayload_DepositCheckingTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmallbankTransactionPayload_DepositCheckingTransactionData.ProtoReflect.Descriptor instead.
func (*SmallbankTransactionPayload_DepositCheckingTransactionData) Descriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{1, 1}
}

func (x *SmallbankTransactionPayload_DepositCheckingTransactionData) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SmallbankTransactionPayload_DepositCheckingTransactionData) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SmallbankTransactionPayload_WriteCheckTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Customer ID
	CustomerId uint32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Amount
	Amount uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SmallbankTransactionPayload_WriteCheckTransactionData) Reset() {
	*x = SmallbankTransactionPayload_WriteCheckTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmallbankTransactionPayload_WriteCheckTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmallbankTransactionPayload_WriteCheckTransactionData) ProtoMessage() {}

func (x *SmallbankTransactionPayload_WriteCheckTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmallbankTransactionPayload_WriteCheckTransactionData.ProtoReflect.Descriptor instead.
func (*SmallbankTransactionPayload_WriteCheckTransactionData) Descriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{1, 2}
}

func (x *SmallbankTransactionPayload_WriteCheckTransactionData) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SmallbankTransactionPayload_WriteCheckTransactionData) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SmallbankTransactionPayload_TransactSavingsTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Customer ID
	CustomerId uint32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Amount
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SmallbankTransactionPayload_TransactSavingsTransactionData) Reset() {
	*x = SmallbankTransactionPayload_TransactSavingsTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmallbankTransactionPayload_TransactSavingsTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmallbankTransactionPayload_TransactSavingsTransactionData) ProtoMessage() {}

func (x *SmallbankTransactionPayload_TransactSavingsTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmallbankTransactionPayload_TransactSavingsTransactionData.ProtoReflect.Descriptor instead.
func (*SmallbankTransactionPayload_TransactSavingsTransactionData) Descriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{1, 3}
}

func (x *SmallbankTransactionPayload_TransactSavingsTransactionData) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SmallbankTransactionPayload_TransactSavingsTransactionData) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SmallbankTransactionPayload_SendPaymentTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source Customer ID
	SourceCustomerId uint32 `protobuf:"varint,1,opt,name=source_customer_id,json=sourceCustomerId,proto3" json:"source_customer_id,omitempty"`
	// Destination Customer ID
	DestCustomerId uint32 `protobuf:"varint,2,opt,name=dest_customer_id,json=destCustomerId,proto3" json:"dest_customer_id,omitempty"`
	// Amount
	Amount uint32 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SmallbankTransactionPayload_SendPaymentTransactionData) Reset() {
	*x = SmallbankTransactionPayload_SendPaymentTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmallbankTransactionPayload_SendPaymentTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmallbankTransactionPayload_SendPaymentTransactionData) ProtoMessage() {}

func (x *SmallbankTransactionPayload_SendPaymentTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmallbankTransactionPayload_SendPaymentTransactionData.ProtoReflect.Descriptor instead.
func (*SmallbankTransactionPayload_SendPaymentTransactionData) Descriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{1, 4}
}

func (x *SmallbankTransactionPayload_SendPaymentTransactionData) GetSourceCustomerId() uint32 {
	if x != nil {
		return x.SourceCustomerId
	}
	return 0
}

func (x *SmallbankTransactionPayload_SendPaymentTransactionData) GetDestCustomerId() uint32 {
	if x != nil {
		return x.DestCustomerId
	}
	return 0
}

func (x *SmallbankTransactionPayload_SendPaymentTransactionData) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SmallbankTransactionPayload_AmalgamateTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source Customer ID
	SourceCustomerId uint32 `protobuf:"varint,1,opt,name=source_customer_id,json=sourceCustomerId,proto3" json:"source_customer_id,omitempty"`
	// Destination Customer ID
	DestCustomerId uint32 `protobuf:"varint,2,opt,name=dest_customer_id,json=destCustomerId,proto3" json:"dest_customer_id,omitempty"`
}

func (x *SmallbankTransactionPayload_AmalgamateTransactionData) Reset() {
	*x = SmallbankTransactionPayload_AmalgamateTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmallbankTransactionPayload_AmalgamateTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmallbankTransactionPayload_AmalgamateTransactionData) ProtoMessage() {}

func (x *SmallbankTransactionPayload_AmalgamateTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmallbankTransactionPayload_AmalgamateTransactionData.ProtoReflect.Descriptor instead.
func (*SmallbankTransactionPayload_AmalgamateTransactionData) Descriptor() ([]byte, []int) {
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP(), []int{1, 5}
}

func (x *SmallbankTransactionPayload_AmalgamateTransactionData) GetSourceCustomerId() uint32 {
	if x != nil {
		return x.SourceCustomerId
	}
	return 0
}

func (x *SmallbankTransactionPayload_AmalgamateTransactionData) GetDestCustomerId() uint32 {
	if x != nil {
		return x.DestCustomerId
	}
	return 0
}

var File_protobuf_smallbank_pb2_smallbank_proto protoreflect.FileDescriptor

var file_protobuf_smallbank_pb2_smallbank_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x62, 0x32, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x61,
	0x76, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xad,
	0x0c, 0x0a, 0x1b, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4b,
	0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x6e, 0x6b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x60, 0x0a, 0x0e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x66, 0x0a,
	0x10, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x57, 0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x53, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x66,
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x61, 0x76, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x53,
	0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x53,
	0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5a, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x53,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x56, 0x0a, 0x0a, 0x61, 0x6d, 0x61, 0x6c, 0x67, 0x61, 0x6d, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x6e, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x41, 0x6d, 0x61, 0x6c, 0x67, 0x61, 0x6d, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a,
	0x61, 0x6d, 0x61, 0x6c, 0x67, 0x61, 0x6d, 0x61, 0x74, 0x65, 0x1a, 0xd6, 0x01, 0x0a, 0x1c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x61, 0x76,
	0x69, 0x6e, 0x67, 0x73, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x61, 0x76, 0x69, 0x6e,
	0x67, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x1a, 0x59, 0x0a, 0x1e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x54,
	0x0a, 0x19, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x59, 0x0a, 0x1e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0x8c, 0x01, 0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10,
	0x64, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x73,
	0x0a, 0x19, 0x41, 0x6d, 0x61, 0x6c, 0x67, 0x61, 0x6d, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x12, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x73,
	0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x43,
	0x48, 0x45, 0x43, 0x4b, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41,
	0x43, 0x54, 0x5f, 0x53, 0x41, 0x56, 0x49, 0x4e, 0x47, 0x53, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x45, 0x4e, 0x44, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x0e,
	0x0a, 0x0a, 0x41, 0x4d, 0x41, 0x4c, 0x47, 0x41, 0x4d, 0x41, 0x54, 0x45, 0x10, 0x06, 0x42, 0x0f,
	0x5a, 0x0d, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x62, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_smallbank_pb2_smallbank_proto_rawDescOnce sync.Once
	file_protobuf_smallbank_pb2_smallbank_proto_rawDescData = file_protobuf_smallbank_pb2_smallbank_proto_rawDesc
)

func file_protobuf_smallbank_pb2_smallbank_proto_rawDescGZIP() []byte {
	file_protobuf_smallbank_pb2_smallbank_proto_rawDescOnce.Do(func() {
		file_protobuf_smallbank_pb2_smallbank_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_smallbank_pb2_smallbank_proto_rawDescData)
	})
	return file_protobuf_smallbank_pb2_smallbank_proto_rawDescData
}

var file_protobuf_smallbank_pb2_smallbank_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_smallbank_pb2_smallbank_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protobuf_smallbank_pb2_smallbank_proto_goTypes = []interface{}{
	(SmallbankTransactionPayload_PayloadType)(0), // 0: SmallbankTransactionPayload.PayloadType
	(*Account)(nil),                     // 1: Account
	(*SmallbankTransactionPayload)(nil), // 2: SmallbankTransactionPayload
	(*SmallbankTransactionPayload_CreateAccountTransactionData)(nil),   // 3: SmallbankTransactionPayload.CreateAccountTransactionData
	(*SmallbankTransactionPayload_DepositCheckingTransactionData)(nil), // 4: SmallbankTransactionPayload.DepositCheckingTransactionData
	(*SmallbankTransactionPayload_WriteCheckTransactionData)(nil),      // 5: SmallbankTransactionPayload.WriteCheckTransactionData
	(*SmallbankTransactionPayload_TransactSavingsTransactionData)(nil), // 6: SmallbankTransactionPayload.TransactSavingsTransactionData
	(*SmallbankTransactionPayload_SendPaymentTransactionData)(nil),     // 7: SmallbankTransactionPayload.SendPaymentTransactionData
	(*SmallbankTransactionPayload_AmalgamateTransactionData)(nil),      // 8: SmallbankTransactionPayload.AmalgamateTransactionData
}
var file_protobuf_smallbank_pb2_smallbank_proto_depIdxs = []int32{
	0, // 0: SmallbankTransactionPayload.payload_type:type_name -> SmallbankTransactionPayload.PayloadType
	3, // 1: SmallbankTransactionPayload.create_account:type_name -> SmallbankTransactionPayload.CreateAccountTransactionData
	4, // 2: SmallbankTransactionPayload.deposit_checking:type_name -> SmallbankTransactionPayload.DepositCheckingTransactionData
	5, // 3: SmallbankTransactionPayload.write_check:type_name -> SmallbankTransactionPayload.WriteCheckTransactionData
	6, // 4: SmallbankTransactionPayload.transact_savings:type_name -> SmallbankTransactionPayload.TransactSavingsTransactionData
	7, // 5: SmallbankTransactionPayload.send_payment:type_name -> SmallbankTransactionPayload.SendPaymentTransactionData
	8, // 6: SmallbankTransactionPayload.amalgamate:type_name -> SmallbankTransactionPayload.AmalgamateTransactionData
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protobuf_smallbank_pb2_smallbank_proto_init() }
func file_protobuf_smallbank_pb2_smallbank_proto_init() {
	if File_protobuf_smallbank_pb2_smallbank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmallbankTransactionPayload); i {
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
		file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmallbankTransactionPayload_CreateAccountTransactionData); i {
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
		file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmallbankTransactionPayload_DepositCheckingTransactionData); i {
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
		file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmallbankTransactionPayload_WriteCheckTransactionData); i {
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
		file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmallbankTransactionPayload_TransactSavingsTransactionData); i {
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
		file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmallbankTransactionPayload_SendPaymentTransactionData); i {
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
		file_protobuf_smallbank_pb2_smallbank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmallbankTransactionPayload_AmalgamateTransactionData); i {
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
			RawDescriptor: file_protobuf_smallbank_pb2_smallbank_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_smallbank_pb2_smallbank_proto_goTypes,
		DependencyIndexes: file_protobuf_smallbank_pb2_smallbank_proto_depIdxs,
		EnumInfos:         file_protobuf_smallbank_pb2_smallbank_proto_enumTypes,
		MessageInfos:      file_protobuf_smallbank_pb2_smallbank_proto_msgTypes,
	}.Build()
	File_protobuf_smallbank_pb2_smallbank_proto = out.File
	file_protobuf_smallbank_pb2_smallbank_proto_rawDesc = nil
	file_protobuf_smallbank_pb2_smallbank_proto_goTypes = nil
	file_protobuf_smallbank_pb2_smallbank_proto_depIdxs = nil
}
