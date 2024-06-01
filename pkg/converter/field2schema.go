package converter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jdxhu/protoc-gen-ch-schema/protos"
)

const (
	sqlTplMergeTree          = "CREATE TABLE IF NOT EXISTS {table_name}\n(\n\t{fields}\n) ENGINE = MergeTree()[\nORDER BY {order_by_expr}][\nPARTITION BY {partion_expr}][\nTTL {ttl_expr}][\nSETTINGS {settings}];"
	sqlTplReplacingMergeTree = "CREATE TABLE IF NOT EXISTS {table_name}\n(\n\t{fields}\n) ENGINE = ReplacingMergeTree([{replaceing_ver}[, {replacing_deleted_row}]])[\nORDER BY {order_by_expr}][\nPARTITION BY {partion_expr}][\nTTL {ttl_expr}][\nSETTINGS {settings}];"
	sqlTplSummingMergeTree   = "CREATE TABLE IF NOT EXISTS {table_name}\n(\n\t{fields}\n) ENGINE = SummingMergeTree([({summing_columns})])[\nORDER BY {order_by_expr}][\nPARTITION BY {partion_expr}][\nTTL {ttl_expr}][\nSETTINGS {settings}];"

	// defaultCreateTimeField is the default column name for create time
	defaultCreateTimeField = "ctime"

	// defaultReplaceingVer is the default column name for replacing version
	defaultReplaceingVer = "ctime"
	defaultReplaceingDel = "deleted"
)

func CheckFill(tplString string, kwargs map[string]string) string {
	filled := false
	ret := ""
	for i := 0; i < len(tplString); i++ {
		if tplString[i] == '{' {
			for j := i + 1; j < len(tplString); j++ {
				if tplString[j] == '}' {
					key := tplString[i+1 : j]
					if v, ok := kwargs[key]; ok {
						ret += v
						filled = true
					}
					i = j
					break
				}
			}
		} else if tplString[i] == ']' {
			if filled {
				ret += CheckFill(tplString[i+1:], kwargs)
			} else {
				ret = CheckFill(tplString[i+1:], kwargs)
			}
			break
		} else if tplString[i] == '[' {
			ret += CheckFill(tplString[i+1:], kwargs)
			break
		} else {
			ret += string(tplString[i])
		}
	}
	return ret
}

func clickhouseSchema(opts *protos.ClickHouseMessageOptions, schema []*Field) (sql string, err error) {
	schemaFlatten := fieldFlatten(schema, "")

	// add default create time field
	schemaFlatten = append(schemaFlatten, &Field{
		Name:                   defaultCreateTimeField,
		Type:                   "DateTime",
		Description:            "create time",
		DefaultValueExpression: "now()",
	})
	switch opts.TableEngine {
	case protos.ClickHouseMessageOptions_MERGE_TREE:
		return sqlGenTableCreateMergeTree(opts, schemaFlatten)
	case protos.ClickHouseMessageOptions_REPLACING_MERGE_TREE:
		return sqlGenTableCreateReplacingMergeTree(opts, schemaFlatten)
	case protos.ClickHouseMessageOptions_SUMMING_MERGE_TREE:
		return sqlGenTableCreateSummingMergeTree(opts, schemaFlatten)
	case protos.ClickHouseMessageOptions_TABLE_ENGINE_UNSPECIFIED:
		return sqlGenTableCreateMergeTree(opts, schemaFlatten)
	default:
		return "", fmt.Errorf("table engine not supported: %s", opts.TableEngine)
	}
}

func fieldFlatten(schema []*Field, prefix string) (fields []*Field) {
	for _, field := range schema {
		if field.Type != chTypeNested {
			field.Name = prefix + field.Name
			fields = append(fields, field)
		} else {
			fields = append(fields, fieldFlatten(field.Fields, prefix+field.Name+"_")...)
		}
	}
	return
}

func sqlGenField(field *Field) string {
	ret := "{name} {type}[ {nullable}][ DEFAULT {default}][ COMMENT '{comment}']"
	fm := map[string]string{}
	fm["name"] = field.Name
	if field.Mode == ModeRepeated {
		fm["type"] = "Array(" + field.Type + ")"
	} else {
		fm["type"] = field.Type
	}
	// 为了保证数据插入不会报错，暂时不支持非空约束
	// if field.Mode == ModeRequired {
	// 	fm["nullable"] = "NOT NULL"
	// }
	if field.DefaultValueExpression != "" {
		fm["default"] = field.DefaultValueExpression
	}
	if field.Description != "" {
		fm["comment"] = field.Description
	}
	return CheckFill(ret, fm)
}

func sqlGenTTL(ttl string) (string, error) {
	if ttl == "" {
		return "", fmt.Errorf("ttl is empty")
	}
	lt := len(ttl)
	unit := ttl[lt-1:]
	cnt, err := strconv.Atoi(ttl[:lt-1])
	if err != nil {
		return "", err
	}

	var unitFmt string
	switch unit {
	case "y":
		unitFmt = "YEAR"
	case "q":
		unitFmt = "QUARTER"
	case "m":
		unitFmt = "MONTH"
	case "w":
		unitFmt = "WEEK"
	case "d":
		unitFmt = "DAY"
	case "h":
		unitFmt = "HOUR"
	case "M":
		unitFmt = "MINUTE"
	case "s":
		unitFmt = "SECOND"
	default:
		return "", fmt.Errorf("invalid ttl unit: '%s'", unit)
	}
	return fmt.Sprintf("%s + INTERVAL %d %s", defaultCreateTimeField, cnt, unitFmt), nil
}

func sqlGenSettings(opts *protos.ClickHouseMessageOptions) string {
	settings := []string{}
	for _, s := range opts.Settings {
		settings = append(settings, fmt.Sprintf("%s = '%s'", s.Name, s.Value))
	}
	return strings.Join(settings, ",")
}

func sqlGenTableCommonFill(sql string, kwargs map[string]string, opts *protos.ClickHouseMessageOptions, schemas []*Field) (string, error) {
	kwargs["table_name"] = opts.TableName
	fields := []string{}
	for _, schema := range schemas {
		fields = append(fields, sqlGenField(schema))
	}
	fieldsStr := strings.Join(fields, ",\n\t")
	kwargs["fields"] = fieldsStr
	if opts.OrderBy != "" {
		kwargs["order_by_expr"] = opts.OrderBy
	} else {
		kwargs["order_by_expr"] = defaultCreateTimeField
	}
	switch opts.TimePartition {
	case protos.ClickHouseMessageOptions_MONTH:
		kwargs["partion_expr"] = fmt.Sprintf("toYYYYMM(%s)", defaultCreateTimeField)
	case protos.ClickHouseMessageOptions_DAY:
		kwargs["partion_expr"] = fmt.Sprintf("toYYYYMMDD(%s)", defaultCreateTimeField)
	case protos.ClickHouseMessageOptions_HOUR:
		kwargs["partion_expr"] = fmt.Sprintf("toYYYYMMDDHH(%s)", defaultCreateTimeField)
	}
	if opts.Ttl != "" {
		ttl_expr, err := sqlGenTTL(opts.Ttl)
		if err != nil {
			return "", err
		}
		kwargs["ttl_expr"] = ttl_expr
	}
	if opts.Settings != nil {
		kwargs["settings"] = sqlGenSettings(opts)
	}
	return CheckFill(sql, kwargs), nil
}

func sqlGenTableCreateMergeTree(opts *protos.ClickHouseMessageOptions, schemas []*Field) (string, error) {
	return sqlGenTableCommonFill(sqlTplMergeTree, map[string]string{}, opts, schemas)
}

func sqlGenTableCreateReplacingMergeTree(opts *protos.ClickHouseMessageOptions, schemas []*Field) (string, error) {
	fm := map[string]string{}
	if opts.ReplacingVersionCol == "" {
		fm["replaceing_ver"] = defaultReplaceingVer
	} else {
		fm["replaceing_ver"] = opts.ReplacingVersionCol
	}
	if opts.ReplacingDeletedCol == "" {
		fm["replacing_deleted_row"] = defaultReplaceingDel
		schemas = append(schemas, &Field{
			Name:        defaultReplaceingDel,
			Type:        "UInt8",
			Description: "state: 0, deleted: 1",
		})
	} else {
		fm["replacing_deleted_row"] = opts.ReplacingDeletedCol
	}

	return sqlGenTableCommonFill(sqlTplReplacingMergeTree, fm, opts, schemas)
}

func sqlGenTableCreateSummingMergeTree(opts *protos.ClickHouseMessageOptions, schemas []*Field) (string, error) {
	fm := map[string]string{}
	if opts.SummingAggregateCols != nil {
		fm["summing_columns"] = strings.Join(opts.SummingAggregateCols, ", ")

	}
	return sqlGenTableCommonFill(sqlTplSummingMergeTree, fm, opts, schemas)
}
