// Code generated by protoc-gen-go.
// source: chaincodeevent.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Chaincode is used for events and registrations that are specific to chaincode
// string type - "chaincode"
type ChaincodeEvent struct {
	Uuid      string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	EventName string `protobuf:"bytes,2,opt,name=eventName" json:"eventName,omitempty"`
	Payload   []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *ChaincodeEvent) Reset()         { *m = ChaincodeEvent{} }
func (m *ChaincodeEvent) String() string { return proto.CompactTextString(m) }
func (*ChaincodeEvent) ProtoMessage()    {}
