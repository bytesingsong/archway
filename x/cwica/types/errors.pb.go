// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/cwica/v1/errors.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ModuleErrors defines the module level error codes
type ModuleErrors int32

const (
	// ERR_UNKNOWN is the default error code
	ModuleErrors_ERR_UNKNOWN ModuleErrors = 0
	// ERR_PACKET_TIMEOUT is the error code for packet timeout
	ModuleErrors_ERR_PACKET_TIMEOUT ModuleErrors = 1
	// ERR_EXEC_FAILURE is the error code for tx execution failure
	ModuleErrors_ERR_EXEC_FAILURE ModuleErrors = 2
)

var ModuleErrors_name = map[int32]string{
	0: "ERR_UNKNOWN",
	1: "ERR_PACKET_TIMEOUT",
	2: "ERR_EXEC_FAILURE",
}

var ModuleErrors_value = map[string]int32{
	"ERR_UNKNOWN":        0,
	"ERR_PACKET_TIMEOUT": 1,
	"ERR_EXEC_FAILURE":   2,
}

func (x ModuleErrors) String() string {
	return proto.EnumName(ModuleErrors_name, int32(x))
}

func (ModuleErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2771426f64e6ac7, []int{0}
}

func init() {
	proto.RegisterEnum("archway.cwica.v1.ModuleErrors", ModuleErrors_name, ModuleErrors_value)
}

func init() { proto.RegisterFile("archway/cwica/v1/errors.proto", fileDescriptor_c2771426f64e6ac7) }

var fileDescriptor_c2771426f64e6ac7 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0x4a, 0xce,
	0x28, 0x4f, 0xac, 0xd4, 0x4f, 0x2e, 0xcf, 0x4c, 0x4e, 0xd4, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x2a,
	0xca, 0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x4a, 0xeb, 0x81, 0xa5,
	0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x92, 0xfa, 0x20, 0x16, 0x44, 0x9d,
	0x96, 0x2f, 0x17, 0x8f, 0x6f, 0x7e, 0x4a, 0x69, 0x4e, 0xaa, 0x2b, 0x58, 0xb7, 0x10, 0x3f, 0x17,
	0xb7, 0x6b, 0x50, 0x50, 0x7c, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x83, 0x90, 0x18,
	0x97, 0x10, 0x48, 0x20, 0xc0, 0xd1, 0xd9, 0xdb, 0x35, 0x24, 0x3e, 0xc4, 0xd3, 0xd7, 0xd5, 0x3f,
	0x34, 0x44, 0x80, 0x51, 0x48, 0x84, 0x4b, 0x00, 0x24, 0xee, 0x1a, 0xe1, 0xea, 0x1c, 0xef, 0xe6,
	0xe8, 0xe9, 0x13, 0x1a, 0xe4, 0x2a, 0xc0, 0xe4, 0xe4, 0x75, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x50, 0xb7, 0xe9, 0xe6, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0xc3, 0xf8, 0xfa, 0x15, 0x50, 0xcf,
	0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x5d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x71, 0x10, 0xc7, 0xa3, 0xea, 0x00, 0x00, 0x00,
}
