// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: ch_table.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// If set will output the schema file order based on the provided value.
type ClickHouseMessageOptions_FieldOrder int32

const (
	ClickHouseMessageOptions_FIELD_ORDER_UNSPECIFIED ClickHouseMessageOptions_FieldOrder = 0
	ClickHouseMessageOptions_FIELD_ORDER_BY_NUMBER   ClickHouseMessageOptions_FieldOrder = 1
)

// Enum value maps for ClickHouseMessageOptions_FieldOrder.
var (
	ClickHouseMessageOptions_FieldOrder_name = map[int32]string{
		0: "FIELD_ORDER_UNSPECIFIED",
		1: "FIELD_ORDER_BY_NUMBER",
	}
	ClickHouseMessageOptions_FieldOrder_value = map[string]int32{
		"FIELD_ORDER_UNSPECIFIED": 0,
		"FIELD_ORDER_BY_NUMBER":   1,
	}
)

func (x ClickHouseMessageOptions_FieldOrder) Enum() *ClickHouseMessageOptions_FieldOrder {
	p := new(ClickHouseMessageOptions_FieldOrder)
	*p = x
	return p
}

func (x ClickHouseMessageOptions_FieldOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClickHouseMessageOptions_FieldOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_ch_table_proto_enumTypes[0].Descriptor()
}

func (ClickHouseMessageOptions_FieldOrder) Type() protoreflect.EnumType {
	return &file_ch_table_proto_enumTypes[0]
}

func (x ClickHouseMessageOptions_FieldOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClickHouseMessageOptions_FieldOrder.Descriptor instead.
func (ClickHouseMessageOptions_FieldOrder) EnumDescriptor() ([]byte, []int) {
	return file_ch_table_proto_rawDescGZIP(), []int{0, 0}
}

// 使用的表引擎 默认使用 MergeTree
type ClickHouseMessageOptions_TableEngine int32

const (
	ClickHouseMessageOptions_TABLE_ENGINE_UNSPECIFIED ClickHouseMessageOptions_TableEngine = 0
	ClickHouseMessageOptions_MERGE_TREE               ClickHouseMessageOptions_TableEngine = 1
	ClickHouseMessageOptions_REPLACING_MERGE_TREE     ClickHouseMessageOptions_TableEngine = 2
	ClickHouseMessageOptions_SUMMING_MERGE_TREE       ClickHouseMessageOptions_TableEngine = 3
)

// Enum value maps for ClickHouseMessageOptions_TableEngine.
var (
	ClickHouseMessageOptions_TableEngine_name = map[int32]string{
		0: "TABLE_ENGINE_UNSPECIFIED",
		1: "MERGE_TREE",
		2: "REPLACING_MERGE_TREE",
		3: "SUMMING_MERGE_TREE",
	}
	ClickHouseMessageOptions_TableEngine_value = map[string]int32{
		"TABLE_ENGINE_UNSPECIFIED": 0,
		"MERGE_TREE":               1,
		"REPLACING_MERGE_TREE":     2,
		"SUMMING_MERGE_TREE":       3,
	}
)

func (x ClickHouseMessageOptions_TableEngine) Enum() *ClickHouseMessageOptions_TableEngine {
	p := new(ClickHouseMessageOptions_TableEngine)
	*p = x
	return p
}

func (x ClickHouseMessageOptions_TableEngine) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClickHouseMessageOptions_TableEngine) Descriptor() protoreflect.EnumDescriptor {
	return file_ch_table_proto_enumTypes[1].Descriptor()
}

func (ClickHouseMessageOptions_TableEngine) Type() protoreflect.EnumType {
	return &file_ch_table_proto_enumTypes[1]
}

func (x ClickHouseMessageOptions_TableEngine) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClickHouseMessageOptions_TableEngine.Descriptor instead.
func (ClickHouseMessageOptions_TableEngine) EnumDescriptor() ([]byte, []int) {
	return file_ch_table_proto_rawDescGZIP(), []int{0, 1}
}

// 数据分区设置 一般按时间分区 月/天/小时
// 默认不分区 大部分情况按月分区即可满足需求
type ClickHouseMessageOptions_TIME_PARTITION int32

const (
	ClickHouseMessageOptions_TIME_PARTITION_UNSPECIFIED ClickHouseMessageOptions_TIME_PARTITION = 0
	ClickHouseMessageOptions_MONTH                      ClickHouseMessageOptions_TIME_PARTITION = 1
	ClickHouseMessageOptions_DAY                        ClickHouseMessageOptions_TIME_PARTITION = 2
	ClickHouseMessageOptions_HOUR                       ClickHouseMessageOptions_TIME_PARTITION = 3
)

// Enum value maps for ClickHouseMessageOptions_TIME_PARTITION.
var (
	ClickHouseMessageOptions_TIME_PARTITION_name = map[int32]string{
		0: "TIME_PARTITION_UNSPECIFIED",
		1: "MONTH",
		2: "DAY",
		3: "HOUR",
	}
	ClickHouseMessageOptions_TIME_PARTITION_value = map[string]int32{
		"TIME_PARTITION_UNSPECIFIED": 0,
		"MONTH":                      1,
		"DAY":                        2,
		"HOUR":                       3,
	}
)

func (x ClickHouseMessageOptions_TIME_PARTITION) Enum() *ClickHouseMessageOptions_TIME_PARTITION {
	p := new(ClickHouseMessageOptions_TIME_PARTITION)
	*p = x
	return p
}

func (x ClickHouseMessageOptions_TIME_PARTITION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClickHouseMessageOptions_TIME_PARTITION) Descriptor() protoreflect.EnumDescriptor {
	return file_ch_table_proto_enumTypes[2].Descriptor()
}

func (ClickHouseMessageOptions_TIME_PARTITION) Type() protoreflect.EnumType {
	return &file_ch_table_proto_enumTypes[2]
}

func (x ClickHouseMessageOptions_TIME_PARTITION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClickHouseMessageOptions_TIME_PARTITION.Descriptor instead.
func (ClickHouseMessageOptions_TIME_PARTITION) EnumDescriptor() ([]byte, []int) {
	return file_ch_table_proto_rawDescGZIP(), []int{0, 2}
}

type ClickHouseMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies tabel in ClickHouse for the message.
	// table_name should be db.table format
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// If true, ClickHouse field names will default to a field's JSON name,
	// not its original/proto field name.
	UseJsonNames bool `protobuf:"varint,2,opt,name=use_json_names,json=useJsonNames,proto3" json:"use_json_names,omitempty"`
	// If set, adds defined extra fields to a JSON representation of the message.
	// Value format: "<field name>:<ClickHouse field type>" for basic types
	// or "<field name>:Nested:<protobuf type>" for message types.
	// "NULLABLE" by default, different mode may be set via optional suffix ":<mode>"
	ExtraFields      []string                             `protobuf:"bytes,3,rep,name=extra_fields,json=extraFields,proto3" json:"extra_fields,omitempty"`
	OutputFieldOrder ClickHouseMessageOptions_FieldOrder  `protobuf:"varint,4,opt,name=output_field_order,json=outputFieldOrder,proto3,enum=gen_ch_schema.ClickHouseMessageOptions_FieldOrder" json:"output_field_order,omitempty"`
	TableEngine      ClickHouseMessageOptions_TableEngine `protobuf:"varint,6,opt,name=table_engine,json=tableEngine,proto3,enum=gen_ch_schema.ClickHouseMessageOptions_TableEngine" json:"table_engine,omitempty"`
	// 表的排序键 可以大幅提高聚合查询的性能
	// 排序依照click house官方文档定义
	// 可以使用复杂表达式 例如 "toYYYYMMDD(ctime),id"
	// default "ctime"
	OrderBy       string                                  `protobuf:"bytes,7,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	TimePartition ClickHouseMessageOptions_TIME_PARTITION `protobuf:"varint,8,opt,name=time_partition,json=timePartition,proto3,enum=gen_ch_schema.ClickHouseMessageOptions_TIME_PARTITION" json:"time_partition,omitempty"`
	// 数据过期设置 指定数据过期时间 过期后数据会被删除
	// 支持月/周/天/小时定义 例如: 3m 2w 1d 6h
	// y: 年，q: 季度，m: 月，w: 周，d: 天，h: 小时，M: 分钟, s: 秒
	// 默认不设置过期时间
	// todo: 后续可能提供过期之后的操作 例如转存到文件或者条件删除
	Ttl      string                              `protobuf:"bytes,9,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Settings []*ClickHouseMessageOptions_Setting `protobuf:"bytes,10,rep,name=settings,proto3" json:"settings,omitempty"`
	// RepalcingMergeTree的版本列
	// 用于去重 保留版本列最大的数据
	// 默认为ctime
	ReplacingVersionCol string `protobuf:"bytes,11,opt,name=replacing_version_col,json=replacingVersionCol,proto3" json:"replacing_version_col,omitempty"`
	// RepalcingMergeTree的删除状态列
	// 用于标记数据是否待删除
	// 默认为is_deleted
	ReplacingDeletedCol string `protobuf:"bytes,12,opt,name=replacing_deleted_col,json=replacingDeletedCol,proto3" json:"replacing_deleted_col,omitempty"`
	// SummingMergeTree的聚合列
	// 插入时，聚合列的值会累加
	// 默认为空 代表所有数字类型的字段都会参与聚合
	SummingAggregateCols []string `protobuf:"bytes,13,rep,name=summing_aggregate_cols,json=summingAggregateCols,proto3" json:"summing_aggregate_cols,omitempty"`
}

func (x *ClickHouseMessageOptions) Reset() {
	*x = ClickHouseMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ch_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickHouseMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickHouseMessageOptions) ProtoMessage() {}

func (x *ClickHouseMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ch_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickHouseMessageOptions.ProtoReflect.Descriptor instead.
func (*ClickHouseMessageOptions) Descriptor() ([]byte, []int) {
	return file_ch_table_proto_rawDescGZIP(), []int{0}
}

func (x *ClickHouseMessageOptions) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *ClickHouseMessageOptions) GetUseJsonNames() bool {
	if x != nil {
		return x.UseJsonNames
	}
	return false
}

func (x *ClickHouseMessageOptions) GetExtraFields() []string {
	if x != nil {
		return x.ExtraFields
	}
	return nil
}

func (x *ClickHouseMessageOptions) GetOutputFieldOrder() ClickHouseMessageOptions_FieldOrder {
	if x != nil {
		return x.OutputFieldOrder
	}
	return ClickHouseMessageOptions_FIELD_ORDER_UNSPECIFIED
}

func (x *ClickHouseMessageOptions) GetTableEngine() ClickHouseMessageOptions_TableEngine {
	if x != nil {
		return x.TableEngine
	}
	return ClickHouseMessageOptions_TABLE_ENGINE_UNSPECIFIED
}

func (x *ClickHouseMessageOptions) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ClickHouseMessageOptions) GetTimePartition() ClickHouseMessageOptions_TIME_PARTITION {
	if x != nil {
		return x.TimePartition
	}
	return ClickHouseMessageOptions_TIME_PARTITION_UNSPECIFIED
}

func (x *ClickHouseMessageOptions) GetTtl() string {
	if x != nil {
		return x.Ttl
	}
	return ""
}

func (x *ClickHouseMessageOptions) GetSettings() []*ClickHouseMessageOptions_Setting {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *ClickHouseMessageOptions) GetReplacingVersionCol() string {
	if x != nil {
		return x.ReplacingVersionCol
	}
	return ""
}

func (x *ClickHouseMessageOptions) GetReplacingDeletedCol() string {
	if x != nil {
		return x.ReplacingDeletedCol
	}
	return ""
}

func (x *ClickHouseMessageOptions) GetSummingAggregateCols() []string {
	if x != nil {
		return x.SummingAggregateCols
	}
	return nil
}

// settings 配置键 可以自定义配置
type ClickHouseMessageOptions_Setting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ClickHouseMessageOptions_Setting) Reset() {
	*x = ClickHouseMessageOptions_Setting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ch_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickHouseMessageOptions_Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickHouseMessageOptions_Setting) ProtoMessage() {}

func (x *ClickHouseMessageOptions_Setting) ProtoReflect() protoreflect.Message {
	mi := &file_ch_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickHouseMessageOptions_Setting.ProtoReflect.Descriptor instead.
func (*ClickHouseMessageOptions_Setting) Descriptor() ([]byte, []int) {
	return file_ch_table_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClickHouseMessageOptions_Setting) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClickHouseMessageOptions_Setting) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var file_ch_table_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ClickHouseMessageOptions)(nil),
		Field:         18123,
		Name:          "gen_ch_schema.clickhouse_opts",
		Tag:           "bytes,18123,opt,name=clickhouse_opts",
		Filename:      "ch_table.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// ClickHouse message schema generation options.
	//
	// The field number is a globally unique id for this option, assigned by
	// protobuf-global-extension-registry@google.com
	//
	// optional gen_ch_schema.ClickHouseMessageOptions clickhouse_opts = 18123;
	E_ClickhouseOpts = &file_ch_table_proto_extTypes[0]
)

var File_ch_table_proto protoreflect.FileDescriptor

var file_ch_table_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x67, 0x65, 0x6e, 0x5f, 0x63, 0x68, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xed, 0x07, 0x0a, 0x18, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x75, 0x73, 0x65, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x60, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x65, 0x6e, 0x5f, 0x63, 0x68, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33,
	0x2e, 0x67, 0x65, 0x6e, 0x5f, 0x63, 0x68, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x5d, 0x0a, 0x0e, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x65, 0x6e, 0x5f, 0x63, 0x68, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x49, 0x4d,
	0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x52, 0x0d, 0x74, 0x69, 0x6d,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x4b, 0x0a, 0x08,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x67, 0x65, 0x6e, 0x5f, 0x63, 0x68, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x12, 0x32, 0x0a,
	0x15, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x6c, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x75, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x14, 0x73, 0x75, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x73, 0x1a, 0x33, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x44, 0x0a, 0x0a,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x01, 0x22, 0x6d, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x47, 0x49, 0x4e,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x5f, 0x54, 0x52, 0x45, 0x45, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x52,
	0x47, 0x45, 0x5f, 0x54, 0x52, 0x45, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x55, 0x4d,
	0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x5f, 0x54, 0x52, 0x45, 0x45, 0x10,
	0x03, 0x22, 0x4e, 0x0a, 0x0e, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52, 0x10,
	0x03, 0x3a, 0x73, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcb, 0x8d, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x67, 0x65, 0x6e, 0x5f, 0x63, 0x68, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6c,
	0x69, 0x63, 0x6b, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x64, 0x78, 0x68, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x63, 0x68, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ch_table_proto_rawDescOnce sync.Once
	file_ch_table_proto_rawDescData = file_ch_table_proto_rawDesc
)

func file_ch_table_proto_rawDescGZIP() []byte {
	file_ch_table_proto_rawDescOnce.Do(func() {
		file_ch_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_ch_table_proto_rawDescData)
	})
	return file_ch_table_proto_rawDescData
}

var file_ch_table_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_ch_table_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ch_table_proto_goTypes = []interface{}{
	(ClickHouseMessageOptions_FieldOrder)(0),     // 0: gen_ch_schema.ClickHouseMessageOptions.FieldOrder
	(ClickHouseMessageOptions_TableEngine)(0),    // 1: gen_ch_schema.ClickHouseMessageOptions.TableEngine
	(ClickHouseMessageOptions_TIME_PARTITION)(0), // 2: gen_ch_schema.ClickHouseMessageOptions.TIME_PARTITION
	(*ClickHouseMessageOptions)(nil),             // 3: gen_ch_schema.ClickHouseMessageOptions
	(*ClickHouseMessageOptions_Setting)(nil),     // 4: gen_ch_schema.ClickHouseMessageOptions.Setting
	(*descriptorpb.MessageOptions)(nil),          // 5: google.protobuf.MessageOptions
}
var file_ch_table_proto_depIdxs = []int32{
	0, // 0: gen_ch_schema.ClickHouseMessageOptions.output_field_order:type_name -> gen_ch_schema.ClickHouseMessageOptions.FieldOrder
	1, // 1: gen_ch_schema.ClickHouseMessageOptions.table_engine:type_name -> gen_ch_schema.ClickHouseMessageOptions.TableEngine
	2, // 2: gen_ch_schema.ClickHouseMessageOptions.time_partition:type_name -> gen_ch_schema.ClickHouseMessageOptions.TIME_PARTITION
	4, // 3: gen_ch_schema.ClickHouseMessageOptions.settings:type_name -> gen_ch_schema.ClickHouseMessageOptions.Setting
	5, // 4: gen_ch_schema.clickhouse_opts:extendee -> google.protobuf.MessageOptions
	3, // 5: gen_ch_schema.clickhouse_opts:type_name -> gen_ch_schema.ClickHouseMessageOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	5, // [5:6] is the sub-list for extension type_name
	4, // [4:5] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ch_table_proto_init() }
func file_ch_table_proto_init() {
	if File_ch_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ch_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickHouseMessageOptions); i {
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
		file_ch_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickHouseMessageOptions_Setting); i {
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
			RawDescriptor: file_ch_table_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_ch_table_proto_goTypes,
		DependencyIndexes: file_ch_table_proto_depIdxs,
		EnumInfos:         file_ch_table_proto_enumTypes,
		MessageInfos:      file_ch_table_proto_msgTypes,
		ExtensionInfos:    file_ch_table_proto_extTypes,
	}.Build()
	File_ch_table_proto = out.File
	file_ch_table_proto_rawDesc = nil
	file_ch_table_proto_goTypes = nil
	file_ch_table_proto_depIdxs = nil
}