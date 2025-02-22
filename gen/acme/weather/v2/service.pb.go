// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: acme/weather/v2/service.proto

package weatherv2

import (
	shared "github.com/sudorandom/reflect-bug/gen/acme/weather/shared"
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

type GetWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GetWeatherRequest) Reset() {
	*x = GetWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_weather_v2_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherRequest) ProtoMessage() {}

func (x *GetWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acme_weather_v2_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherRequest.ProtoReflect.Descriptor instead.
func (*GetWeatherRequest) Descriptor() ([]byte, []int) {
	return file_acme_weather_v2_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetWeatherRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GetWeatherRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GetWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float32          `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Condition   shared.Condition `protobuf:"varint,2,opt,name=condition,proto3,enum=acme.weather.shared.Condition" json:"condition,omitempty"`
}

func (x *GetWeatherResponse) Reset() {
	*x = GetWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_weather_v2_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherResponse) ProtoMessage() {}

func (x *GetWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acme_weather_v2_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherResponse.ProtoReflect.Descriptor instead.
func (*GetWeatherResponse) Descriptor() ([]byte, []int) {
	return file_acme_weather_v2_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetWeatherResponse) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *GetWeatherResponse) GetCondition() shared.Condition {
	if x != nil {
		return x.Condition
	}
	return shared.Condition(0)
}

var File_acme_weather_v2_service_proto protoreflect.FileDescriptor

var file_acme_weather_v2_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x1a, 0x1f, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0x74, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x63,
	0x6d, 0x65, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x67, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x63, 0x6d,
	0x65, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0xc2, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x64, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2d, 0x62, 0x75, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61,
	0x63, 0x6d, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x3b, 0x77,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x41, 0x57, 0x58, 0xaa, 0x02,
	0x0f, 0x41, 0x63, 0x6d, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x32,
	0xca, 0x02, 0x0f, 0x41, 0x63, 0x6d, 0x65, 0x5c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5c,
	0x56, 0x32, 0xe2, 0x02, 0x1b, 0x41, 0x63, 0x6d, 0x65, 0x5c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x11, 0x41, 0x63, 0x6d, 0x65, 0x3a, 0x3a, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acme_weather_v2_service_proto_rawDescOnce sync.Once
	file_acme_weather_v2_service_proto_rawDescData = file_acme_weather_v2_service_proto_rawDesc
)

func file_acme_weather_v2_service_proto_rawDescGZIP() []byte {
	file_acme_weather_v2_service_proto_rawDescOnce.Do(func() {
		file_acme_weather_v2_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_acme_weather_v2_service_proto_rawDescData)
	})
	return file_acme_weather_v2_service_proto_rawDescData
}

var file_acme_weather_v2_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_acme_weather_v2_service_proto_goTypes = []any{
	(*GetWeatherRequest)(nil),  // 0: acme.weather.v2.GetWeatherRequest
	(*GetWeatherResponse)(nil), // 1: acme.weather.v2.GetWeatherResponse
	(shared.Condition)(0),      // 2: acme.weather.shared.Condition
}
var file_acme_weather_v2_service_proto_depIdxs = []int32{
	2, // 0: acme.weather.v2.GetWeatherResponse.condition:type_name -> acme.weather.shared.Condition
	0, // 1: acme.weather.v2.WeatherService.GetWeather:input_type -> acme.weather.v2.GetWeatherRequest
	1, // 2: acme.weather.v2.WeatherService.GetWeather:output_type -> acme.weather.v2.GetWeatherResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_acme_weather_v2_service_proto_init() }
func file_acme_weather_v2_service_proto_init() {
	if File_acme_weather_v2_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acme_weather_v2_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetWeatherRequest); i {
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
		file_acme_weather_v2_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetWeatherResponse); i {
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
			RawDescriptor: file_acme_weather_v2_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_acme_weather_v2_service_proto_goTypes,
		DependencyIndexes: file_acme_weather_v2_service_proto_depIdxs,
		MessageInfos:      file_acme_weather_v2_service_proto_msgTypes,
	}.Build()
	File_acme_weather_v2_service_proto = out.File
	file_acme_weather_v2_service_proto_rawDesc = nil
	file_acme_weather_v2_service_proto_goTypes = nil
	file_acme_weather_v2_service_proto_depIdxs = nil
}
