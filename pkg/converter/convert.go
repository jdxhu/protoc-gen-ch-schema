package converter

import (
	"fmt"
	"io"
	"path"
	"sort"
	"strings"

	"github.com/golang/glog"
	"github.com/jdxhu/protoc-gen-ch-schema/protos"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	descriptor "google.golang.org/protobuf/types/descriptorpb"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

const (
	chTypeInt32    = "Int32"
	chTypeInt64    = "Int64"
	chTypeUInt32   = "UInt32"
	chTypeUInt64   = "UInt64"
	chTypeFloat32  = "Float32"
	chTypeFloat64  = "Float64"
	chTypeUInt8    = "UInt8"
	chTypeString   = "String"
	chTypeDateTime = "DateTime"
	chTypeEnum     = "Enum"
	chTypeNested   = "Nested"

	ModeOptional = "NULLABLE"
	ModeRequired = "REQUIRED"
	ModeRepeated = "REPEATED"
)

var (
	typeFromWKT = map[string]string{
		".google.protobuf.Int32Value":  chTypeInt32,
		".google.protobuf.Int64Value":  chTypeInt64,
		".google.protobuf.UInt32Value": chTypeUInt32,
		".google.protobuf.UInt64Value": chTypeUInt64,
		".google.protobuf.DoubleValue": chTypeFloat64,
		".google.protobuf.FloatValue":  chTypeFloat32,
		".google.protobuf.BoolValue":   chTypeUInt8,
		".google.protobuf.StringValue": chTypeString,
		".google.protobuf.BytesValue":  chTypeString,
		".google.protobuf.Duration":    chTypeInt64,
		".google.protobuf.Timestamp":   chTypeDateTime,
	}
	typeFromFieldType = map[descriptor.FieldDescriptorProto_Type]string{
		descriptor.FieldDescriptorProto_TYPE_DOUBLE: chTypeFloat64,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:  chTypeFloat32,

		descriptor.FieldDescriptorProto_TYPE_INT64:    chTypeInt64,
		descriptor.FieldDescriptorProto_TYPE_UINT64:   chTypeUInt64,
		descriptor.FieldDescriptorProto_TYPE_INT32:    chTypeInt32,
		descriptor.FieldDescriptorProto_TYPE_UINT32:   chTypeUInt32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:  chTypeUInt64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32:  chTypeUInt32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32: chTypeInt32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: chTypeInt64,
		descriptor.FieldDescriptorProto_TYPE_SINT32:   chTypeInt32,
		descriptor.FieldDescriptorProto_TYPE_SINT64:   chTypeInt64,

		descriptor.FieldDescriptorProto_TYPE_STRING: chTypeString,
		descriptor.FieldDescriptorProto_TYPE_BYTES:  chTypeString,
		descriptor.FieldDescriptorProto_TYPE_ENUM:   chTypeEnum,

		descriptor.FieldDescriptorProto_TYPE_BOOL: chTypeUInt8,

		descriptor.FieldDescriptorProto_TYPE_GROUP:   chTypeNested,
		descriptor.FieldDescriptorProto_TYPE_MESSAGE: chTypeNested,
	}

	modeFromFieldLabel = map[descriptor.FieldDescriptorProto_Label]string{
		descriptor.FieldDescriptorProto_LABEL_OPTIONAL: ModeOptional,
		descriptor.FieldDescriptorProto_LABEL_REQUIRED: ModeRequired,
		descriptor.FieldDescriptorProto_LABEL_REPEATED: ModeRepeated,
	}
)

// Field describes the schema of a field in ClickHouse.
type Field struct {
	Name                   string   `json:"name"`
	Type                   string   `json:"type"`
	Mode                   string   `json:"mode"`
	Description            string   `json:"description,omitempty"`
	Fields                 []*Field `json:"fields,omitempty"`
	DefaultValueExpression string   `json:"defaultValueExpression,omitempty"`
}

func registerType(pkgName *string, msg *descriptor.DescriptorProto, comments Comments, path string) {
	pkg := globalPkg
	if pkgName != nil {
		for _, node := range strings.Split(*pkgName, ".") {
			if pkg == globalPkg && node == "" {
				// Skips leading "."
				continue
			}

			child, ok := pkg.children[node]
			if !ok {
				child = &ProtoPackage{
					name:     pkg.name + "." + node,
					parent:   pkg,
					children: make(map[string]*ProtoPackage),
					types:    make(map[string]*descriptor.DescriptorProto),
					comments: make(map[string]Comments),
					path:     make(map[string]string),
				}
				pkg.children[node] = child
			}
			pkg = child
		}
	}

	pkg.types[msg.GetName()] = msg
	pkg.comments[msg.GetName()] = comments
	pkg.path[msg.GetName()] = path
}

func convertField(
	curPkg *ProtoPackage,
	desc *descriptor.FieldDescriptorProto,
	msgOpts *protos.ClickHouseMessageOptions,
	parentMessages map[*descriptor.DescriptorProto]bool,
	comments Comments,
	path string) (*Field, error) {

	field := &Field{
		Name: desc.GetName(),
	}
	if msgOpts.GetUseJsonNames() && desc.GetJsonName() != "" {
		field.Name = desc.GetJsonName()
	}

	var ok bool
	field.Mode, ok = modeFromFieldLabel[desc.GetLabel()]
	if !ok {
		return nil, fmt.Errorf("unrecognized field label: %s", desc.GetLabel().String())
	}

	field.Type, ok = typeFromFieldType[desc.GetType()]
	if !ok {
		return nil, fmt.Errorf("unrecognized field type: %s", desc.GetType().String())
	}

	if comment := comments.Get(path); comment != "" {
		field.Description = comment
	}

	opts := desc.GetOptions()
	if opts != nil && proto.HasExtension(opts, protos.E_Clickhouse) {
		opt := proto.GetExtension(opts, protos.E_Clickhouse).(*protos.ClickHouseFieldOptions)
		if opt.Ignore {
			// skip the field below
			return nil, nil
		}

		if opt.Require {
			field.Mode = ModeRequired
		}

		if len(opt.TypeOverride) > 0 {
			field.Type = opt.TypeOverride
		}

		if len(opt.Name) > 0 {
			field.Name = opt.Name
		}

		if len(opt.Description) > 0 {
			field.Description = opt.Description
		}

		if len(opt.DefaultValueExpression) > 0 {
			field.DefaultValueExpression = opt.DefaultValueExpression
		}
	}

	if len(field.Description) > 1024 {
		field.Description = field.Description[:1021] + "..."
	}

	if field.Type != chTypeNested {
		return field, nil
	}
	if t, ok := typeFromWKT[desc.GetTypeName()]; ok {
		field.Type = t
		return field, nil
	}

	fields, err := convertFieldsForType(curPkg, desc.GetTypeName(), parentMessages)
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 { // discard RECORDs that would have zero fields
		return nil, nil
	}

	field.Fields = fields

	return field, nil
}

func convertExtraField(curPkg *ProtoPackage, extraFieldDefinition string, parentMessages map[*descriptor.DescriptorProto]bool) (*Field, error) {
	parts := strings.Split(extraFieldDefinition, ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("expecting at least 2 parts in extra field definition separated by colon, got %d", len(parts))
	}

	field := &Field{
		Name: parts[0],
		Type: parts[1],
		Mode: ModeOptional,
	}

	modeIndex := 2
	if field.Type == chTypeNested {
		modeIndex = 3
	}
	if len(parts) > modeIndex {
		field.Mode = parts[modeIndex]
	}

	if field.Type != chTypeNested {
		return field, nil
	}

	if len(parts) < 3 {
		return nil, fmt.Errorf("extra field %s has no type defined", field.Type)
	}

	typeName := parts[2]

	if t, ok := typeFromWKT[typeName]; ok {
		field.Type = t
		return field, nil
	}

	fields, err := convertFieldsForType(curPkg, typeName, parentMessages)
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 { // discard RECORDs that would have zero fields
		return nil, nil
	}

	field.Fields = fields

	return field, nil
}

func convertFieldsForType(curPkg *ProtoPackage,
	typeName string,
	parentMessages map[*descriptor.DescriptorProto]bool) ([]*Field, error) {
	recordType, ok, comments, path := curPkg.lookupType(typeName)
	if !ok {
		return nil, fmt.Errorf("no such message type named %s", typeName)
	}

	fieldMsgOpts, err := getClickhouseMessageOptions(recordType)
	if err != nil {
		return nil, err
	}

	return convertMessageType(curPkg, recordType, fieldMsgOpts, parentMessages, comments, path)
}

func convertMessageType(
	curPkg *ProtoPackage,
	msg *descriptor.DescriptorProto,
	opts *protos.ClickHouseMessageOptions,
	parentMessages map[*descriptor.DescriptorProto]bool,
	comments Comments,
	path string) (schema []*Field, err error) {

	if parentMessages[msg] {
		glog.Infof("Detected recursion for message %s, ignoring subfields", *msg.Name)
		return
	}

	if glog.V(4) {
		glog.Info("Converting message: ", prototext.Format(msg))
	}

	parentMessages[msg] = true
	fields := msg.GetField()
	// Sort fields by the field numbers if the option is set.
	if opts.GetOutputFieldOrder() == protos.ClickHouseMessageOptions_FIELD_ORDER_BY_NUMBER {
		sort.Slice(fields, func(i, j int) bool {
			return fields[i].GetNumber() < fields[j].GetNumber()
		})
	}
	for fieldIndex, fieldDesc := range fields {
		fieldCommentPath := fmt.Sprintf("%s.%d.%d", path, fieldPath, fieldIndex)
		field, err := convertField(curPkg, fieldDesc, opts, parentMessages, comments, fieldCommentPath)
		if err != nil {
			glog.Errorf("Failed to convert field %s in %s: %v", fieldDesc.GetName(), msg.GetName(), err)
			return nil, err
		}

		// if we got no error and the field is nil, skip it
		if field != nil {
			schema = append(schema, field)
		}
	}

	for _, extraField := range opts.GetExtraFields() {
		field, err := convertExtraField(curPkg, extraField, parentMessages)
		if err != nil {
			glog.Errorf("Failed to convert extra field %s in %s: %v", extraField, msg.GetName(), err)
			return nil, err
		}

		schema = append(schema, field)
	}

	parentMessages[msg] = false

	return
}

func convertFile(file *descriptor.FileDescriptorProto, reqParams string) ([]*plugin.CodeGeneratorResponse_File, error) {
	name := path.Base(file.GetName())
	pkg, ok := globalPkg.relativelyLookupPackage(file.GetPackage())
	if !ok {
		return nil, fmt.Errorf("no such package found: %s", file.GetPackage())
	}

	comments := ParseComments(file)
	response := []*plugin.CodeGeneratorResponse_File{}
	for msgIndex, msg := range file.GetMessageType() {
		path := fmt.Sprintf("%d.%d", messagePath, msgIndex)

		opts, err := getClickhouseMessageOptions(msg)
		if err != nil {
			return nil, err
		}
		if opts == nil {
			continue
		}

		tableName := opts.GetTableName()
		if len(tableName) == 0 {
			continue
		}

		glog.V(2).Info("Generating schema for a message type ", msg.GetName())
		schema, err := convertMessageType(pkg, msg, opts, make(map[*descriptor.DescriptorProto]bool), comments, path)
		if err != nil {
			glog.Errorf("Failed to convert %s: %v", name, err)
			return nil, err
		}

		// bytes, err := json.MarshalIndent(schema, "", " ")
		ddl, err := clickhouseSchema(opts, schema)
		if err != nil {
			glog.Error("Failed to encode schema", err)
			return nil, err
		}
		// ddl := string(bytes)
		var pth string
		if strings.Contains(reqParams, "source-path-relative") {
			pth = fmt.Sprintf("%s/%s.ch.sql", strings.Replace(file.GetPackage(), ".", "/", -1), tableName)
		} else {
			pth = fmt.Sprintf("%s.ch.sql", tableName)
		}

		resFile := &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(pth),
			Content: proto.String(ddl),
		}
		response = append(response, resFile)
	}

	return response, nil
}

// getClickhouseMessageOptions returns the clickhouse options for the given message.
// If an error is encountered, it is returned instead. If no error occurs, but
// the message has no gen_ch_schema.clickhouse_opts option, this function returns
// nil, nil.
func getClickhouseMessageOptions(msg *descriptor.DescriptorProto) (*protos.ClickHouseMessageOptions, error) {
	options := msg.GetOptions()
	if options == nil {
		return nil, nil
	}

	if !proto.HasExtension(options, protos.E_ClickhouseOpts) {
		return nil, nil
	}

	return proto.GetExtension(options, protos.E_ClickhouseOpts).(*protos.ClickHouseMessageOptions), nil
}

func Convert(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	generateTargets := make(map[string]bool)
	for _, file := range req.GetFileToGenerate() {
		generateTargets[file] = true
	}

	res := &plugin.CodeGeneratorResponse{
		SupportedFeatures: proto.Uint64(uint64(plugin.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)),
	}
	for _, file := range req.GetProtoFile() {
		for msgIndex, msg := range file.GetMessageType() {
			glog.V(1).Infof("Loading a message type %s from package %s", msg.GetName(), file.GetPackage())
			registerType(file.Package, msg, ParseComments(file), fmt.Sprintf("%d.%d", messagePath, msgIndex))
		}
	}
	for _, file := range req.GetProtoFile() {
		if _, ok := generateTargets[file.GetName()]; ok {
			glog.V(1).Info("Converting ", file.GetName())
			converted, err := convertFile(file, req.GetParameter())
			if err != nil {
				res.Error = proto.String(fmt.Sprintf("Failed to convert %s: %v", file.GetName(), err))
				return res, err
			}
			res.File = append(res.File, converted...)
		}
	}
	return res, nil
}

// ConvertFrom converts input from protoc to a CodeGeneratorRequest and starts conversion
// Returning a CodeGeneratorResponse containing either an error or the results of converting the given proto
func ConvertFrom(rd io.Reader) (*plugin.CodeGeneratorResponse, error) {
	glog.V(1).Info("Reading code generation request")
	input, err := io.ReadAll(rd)
	if err != nil {
		glog.Error("Failed to read request:", err)
		return nil, err
	}
	req := &plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(input, req)
	if err != nil {
		glog.Error("Can't unmarshal input:", err)
		return nil, err
	}

	glog.V(1).Info("Converting input")
	return Convert(req)
}
