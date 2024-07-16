package converter

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/prototext"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

func sqlPre(sql string) string {
	sql = strings.ReplaceAll(sql, "\n", "")
	sql = strings.ReplaceAll(sql, "\t", "")
	sql = strings.ReplaceAll(sql, " ", "")
	return sql
}

func testConvert(t *testing.T, input string, expectedOutputs map[string]string, extras ...func(request *plugin.CodeGeneratorRequest)) {
	req := plugin.CodeGeneratorRequest{}
	if err := prototext.Unmarshal([]byte(input), &req); err != nil {
		t.Fatal("Failed to parse test input: ", err)
	}

	// apply custom transformations, if any
	for _, extra := range extras {
		extra(&req)
	}

	expectedSchema := make(map[string]string)
	for filename, data := range expectedOutputs {
		expectedSchema[filename] = data
	}

	res, err := Convert(&req)
	if err != nil {
		t.Fatal("Conversion failed. ", err)
	}
	if res.Error != nil {
		t.Fatal("Conversion failed. ", res.Error)
	}

	actualSchema := make(map[string]string)
	for _, file := range res.GetFile() {
		actualSchema[file.GetName()] = file.GetContent()
	}

	if len(actualSchema) != len(expectedSchema) {
		t.Errorf("Expected %d files generated, but actually %d files:\nExpectation: %s\n Actual: %s",
			len(expectedSchema), len(actualSchema), expectedSchema, actualSchema)
	}

	for name, actual := range actualSchema {
		expected, ok := expectedSchema[name]
		if !ok {
			t.Error("Unexpected file generated: ", name)
		}
		if sqlPre(expected) != sqlPre(actual) {
			t.Errorf("Expected the content of %s to be \"%v\" but got \"%v\"", name, expected, actual)
		}
	}
}

// TestSimple tries a simple code generator request.
func TestSimple(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_ch_schema.clickhouse_opts] <table_name: "foo_table"> >
				>
			>
		`,
		map[string]string{
			"foo_table.ch.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\ti1 Int32,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

// TestIgnoreNonTargetMessage checks if the generator ignores messages without gen_ch_schema.table_name option.
func TestIgnoreNonTargetMessage(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
				>
				message_type <
					name: "BarProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_ch_schema.clickhouse_opts] <table_name: "bar_table"> >
				>
				message_type <
					name: "BazProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
				>
			>
		`,
		map[string]string{
			"bar_table.ch.sql": "CREATE TABLE IF NOT EXISTS bar_table\n(\n\ti1 Int32,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

// TestIgnoreNonTargetFile checks if the generator ignores messages in non target files.
func TestIgnoreNonTargetFile(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_ch_schema.clickhouse_opts] <table_name: "foo_table"> >
				>
			>
			proto_file <
				name: "bar.proto"
				package: "example_package.nested"
				message_type <
					name: "BarProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_ch_schema.clickhouse_opts] <table_name: "bar_table"> >
				>
			>
		`,
		map[string]string{
			"foo_table.ch.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\ti1 Int32,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

// TestStopsAtRecursiveMessage verifies that generator ignores nested fields if finds message is recursive.
// Proceeding in such case without limit would cause infinite recursion.
func TestStopsAtRecursiveMessage(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.recursive"
				message_type <
					name: "FooProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field <
                        name: "bar" number: 2 type: TYPE_MESSAGE label: LABEL_OPTIONAL
                        type_name: "BarProto" >
					options < [gen_ch_schema.clickhouse_opts] <table_name: "foo_table"> >
				>
				message_type <
					name: "BarProto"
					field < name: "i2" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field <
                        name: "foo" number: 2 type: TYPE_MESSAGE label: LABEL_OPTIONAL
                        type_name: "FooProto" >
				>
			>
		`,
		map[string]string{
			"foo_table.ch.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\ti1 Int32,\n\tbar_i2 Int32,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

// TestTypes tests the generator with various field types
func TestTypes(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "i32" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i64" number: 2 type: TYPE_INT64 label: LABEL_OPTIONAL >
					field < name: "ui32" number: 3 type: TYPE_UINT32 label: LABEL_OPTIONAL >
					field < name: "ui64" number: 4 type: TYPE_UINT64 label: LABEL_OPTIONAL >
					field < name: "si32" number: 5 type: TYPE_SINT32 label: LABEL_OPTIONAL >
					field < name: "si64" number: 6 type: TYPE_SINT64 label: LABEL_OPTIONAL >
					field < name: "ufi32" number: 7 type: TYPE_FIXED32 label: LABEL_OPTIONAL >
					field < name: "ufi64" number: 8 type: TYPE_FIXED64 label: LABEL_OPTIONAL >
					field < name: "sfi32" number: 9 type: TYPE_SFIXED32 label: LABEL_OPTIONAL >
					field < name: "sfi64" number: 10 type: TYPE_SFIXED64 label: LABEL_OPTIONAL >
					field < name: "d" number: 11 type: TYPE_DOUBLE label: LABEL_OPTIONAL >
					field < name: "f" number: 12 type: TYPE_FLOAT label: LABEL_OPTIONAL >
					field < name: "bool" number: 16 type: TYPE_BOOL label: LABEL_OPTIONAL >
					field < name: "str" number: 13 type: TYPE_STRING label: LABEL_OPTIONAL >
					field < name: "bytes" number: 14 type: TYPE_BYTES label: LABEL_OPTIONAL >
					field <
						name: "enum1" number: 15 type: TYPE_ENUM label: LABEL_OPTIONAL
						type_name: ".example_package.nested.FooProto.Enum1"
					>
					field <
						name: "enum2" number: 16 type: TYPE_ENUM label: LABEL_OPTIONAL
						type_name: "FooProto.Enum1"
					>
					field <
						name: "grp1" number: 17 type: TYPE_GROUP label: LABEL_OPTIONAL
						type_name: ".example_package.nested.FooProto.Group1"
					>
					field <
						name: "grp2" number: 18 type: TYPE_GROUP label: LABEL_OPTIONAL
						type_name: "FooProto.Group1"
					>
					field <
						name: "msg1" number: 19 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".example_package.nested.FooProto.Nested1"
					>
					field <
						name: "msg2" number: 20 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: "FooProto.Nested1"
					>
					field <
						name: "msg3" number: 21 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".example_package.nested2.BarProto"
					>
					field <
						name: "msg4" number: 22 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: "nested2.BarProto"
					>
					field <
						name: "msg2" number: 23 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: "FooProto.EmptyNested1"
					>
					nested_type <
						name: "Group1"
						field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					>
					nested_type <
						name: "Nested1"
						field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					>
					nested_type <
						name: "EmptyNested1"
					>
					enum_type < name: "Enum1" value < name: "E1" number: 1 > value < name: "E2" number: 2 > >
					options < [gen_ch_schema.clickhouse_opts] <table_name: "foo_table"> >
				>
			>
			proto_file <
				name: "bar.proto"
				package: "example_package.nested2"
				message_type <
					name: "BarProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i2" number: 2 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i3" number: 3 type: TYPE_INT32 label: LABEL_OPTIONAL >
				>
			>
		`,
		map[string]string{
			"foo_table.ch.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\ti32 Int32,\n\ti64 Int64,\n\tui32 UInt32,\n\tui64 UInt64,\n\tsi32 Int32,\n\tsi64 Int64,\n\tufi32 UInt32,\n\tufi64 UInt64,\n\tsfi32 Int32,\n\tsfi64 Int64,\n\td Float64,\n\tf Float32,\n\tbool Bool,\n\tstr String,\n\tbytes String,\n\tenum1 Enum,\n\tenum2 Enum,\n\tgrp1_i1 Int32,\n\tgrp2_i1 Int32,\n\tmsg1_i1 Int32,\n\tmsg2_i1 Int32,\n\tmsg3_i1 Int32,\n\tmsg3_i2 Int32,\n\tmsg3_i3 Int32,\n\tmsg4_i1 Int32,\n\tmsg4_i2 Int32,\n\tmsg4_i3 Int32,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

// TestWellKnownTypes tests the generator with various well-known message types
// which have custom JSON serialization.
func TestWellKnownTypes(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package"
				message_type <
					name: "FooProto"
					field <
						name: "i32" number: 1 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Int32Value"
					>
					field <
						name: "i64" number: 2 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Int64Value"
					>
					field <
						name: "ui32" number: 3 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.UInt32Value"
					>
					field <
						name: "ui64" number: 4 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.UInt64Value"
					>
					field <
						name: "d" number: 5 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.DoubleValue"
					>
					field <
						name: "f" number: 6 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.FloatValue"
					>
					field <
						name: "bool" number: 7 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.BoolValue"
					>
					field <
						name: "str" number: 8 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.StringValue"
					>
					field <
						name: "bytes" number: 9 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.BytesValue"
					>
					field <
						name: "du" number: 10 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Duration"
					>
					field <
						name: "t" number: 11 type: TYPE_MESSAGE label: LABEL_OPTIONAL
						type_name: ".google.protobuf.Timestamp"
					>
					options < [gen_ch_schema.clickhouse_opts] <table_name: "foo_table"> >
				>
			>
		`,
		map[string]string{
			"example_package/foo_table.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\ti32 Int32,\n\ti64 Int64,\n\tui32 UInt32,\n\tui64 UInt64,\n\td Float64,\n\tf Float32,\n\tbool UInt8,\n\tstr String,\n\tbytes String,\n\tdu Int64,\n\tt DateTime,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

// TestModes tests the generator with different label values.
func TestModes(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i2" number: 2 type: TYPE_INT32 label: LABEL_REQUIRED >
					field < name: "i3" number: 3 type: TYPE_INT32 label: LABEL_REPEATED >
					options < [gen_ch_schema.clickhouse_opts] <table_name: "foo_table"> >
				>
			>
		`,
		map[string]string{
			"example_package/nested/foo_table.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\ti1 Int32,\n\ti2 Int32,\n\ti3 Array(Int32),\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

func TestExtraFields(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package"
				message_type <
					name: "FooProto"
					field <
						name: "i1"
						number: 1
						type: TYPE_INT32
						label: LABEL_OPTIONAL
					>
					options <
						[gen_ch_schema.clickhouse_opts]: <
							table_name: "foo_table"
							extra_fields: [
								"i2:Int32",
								"i3:String:REPEATED",
								"i4:DateTime:REQUIRED",
								"i5:Nested:example_package.nested2.BarProto",
								"i6:Nested:.google.protobuf.DoubleValue:REQUIRED"
							]
						>
					>
				>
			>
			proto_file <
				name: "bar.proto"
				package: "example_package.nested2"
				message_type <
					name: "BarProto"
					field < name: "i1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i2" number: 2 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "i3" number: 3 type: TYPE_INT32 label: LABEL_OPTIONAL >
				>
			>
		`,
		map[string]string{
			"example_package/foo_table.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\ti1 Int32,\n\ti2 Int32,\n\ti3 Array(String),\n\ti4 DateTime,\n\ti5_i1 Int32,\n\ti5_i2 Int32,\n\ti5_i3 Int32,\n\ti6 Float64,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

func TestOrderSchemaByFieldNumber(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package.nested"
				message_type <
					name: "FooProto"
					field < name: "first" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "third" number: 3 type: TYPE_INT32 label: LABEL_REQUIRED >
					field < name: "second" number: 2 type: TYPE_INT32 label: LABEL_REPEATED >
					options < [gen_ch_schema.clickhouse_opts] <
						table_name: "foo_table"
						output_field_order: 1
						>
					>
				>
			>
		`,
		map[string]string{
			"foo_table.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\tfirst Int32,\n\tsecond Array(Int32),\n\tthird Int32,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

func TestNestedOrderSchemaByFieldNumber(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package"
				message_type <
					name: "FooProto"
					field < name: "f_1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "f_3" number: 3 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field <
							name: "bar"
							number: 2
							type: TYPE_MESSAGE
							label: LABEL_OPTIONAL
              type_name: "BarProto"
						>
					options < [gen_ch_schema.clickhouse_opts] <
						table_name: "foo_table"
						output_field_order: 1
						>
					>
				>
				message_type <
					name: "BarProto"
					field < name: "b_2" number: 2 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "b_3" number: 3 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "b_1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_ch_schema.clickhouse_opts] <
					output_field_order: 1
					>
					>
				>
				>
		`,
		map[string]string{
			"foo_table.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\tf_1 Int32,\n\tbar_b_1 Int32,\n\tbar_b_2 Int32,\n\tbar_b_3 Int32,\n\tf_3 Int32,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}

func TestMultipleMessageOrderByFieldNumber(t *testing.T) {
	testConvert(t, `
			file_to_generate: "foo.proto"
			proto_file <
				name: "foo.proto"
				package: "example_package"
				message_type <
					name: "FooProto"
					field < name: "f_1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "f_3" number: 4 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field <
							name: "bar_ordered"
							number: 2
							type: TYPE_MESSAGE
							label: LABEL_OPTIONAL
              type_name: "BarOrderedProto"
						>
						field <
							name: "bar_unordered"
							number: 3
							type: TYPE_MESSAGE
							label: LABEL_OPTIONAL
              type_name: "BarUnOrderedProto"
						>
					options < [gen_ch_schema.clickhouse_opts] <
						table_name: "foo_table"
						output_field_order: 1
						>
					>
				>
				message_type <
					name: "BarOrderedProto"
					field < name: "b_2" number: 2 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "b_3" number: 3 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "b_1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
					options < [gen_ch_schema.clickhouse_opts] <
						output_field_order: 1
					>
					>
				>
				message_type <
					name: "BarUnOrderedProto"
					field < name: "b_2" number: 2 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "b_3" number: 3 type: TYPE_INT32 label: LABEL_OPTIONAL >
					field < name: "b_1" number: 1 type: TYPE_INT32 label: LABEL_OPTIONAL >
				>
				>
		`,
		map[string]string{
			"foo_table.sql": "CREATE TABLE IF NOT EXISTS foo_table\n(\n\tf_1 Int32,\n\tbar_ordered_b_1 Int32,\n\tbar_ordered_b_2 Int32,\n\tbar_ordered_b_3 Int32,\n\tbar_unordered_b_2 Int32,\n\tbar_unordered_b_3 Int32,\n\tbar_unordered_b_1 Int32,\n\tf_3 Int32,\n\tctime DateTime DEFAULT now() COMMENT 'create time'\n) ENGINE = MergeTree()\nORDER BY ctime;",
		})
}
