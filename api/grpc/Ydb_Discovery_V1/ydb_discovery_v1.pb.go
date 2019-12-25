// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_discovery_v1.proto

package Ydb_Discovery_V1

import (
	_ "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Discovery"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("ydb_discovery_v1.proto", fileDescriptor_52888d5f10ffdaef) }

var fileDescriptor_52888d5f10ffdaef = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xab, 0x4c, 0x49, 0x8a,
	0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xaa, 0x8c, 0x2f, 0x33, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x88, 0x4c, 0x49, 0xd2, 0x73, 0x81, 0x89, 0xeb, 0x85, 0x19, 0x4a, 0xe9,
	0x64, 0x67, 0x66, 0x67, 0xe6, 0x16, 0xe9, 0x17, 0x94, 0x26, 0xe5, 0x64, 0x26, 0xeb, 0x27, 0x16,
	0x64, 0xea, 0x83, 0x95, 0x16, 0xeb, 0xa3, 0x18, 0x01, 0xd1, 0x6f, 0xb4, 0x95, 0x91, 0x4b, 0x00,
	0xae, 0x3d, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x8a, 0x8b, 0xd7, 0x27, 0xb3, 0xb8,
	0xc4, 0x35, 0x2f, 0xa5, 0x20, 0x3f, 0x33, 0xaf, 0xa4, 0x58, 0x48, 0x59, 0x0f, 0xd5, 0x1a, 0x14,
	0xd9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x15, 0xfc, 0x8a, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0x5c, 0xb9, 0xd8, 0xc2, 0x33, 0xf2, 0x1d, 0x73, 0x3d, 0x85, 0x64, 0xd0, 0xd4,
	0x43, 0x84, 0x61, 0xa6, 0xc9, 0xe2, 0x90, 0x85, 0x18, 0xe3, 0x24, 0xcb, 0x25, 0x9d, 0x9c, 0x9f,
	0xab, 0x57, 0x99, 0x98, 0x97, 0x92, 0x5a, 0xa1, 0x57, 0x99, 0x92, 0xa4, 0x87, 0xf0, 0x59, 0x99,
	0x61, 0x12, 0x1b, 0xd8, 0x77, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x7c, 0xc7, 0x5b,
	0x37, 0x01, 0x00, 0x00,
}

const (
	ListEndpoints = "/Ydb.Discovery.V1.DiscoveryService/ListEndpoints"
	WhoAmI        = "/Ydb.Discovery.V1.DiscoveryService/WhoAmI"
)