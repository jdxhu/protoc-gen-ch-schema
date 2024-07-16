# protoc-gen-ch-schema

protoc-gen-ch-schema 是一个protibuf扩展，用于将 .proto 的定义文件转换为clickhouse sql建表语句。
该项目参考谷歌开源的 protoc-gen-bq-schema 项目开发完成，不过最后结果是生成clickhouse直接执行的建表语句

需要注意，生成的的ddl语句会新增ctime字段，该字段将用于order by、时间分区键、ttl。

## Installation

```sh
go install github.com/jdxhu/protoc-gen-ch-schema@latest
```

## Usage

```sh
protoc --ch-schema_out=path/to/outdir foo.proto
```

`protoc` 和 `protoc-gen-ch-schema` 必须在系统环境变量`${PATH}`中

生成的文件会以`.sql`为文件结尾，以package为目录存放在outdir中

```sh
protoc --ch-schema_out=path/to/out/dir foo.proto --proto_path=. --proto_path=<path_to_google_proto_folder>/src
```

### Example

对于详细的配置，参考ch_table.proto/ch_field.proto定义文件。
示例：

```protobuf
syntax = "proto2";
package foo;
import "ch_table.proto";
import "ch_field.proto";

message Bar {
  option (gen_ch_schema.clickhouse_opts) = {
    table_name: "bar_db.bar_table"
    table_engine: MERGE_TREE // 目前支持三种表引擎: MERGE_TREE, REPLACING_MERGE_TREE, SUMMING_MERGE_TREE。默认为MERGE_TREE
    time_partition: MONTH  // 分区设置，按clickhouse官方建议，month分区可以满足巨大多数需求。默认不使用分区
    ttl: "3m"  // 数据过期时间，超限的数据会被删除，时间定义：y/q/m/w/d/h/M/s -> 年/季度/月/周/天/小时/分钟/秒。 默认没有数据超时。
    settings: [
        {
            name: "index_granularity"
            value: "8192"
        }
    ]  // settings字段。标识建表时额外的配置项。参考clickhouse官方文档。默认settings为空
    // replacing_version_col: "e"  // 指定replacing表引擎的version字段，参考官方文档定义
    // replacing_deleted_col  // 指定replacing表引擎的is_deleted状态字段，参考官方文档定义
  };
  message Nested {
    repeated int32 a = 1;
    repeated string b = 2;
  }

  message Nested {
    repeated int32 a = 1;
  }

  // Description of field a -- this is an int32
  required int32 a = 1;

  // Nested b structure
  optional Nested b = 2;

  // Repeated c string
  repeated string c = 3;

  optional bool d = 4 [(gen_ch_schema.clickhouse).ignore = true];

  // TIMESTAMP (uint64 in proto) - required in ClickHouse
  optional uint64 e = 5 [
    (gen_ch_schema.clickhouse) = {
      require: true
      type_override: 'DateTime'
    }
  ];
}

message Baz {
  required int32 a = 1;
}
```

`protoc --ch-schema_out=. foo.proto` 会生成 `foo/bar_table.sql`文件.

如果需要保留package相对路径，需要加入 `--ch-schema_opt=source_relative` 配置.

`foo.Baz` 不会在建表字段里，根据 `gen_ch_schema.clickhouse_opts`配置.
