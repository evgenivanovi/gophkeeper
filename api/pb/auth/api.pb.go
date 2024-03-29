// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: auth/api.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_api_proto protoreflect.FileDescriptor

var file_auth_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x73, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68,
	0x41, 0x50, 0x49, 0x12, 0x33, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x13, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x76, 0x67, 0x65,
	0x6e, 0x69, 0x76, 0x61, 0x6e, 0x6f, 0x76, 0x69, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_api_proto_goTypes = []interface{}{
	(*SigninRequest)(nil),  // 0: auth.SigninRequest
	(*SignupRequest)(nil),  // 1: auth.SignupRequest
	(*SigninResponse)(nil), // 2: auth.SigninResponse
	(*SignupResponse)(nil), // 3: auth.SignupResponse
}
var file_auth_api_proto_depIdxs = []int32{
	0, // 0: auth.AuthAPI.Signin:input_type -> auth.SigninRequest
	1, // 1: auth.AuthAPI.Signup:input_type -> auth.SignupRequest
	2, // 2: auth.AuthAPI.Signin:output_type -> auth.SigninResponse
	3, // 3: auth.AuthAPI.Signup:output_type -> auth.SignupResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_api_proto_init() }
func file_auth_api_proto_init() {
	if File_auth_api_proto != nil {
		return
	}
	file_auth_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_api_proto_goTypes,
		DependencyIndexes: file_auth_api_proto_depIdxs,
	}.Build()
	File_auth_api_proto = out.File
	file_auth_api_proto_rawDesc = nil
	file_auth_api_proto_goTypes = nil
	file_auth_api_proto_depIdxs = nil
}
