// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: acme/weather/v1/types.proto

package weatherv1

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

type Condition int32

const (
	Condition_CONDITION_UNSPECIFIED Condition = 0
	Condition_CONDITION_SUNNY       Condition = 1
	Condition_CONDITION_RAINY       Condition = 2
)

// Enum value maps for Condition.
var (
	Condition_name = map[int32]string{
		0: "CONDITION_UNSPECIFIED",
		1: "CONDITION_SUNNY",
		2: "CONDITION_RAINY",
	}
	Condition_value = map[string]int32{
		"CONDITION_UNSPECIFIED": 0,
		"CONDITION_SUNNY":       1,
		"CONDITION_RAINY":       2,
	}
)

func (x Condition) Enum() *Condition {
	p := new(Condition)
	*p = x
	return p
}

func (x Condition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition) Descriptor() protoreflect.EnumDescriptor {
	return file_acme_weather_v1_types_proto_enumTypes[0].Descriptor()
}

func (Condition) Type() protoreflect.EnumType {
	return &file_acme_weather_v1_types_proto_enumTypes[0]
}

func (x Condition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition.Descriptor instead.
func (Condition) EnumDescriptor() ([]byte, []int) {
	return file_acme_weather_v1_types_proto_rawDescGZIP(), []int{0}
}

var File_acme_weather_v1_types_proto protoreflect.FileDescriptor

var file_acme_weather_v1_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61,
	0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2a, 0x50,
	0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x43,
	0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x4e, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43,
	0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x41, 0x49, 0x4e, 0x59, 0x10, 0x02,
	0x42, 0xc0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x64, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x2f, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x2d, 0x62, 0x75, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x63,
	0x6d, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x57, 0x58, 0xaa, 0x02, 0x0f,
	0x41, 0x63, 0x6d, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0f, 0x41, 0x63, 0x6d, 0x65, 0x5c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1b, 0x41, 0x63, 0x6d, 0x65, 0x5c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x11, 0x41, 0x63, 0x6d, 0x65, 0x3a, 0x3a, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acme_weather_v1_types_proto_rawDescOnce sync.Once
	file_acme_weather_v1_types_proto_rawDescData = file_acme_weather_v1_types_proto_rawDesc
)

func file_acme_weather_v1_types_proto_rawDescGZIP() []byte {
	file_acme_weather_v1_types_proto_rawDescOnce.Do(func() {
		file_acme_weather_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_acme_weather_v1_types_proto_rawDescData)
	})
	return file_acme_weather_v1_types_proto_rawDescData
}

var file_acme_weather_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_acme_weather_v1_types_proto_goTypes = []any{
	(Condition)(0), // 0: acme.weather.v1.Condition
}
var file_acme_weather_v1_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_acme_weather_v1_types_proto_init() }
func file_acme_weather_v1_types_proto_init() {
	if File_acme_weather_v1_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_acme_weather_v1_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acme_weather_v1_types_proto_goTypes,
		DependencyIndexes: file_acme_weather_v1_types_proto_depIdxs,
		EnumInfos:         file_acme_weather_v1_types_proto_enumTypes,
	}.Build()
	File_acme_weather_v1_types_proto = out.File
	file_acme_weather_v1_types_proto_rawDesc = nil
	file_acme_weather_v1_types_proto_goTypes = nil
	file_acme_weather_v1_types_proto_depIdxs = nil
}
