protoc --go_out=src ch_field.proto ch_table.proto
protoc --proto_path=. \
    --go_out=.\
  --go_opt=Mch_field.proto=github.com/jdxhu/protoc-gen-ch-schema/protos \
  --go_opt=Mch_table.proto=github.com/jdxhu/protoc-gen-ch-schema/protos \
  ch_field.proto ch_table.proto


protoc-gen-ch-schema --go_out=src ch_field.proto ch_table.proto
protoc --ch-schema_out=. foo.proto --proto_path=. --proto_path=<path_to_google_proto_folder>/src
