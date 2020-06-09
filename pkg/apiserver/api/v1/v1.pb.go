/*
Copyright 2020 The KubeCarrier Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

func init() {
	proto.RegisterFile("v1.proto", fileDescriptor_2e4aa7d76fd7ee8a)
}

var fileDescriptor_2e4aa7d76fd7ee8a = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x82, 0x48, 0x1a, 0x61, 0x05, 0x41, 0xab, 0x60, 0x23, 0x88, 0xd9, 0xbd, 0x9c,
	0x9d, 0x60, 0x11, 0xef, 0xae, 0x38, 0x3c, 0x44, 0xb8, 0xce, 0x6e, 0x77, 0x1d, 0x92, 0x45, 0xb3,
	0xb3, 0xcc, 0xce, 0xc6, 0xff, 0x61, 0xe9, 0xdf, 0xb3, 0xf4, 0x4f, 0xc8, 0x25, 0x1c, 0xe4, 0xba,
	0x61, 0xe6, 0xcd, 0x7b, 0xdf, 0xcb, 0x4f, 0xfb, 0x4a, 0x06, 0x42, 0x46, 0x21, 0x3e, 0x92, 0x01,
	0xab, 0x89, 0x1c, 0x90, 0xd4, 0xc1, 0xc9, 0xbe, 0xba, 0xba, 0x1b, 0x4e, 0xb6, 0x6c, 0xc0, 0x97,
	0xf1, 0x4b, 0x37, 0x0d, 0x90, 0xc2, 0xc0, 0x0e, 0x7d, 0x54, 0xda, 0x7b, 0x64, 0x3d, 0xcc, 0xa3,
	0xc3, 0xd3, 0x5f, 0xf6, 0x5d, 0xff, 0x66, 0xe2, 0x27, 0xcb, 0xcf, 0x9e, 0x93, 0x81, 0xc5, 0xe8,
	0x55, 0xd4, 0xaf, 0xeb, 0xeb, 0x32, 0x3f, 0x3f, 0x58, 0x25, 0x6e, 0x91, 0xa2, 0xb8, 0x68, 0x99,
	0x43, 0x7c, 0x50, 0x6a, 0x97, 0x4d, 0x9d, 0x66, 0x67, 0xa5, 0xc5, 0xee, 0x76, 0x9b, 0x8b, 0x3a,
	0x68, 0xdb, 0x42, 0xb1, 0x71, 0x16, 0x7c, 0x84, 0x62, 0x2e, 0x67, 0xe2, 0x71, 0xaf, 0x6e, 0x1c,
	0xb7, 0xc9, 0xec, 0x94, 0x93, 0x47, 0x35, 0xe1, 0x57, 0xe6, 0x13, 0x8d, 0xea, 0x74, 0x64, 0x20,
	0xb5, 0x59, 0x2f, 0x56, 0x2f, 0xdb, 0xd5, 0xfc, 0xb8, 0x92, 0x33, 0x5a, 0xe6, 0x97, 0x53, 0x90,
	0x25, 0xda, 0xd4, 0x81, 0x1f, 0x7b, 0x88, 0x9b, 0x7d, 0xc0, 0x3b, 0xda, 0x28, 0x0f, 0x99, 0xa6,
	0xf6, 0x6f, 0x47, 0x7d, 0x65, 0x4e, 0x86, 0xd2, 0xf7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4,
	0x91, 0x1e, 0x26, 0x42, 0x01, 0x00, 0x00,
}