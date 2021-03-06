// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/overload/v3/overload.proto

package envoy_config_overload_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/struct"
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

type ResourceMonitor struct {
	// The name of the resource monitor to instantiate. Must match a registered
	// resource monitor type. The built-in resource monitors are:
	//
	// * :ref:`envoy.resource_monitors.fixed_heap
	//   <envoy_api_msg_config.resource_monitor.fixed_heap.v2alpha.FixedHeapConfig>`
	// * :ref:`envoy.resource_monitors.injected_resource
	//   <envoy_api_msg_config.resource_monitor.injected_resource.v2alpha.InjectedResourceConfig>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Configuration for the resource monitor being instantiated.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*ResourceMonitor_TypedConfig
	ConfigType           isResourceMonitor_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ResourceMonitor) Reset()         { *m = ResourceMonitor{} }
func (m *ResourceMonitor) String() string { return proto.CompactTextString(m) }
func (*ResourceMonitor) ProtoMessage()    {}
func (*ResourceMonitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{0}
}

func (m *ResourceMonitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceMonitor.Unmarshal(m, b)
}
func (m *ResourceMonitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceMonitor.Marshal(b, m, deterministic)
}
func (m *ResourceMonitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMonitor.Merge(m, src)
}
func (m *ResourceMonitor) XXX_Size() int {
	return xxx_messageInfo_ResourceMonitor.Size(m)
}
func (m *ResourceMonitor) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMonitor.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMonitor proto.InternalMessageInfo

func (m *ResourceMonitor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isResourceMonitor_ConfigType interface {
	isResourceMonitor_ConfigType()
}

type ResourceMonitor_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ResourceMonitor_TypedConfig) isResourceMonitor_ConfigType() {}

func (m *ResourceMonitor) GetConfigType() isResourceMonitor_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ResourceMonitor) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ResourceMonitor_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResourceMonitor) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResourceMonitor_TypedConfig)(nil),
	}
}

type ThresholdTrigger struct {
	// If the resource pressure is greater than or equal to this value, the trigger
	// will fire.
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThresholdTrigger) Reset()         { *m = ThresholdTrigger{} }
func (m *ThresholdTrigger) String() string { return proto.CompactTextString(m) }
func (*ThresholdTrigger) ProtoMessage()    {}
func (*ThresholdTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{1}
}

func (m *ThresholdTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThresholdTrigger.Unmarshal(m, b)
}
func (m *ThresholdTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThresholdTrigger.Marshal(b, m, deterministic)
}
func (m *ThresholdTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThresholdTrigger.Merge(m, src)
}
func (m *ThresholdTrigger) XXX_Size() int {
	return xxx_messageInfo_ThresholdTrigger.Size(m)
}
func (m *ThresholdTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_ThresholdTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_ThresholdTrigger proto.InternalMessageInfo

func (m *ThresholdTrigger) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Trigger struct {
	// The name of the resource this is a trigger for.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to TriggerOneof:
	//	*Trigger_Threshold
	TriggerOneof         isTrigger_TriggerOneof `protobuf_oneof:"trigger_oneof"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{2}
}

func (m *Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trigger.Unmarshal(m, b)
}
func (m *Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trigger.Marshal(b, m, deterministic)
}
func (m *Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trigger.Merge(m, src)
}
func (m *Trigger) XXX_Size() int {
	return xxx_messageInfo_Trigger.Size(m)
}
func (m *Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Trigger proto.InternalMessageInfo

func (m *Trigger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTrigger_TriggerOneof interface {
	isTrigger_TriggerOneof()
}

type Trigger_Threshold struct {
	Threshold *ThresholdTrigger `protobuf:"bytes,2,opt,name=threshold,proto3,oneof"`
}

func (*Trigger_Threshold) isTrigger_TriggerOneof() {}

func (m *Trigger) GetTriggerOneof() isTrigger_TriggerOneof {
	if m != nil {
		return m.TriggerOneof
	}
	return nil
}

func (m *Trigger) GetThreshold() *ThresholdTrigger {
	if x, ok := m.GetTriggerOneof().(*Trigger_Threshold); ok {
		return x.Threshold
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Trigger) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Trigger_Threshold)(nil),
	}
}

type OverloadAction struct {
	// The name of the overload action. This is just a well-known string that listeners can
	// use for registering callbacks. Custom overload actions should be named using reverse
	// DNS to ensure uniqueness.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A set of triggers for this action. If any of these triggers fire the overload action
	// is activated. Listeners are notified when the overload action transitions from
	// inactivated to activated, or vice versa.
	Triggers             []*Trigger `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OverloadAction) Reset()         { *m = OverloadAction{} }
func (m *OverloadAction) String() string { return proto.CompactTextString(m) }
func (*OverloadAction) ProtoMessage()    {}
func (*OverloadAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{3}
}

func (m *OverloadAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverloadAction.Unmarshal(m, b)
}
func (m *OverloadAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverloadAction.Marshal(b, m, deterministic)
}
func (m *OverloadAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverloadAction.Merge(m, src)
}
func (m *OverloadAction) XXX_Size() int {
	return xxx_messageInfo_OverloadAction.Size(m)
}
func (m *OverloadAction) XXX_DiscardUnknown() {
	xxx_messageInfo_OverloadAction.DiscardUnknown(m)
}

var xxx_messageInfo_OverloadAction proto.InternalMessageInfo

func (m *OverloadAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OverloadAction) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type OverloadManager struct {
	// The interval for refreshing resource usage.
	RefreshInterval *duration.Duration `protobuf:"bytes,1,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
	// The set of resources to monitor.
	ResourceMonitors []*ResourceMonitor `protobuf:"bytes,2,rep,name=resource_monitors,json=resourceMonitors,proto3" json:"resource_monitors,omitempty"`
	// The set of overload actions.
	Actions              []*OverloadAction `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OverloadManager) Reset()         { *m = OverloadManager{} }
func (m *OverloadManager) String() string { return proto.CompactTextString(m) }
func (*OverloadManager) ProtoMessage()    {}
func (*OverloadManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b32999822dfa297, []int{4}
}

func (m *OverloadManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverloadManager.Unmarshal(m, b)
}
func (m *OverloadManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverloadManager.Marshal(b, m, deterministic)
}
func (m *OverloadManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverloadManager.Merge(m, src)
}
func (m *OverloadManager) XXX_Size() int {
	return xxx_messageInfo_OverloadManager.Size(m)
}
func (m *OverloadManager) XXX_DiscardUnknown() {
	xxx_messageInfo_OverloadManager.DiscardUnknown(m)
}

var xxx_messageInfo_OverloadManager proto.InternalMessageInfo

func (m *OverloadManager) GetRefreshInterval() *duration.Duration {
	if m != nil {
		return m.RefreshInterval
	}
	return nil
}

func (m *OverloadManager) GetResourceMonitors() []*ResourceMonitor {
	if m != nil {
		return m.ResourceMonitors
	}
	return nil
}

func (m *OverloadManager) GetActions() []*OverloadAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceMonitor)(nil), "envoy.config.overload.v3.ResourceMonitor")
	proto.RegisterType((*ThresholdTrigger)(nil), "envoy.config.overload.v3.ThresholdTrigger")
	proto.RegisterType((*Trigger)(nil), "envoy.config.overload.v3.Trigger")
	proto.RegisterType((*OverloadAction)(nil), "envoy.config.overload.v3.OverloadAction")
	proto.RegisterType((*OverloadManager)(nil), "envoy.config.overload.v3.OverloadManager")
}

func init() {
	proto.RegisterFile("envoy/config/overload/v3/overload.proto", fileDescriptor_1b32999822dfa297)
}

var fileDescriptor_1b32999822dfa297 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0xad, 0xd3, 0xae, 0xed, 0x5c, 0x4a, 0x4b, 0x54, 0x69, 0xe9, 0x80, 0xa9, 0xab, 0x04, 0xeb,
	0xd0, 0x96, 0xa0, 0x16, 0x0e, 0xf4, 0x82, 0x1a, 0x26, 0x31, 0x26, 0x4d, 0x4c, 0xd1, 0xee, 0xc5,
	0x6b, 0xdc, 0x34, 0x52, 0x66, 0x57, 0x8e, 0x13, 0xad, 0x37, 0x8e, 0x9c, 0x39, 0xf2, 0x17, 0x38,
	0x22, 0x71, 0x40, 0x5c, 0x91, 0xb8, 0xf2, 0x4f, 0x38, 0xa2, 0x9e, 0x50, 0x6c, 0xa7, 0x53, 0x53,
	0xb5, 0xe0, 0x93, 0xe5, 0xf7, 0xde, 0xf7, 0xbd, 0xef, 0xe9, 0x33, 0x3c, 0xc0, 0x24, 0xa6, 0x33,
	0x6b, 0x44, 0xc9, 0xd8, 0xf7, 0x2c, 0x1a, 0x63, 0x16, 0x50, 0xe4, 0x5a, 0x71, 0x6f, 0x71, 0x37,
	0xa7, 0x8c, 0x72, 0xaa, 0x1b, 0x82, 0x68, 0x4a, 0xa2, 0xb9, 0x00, 0xe3, 0xde, 0x6e, 0xd3, 0xa3,
	0xd4, 0x0b, 0xb0, 0x25, 0x78, 0x57, 0xd1, 0xd8, 0x42, 0x64, 0x26, 0x45, 0xbb, 0x7b, 0x59, 0xc8,
	0x8d, 0x18, 0xe2, 0x3e, 0x25, 0x0a, 0x7f, 0x90, 0xc5, 0x43, 0xce, 0xa2, 0x11, 0x57, 0xe8, 0xc3,
	0xc8, 0x9d, 0x22, 0x0b, 0x11, 0x42, 0xb9, 0x10, 0x85, 0x56, 0xc8, 0x11, 0x8f, 0x42, 0x05, 0xef,
	0xaf, 0xc0, 0x31, 0x66, 0xa1, 0x4f, 0x89, 0x4f, 0x3c, 0x45, 0xd9, 0x89, 0x51, 0xe0, 0xbb, 0x88,
	0x63, 0x2b, 0xbd, 0x48, 0xa0, 0xfd, 0x1d, 0xc0, 0x9a, 0x83, 0x43, 0x1a, 0xb1, 0x11, 0x3e, 0xa7,
	0xc4, 0xe7, 0x94, 0xe9, 0xf7, 0x61, 0x81, 0xa0, 0x6b, 0x6c, 0x80, 0x16, 0xe8, 0x6c, 0xdb, 0xa5,
	0xb9, 0x5d, 0x60, 0x5a, 0x0b, 0x38, 0xe2, 0x51, 0x7f, 0x01, 0xef, 0xf0, 0xd9, 0x14, 0xbb, 0x43,
	0x19, 0x80, 0x91, 0x6f, 0x81, 0x4e, 0xa5, 0xdb, 0x30, 0xe5, 0x00, 0x66, 0x3a, 0x80, 0x39, 0x20,
	0xb3, 0xd3, 0x9c, 0x53, 0x11, 0xdc, 0x57, 0x82, 0xda, 0x7f, 0xf6, 0xe9, 0xc7, 0x87, 0x3d, 0x0b,
	0x1e, 0xaf, 0x09, 0xb0, 0x8b, 0x82, 0xe9, 0x04, 0x99, 0x19, 0x37, 0x76, 0x15, 0x56, 0x24, 0x75,
	0x98, 0xd4, 0x3a, 0x2b, 0x94, 0xb5, 0x7a, 0xde, 0x29, 0xca, 0xa7, 0xf6, 0x0d, 0xac, 0x5f, 0x4e,
	0x18, 0x0e, 0x27, 0x34, 0x70, 0x2f, 0x99, 0xef, 0x79, 0x98, 0xe9, 0xc7, 0x70, 0x2b, 0x46, 0x41,
	0x24, 0xfd, 0x03, 0x7b, 0x67, 0x6e, 0x37, 0x74, 0xbd, 0x99, 0x13, 0xe7, 0xf7, 0xcb, 0xc3, 0x9c,
	0x3a, 0x8e, 0x64, 0xf5, 0x9f, 0x27, 0xae, 0x9e, 0x42, 0x73, 0xb3, 0xab, 0x6c, 0x97, 0xf6, 0x57,
	0x00, 0x4b, 0x69, 0xc7, 0x8d, 0x81, 0x9d, 0xc1, 0x6d, 0x9e, 0x8a, 0x0d, 0x4d, 0xa4, 0xf5, 0x64,
	0x5d, 0xb3, 0xde, 0x4a, 0x9f, 0xd3, 0x9c, 0x73, 0x2b, 0xef, 0x1f, 0x25, 0x5e, 0x0f, 0xe0, 0xa3,
	0x7f, 0x78, 0x95, 0x52, 0xbb, 0x01, 0xab, 0x5c, 0x5e, 0x87, 0x94, 0x60, 0x3a, 0xd6, 0xf3, 0x7f,
	0x6c, 0xd0, 0xfe, 0x0c, 0xe0, 0xdd, 0xb7, 0x4a, 0x32, 0x18, 0x25, 0xfb, 0xb2, 0xd9, 0xff, 0x6b,
	0x58, 0x56, 0x55, 0x42, 0x43, 0x6b, 0xe5, 0x3b, 0x95, 0xee, 0xfe, 0x06, 0xfb, 0xaa, 0x75, 0x79,
	0x6e, 0x6f, 0x7d, 0x04, 0x5a, 0x19, 0x38, 0x0b, 0x71, 0xbf, 0x97, 0x98, 0x37, 0xe1, 0xd1, 0x66,
	0xf3, 0xcb, 0xd6, 0xda, 0x5f, 0x34, 0x58, 0x4b, 0x9f, 0xce, 0x11, 0x41, 0x49, 0xdc, 0x27, 0xb0,
	0xce, 0xf0, 0x38, 0xc9, 0x64, 0xe8, 0x13, 0x8e, 0x59, 0x8c, 0x02, 0x61, 0xbd, 0xd2, 0x6d, 0xae,
	0xac, 0xe1, 0x89, 0xfa, 0x67, 0x4e, 0x4d, 0x49, 0xde, 0x28, 0x85, 0xfe, 0x0e, 0xde, 0x63, 0x6a,
	0xd5, 0x86, 0xd7, 0x72, 0xd7, 0xd2, 0x01, 0x0f, 0xd7, 0x0f, 0x98, 0xdd, 0xce, 0xdb, 0x41, 0xeb,
	0x6c, 0x19, 0x0a, 0x75, 0x1b, 0x96, 0x90, 0x98, 0x22, 0x34, 0xf2, 0xa2, 0x6e, 0x67, 0x7d, 0xdd,
	0xe5, 0xb1, 0x9d, 0x54, 0xf8, 0x9f, 0x7f, 0x26, 0x93, 0x90, 0x3d, 0xf8, 0xf6, 0xfe, 0xe7, 0xaf,
	0xa2, 0x56, 0xd7, 0xe0, 0x63, 0x9f, 0xca, 0xa6, 0x53, 0x46, 0x6f, 0x66, 0x6b, 0xfb, 0xdb, 0xd5,
	0xb4, 0xc4, 0x45, 0x92, 0xdc, 0x05, 0xb8, 0x2a, 0x8a, 0x08, 0x7b, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0xdf, 0x68, 0xfd, 0x18, 0x05, 0x00, 0x00,
}
