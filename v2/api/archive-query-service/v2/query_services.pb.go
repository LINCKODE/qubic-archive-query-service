// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v6.31.0
// source: query_services.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_query_services_proto protoreflect.FileDescriptor

var file_query_services_proto_rawDesc = []byte{
	0x0a, 0x14, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x32,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x62, 0x1a, 0x0e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb9, 0x05, 0x0a, 0x13, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xaa,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x30, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e,
	0x76, 0x32, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x71, 0x75, 0x62, 0x69,
	0x63, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x62, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x15, 0x2f, 0x67, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0xb3, 0x01, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46,
	0x6f, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x32, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76,
	0x32, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x54,
	0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x71, 0x75, 0x62,
	0x69, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x46, 0x6f, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x62, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x17, 0x2f, 0x67, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x12, 0xc3, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x36, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63,
	0x2e, 0x76, 0x32, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x62, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1b, 0x2f, 0x67, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x79, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76,
	0x32, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x67, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2d, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_query_services_proto_goTypes = []any{
	(*GetTransactionByHashRequest)(nil),        // 0: qubic.v2.archive.pb.GetTransactionByHashRequest
	(*GetTransactionsForTickRequest)(nil),      // 1: qubic.v2.archive.pb.GetTransactionsForTickRequest
	(*GetTransactionsForIdentityRequest)(nil),  // 2: qubic.v2.archive.pb.GetTransactionsForIdentityRequest
	(*GetTickDataRequest)(nil),                 // 3: qubic.v2.archive.pb.GetTickDataRequest
	(*GetTransactionByHashResponse)(nil),       // 4: qubic.v2.archive.pb.GetTransactionByHashResponse
	(*GetTransactionsForTickResponse)(nil),     // 5: qubic.v2.archive.pb.GetTransactionsForTickResponse
	(*GetTransactionsForIdentityResponse)(nil), // 6: qubic.v2.archive.pb.GetTransactionsForIdentityResponse
	(*GetTickDataResponse)(nil),                // 7: qubic.v2.archive.pb.GetTickDataResponse
}
var file_query_services_proto_depIdxs = []int32{
	0, // 0: qubic.v2.archive.pb.ArchiveQueryService.GetTransactionByHash:input_type -> qubic.v2.archive.pb.GetTransactionByHashRequest
	1, // 1: qubic.v2.archive.pb.ArchiveQueryService.GetTransactionsForTick:input_type -> qubic.v2.archive.pb.GetTransactionsForTickRequest
	2, // 2: qubic.v2.archive.pb.ArchiveQueryService.GetTransactionsForIdentity:input_type -> qubic.v2.archive.pb.GetTransactionsForIdentityRequest
	3, // 3: qubic.v2.archive.pb.ArchiveQueryService.GetTickData:input_type -> qubic.v2.archive.pb.GetTickDataRequest
	4, // 4: qubic.v2.archive.pb.ArchiveQueryService.GetTransactionByHash:output_type -> qubic.v2.archive.pb.GetTransactionByHashResponse
	5, // 5: qubic.v2.archive.pb.ArchiveQueryService.GetTransactionsForTick:output_type -> qubic.v2.archive.pb.GetTransactionsForTickResponse
	6, // 6: qubic.v2.archive.pb.ArchiveQueryService.GetTransactionsForIdentity:output_type -> qubic.v2.archive.pb.GetTransactionsForIdentityResponse
	7, // 7: qubic.v2.archive.pb.ArchiveQueryService.GetTickData:output_type -> qubic.v2.archive.pb.GetTickDataResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_query_services_proto_init() }
func file_query_services_proto_init() {
	if File_query_services_proto != nil {
		return
	}
	file_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_query_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_query_services_proto_goTypes,
		DependencyIndexes: file_query_services_proto_depIdxs,
	}.Build()
	File_query_services_proto = out.File
	file_query_services_proto_rawDesc = nil
	file_query_services_proto_goTypes = nil
	file_query_services_proto_depIdxs = nil
}
