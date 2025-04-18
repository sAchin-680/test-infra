// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protocol/ksandbox/v1alpha1/agent.proto

package protocol

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

// Requests execution of a shell command
type ExecuteCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The command to execute
	Command []string `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`
	// The working directory for the command
	WorkingDir string `protobuf:"bytes,2,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	// The environment variables to set for the command
	Env []*EnvironmentVariable `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
}

func (x *ExecuteCommandRequest) Reset() {
	*x = ExecuteCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCommandRequest) ProtoMessage() {}

func (x *ExecuteCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCommandRequest.ProtoReflect.Descriptor instead.
func (*ExecuteCommandRequest) Descriptor() ([]byte, []int) {
	return file_protocol_ksandbox_v1alpha1_agent_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteCommandRequest) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *ExecuteCommandRequest) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

func (x *ExecuteCommandRequest) GetEnv() []*EnvironmentVariable {
	if x != nil {
		return x.Env
	}
	return nil
}

type EnvironmentVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EnvironmentVariable) Reset() {
	*x = EnvironmentVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentVariable) ProtoMessage() {}

func (x *EnvironmentVariable) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentVariable.ProtoReflect.Descriptor instead.
func (*EnvironmentVariable) Descriptor() ([]byte, []int) {
	return file_protocol_ksandbox_v1alpha1_agent_proto_rawDescGZIP(), []int{1}
}

func (x *EnvironmentVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvironmentVariable) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Reports results of executing a shell command
type ExecuteCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode int32  `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	Stdout   []byte `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr   []byte `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (x *ExecuteCommandResponse) Reset() {
	*x = ExecuteCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCommandResponse) ProtoMessage() {}

func (x *ExecuteCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCommandResponse.ProtoReflect.Descriptor instead.
func (*ExecuteCommandResponse) Descriptor() ([]byte, []int) {
	return file_protocol_ksandbox_v1alpha1_agent_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteCommandResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *ExecuteCommandResponse) GetStdout() []byte {
	if x != nil {
		return x.Stdout
	}
	return nil
}

func (x *ExecuteCommandResponse) GetStderr() []byte {
	if x != nil {
		return x.Stderr
	}
	return nil
}

var File_protocol_ksandbox_v1alpha1_agent_proto protoreflect.FileDescriptor

var file_protocol_ksandbox_v1alpha1_agent_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6b, 0x73, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6b, 0x73, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x78, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x8c, 0x01, 0x0a, 0x15,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72,
	0x12, 0x38, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6b, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0x3f, 0x0a, 0x13, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x65, 0x0a, 0x16, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x64, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65,
	0x72, 0x72, 0x32, 0x6e, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x65, 0x0a, 0x0e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x28, 0x2e,
	0x6b, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6b, 0x73, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x78, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_ksandbox_v1alpha1_agent_proto_rawDescOnce sync.Once
	file_protocol_ksandbox_v1alpha1_agent_proto_rawDescData = file_protocol_ksandbox_v1alpha1_agent_proto_rawDesc
)

func file_protocol_ksandbox_v1alpha1_agent_proto_rawDescGZIP() []byte {
	file_protocol_ksandbox_v1alpha1_agent_proto_rawDescOnce.Do(func() {
		file_protocol_ksandbox_v1alpha1_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_ksandbox_v1alpha1_agent_proto_rawDescData)
	})
	return file_protocol_ksandbox_v1alpha1_agent_proto_rawDescData
}

var file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocol_ksandbox_v1alpha1_agent_proto_goTypes = []interface{}{
	(*ExecuteCommandRequest)(nil),  // 0: ksandbox.v1alpha1.ExecuteCommandRequest
	(*EnvironmentVariable)(nil),    // 1: ksandbox.v1alpha1.EnvironmentVariable
	(*ExecuteCommandResponse)(nil), // 2: ksandbox.v1alpha1.ExecuteCommandResponse
}
var file_protocol_ksandbox_v1alpha1_agent_proto_depIdxs = []int32{
	1, // 0: ksandbox.v1alpha1.ExecuteCommandRequest.env:type_name -> ksandbox.v1alpha1.EnvironmentVariable
	0, // 1: ksandbox.v1alpha1.Agent.ExecuteCommand:input_type -> ksandbox.v1alpha1.ExecuteCommandRequest
	2, // 2: ksandbox.v1alpha1.Agent.ExecuteCommand:output_type -> ksandbox.v1alpha1.ExecuteCommandResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_ksandbox_v1alpha1_agent_proto_init() }
func file_protocol_ksandbox_v1alpha1_agent_proto_init() {
	if File_protocol_ksandbox_v1alpha1_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCommandRequest); i {
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
		file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentVariable); i {
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
		file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCommandResponse); i {
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
			RawDescriptor: file_protocol_ksandbox_v1alpha1_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_ksandbox_v1alpha1_agent_proto_goTypes,
		DependencyIndexes: file_protocol_ksandbox_v1alpha1_agent_proto_depIdxs,
		MessageInfos:      file_protocol_ksandbox_v1alpha1_agent_proto_msgTypes,
	}.Build()
	File_protocol_ksandbox_v1alpha1_agent_proto = out.File
	file_protocol_ksandbox_v1alpha1_agent_proto_rawDesc = nil
	file_protocol_ksandbox_v1alpha1_agent_proto_goTypes = nil
	file_protocol_ksandbox_v1alpha1_agent_proto_depIdxs = nil
}
