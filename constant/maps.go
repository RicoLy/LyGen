package constant

import "LyGen/types"

// MysqlTypeToGoType mysql类型 <=> golang类型
var MysqlTypeToGoType = map[string]string{
	"tinyint":    "int32",
	"smallint":   "int32",
	"mediumint":  "int32",
	"int":        "int32",
	"integer":    "int64",
	"bigint":     "int64",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}

// MysqlTypeToGoNullType MYSQL => golang mysql NULL TYPE
var MysqlTypeToGoNullType = map[string]string{
	"tinyint":    "sql.NullInt32",
	"smallint":   "sql.NullInt32",
	"mediumint":  "sql.NullInt32",
	"int":        "sql.NullInt32",
	"integer":    "sql.NullInt64",
	"bigint":     "sql.NullInt64",
	"float":      "sql.NullFloat64",
	"double":     "sql.NullFloat64",
	"decimal":    "sql.NullFloat64",
	"date":       "sql.NullString",
	"time":       "sql.NullString",
	"year":       "sql.NullString",
	"datetime":   "mysql.NullTime",
	"timestamp":  "mysql.NullTime",
	"char":       "sql.NullString",
	"varchar":    "sql.NullString",
	"tinyblob":   "sql.NullString",
	"tinytext":   "sql.NullString",
	"blob":       "sql.NullString",
	"text":       "sql.NullString",
	"mediumblob": "sql.NullString",
	"mediumtext": "sql.NullString",
	"longblob":   "sql.NullString",
	"longtext":   "sql.NullString",
}

var CmdHelp = []types.CmdEntity{
	{"0", "Set build directory"},
	{"1", "Generate the table markdown document"},
	{"2", "Generate table structure entities"},
	{"21", "myGenerate table structure entities"},
	{"3", "Generate CURD insert, delete, update and select"},
	{"31", "myGenerate CURD insert, delete, update and select"},
	{"32", "myGenerate Sql Entity And CURD insert, delete, update and select"},
	{"4", "Sets the struct mapping name"},
	{"5", "Find or set the table name"},
	{"7, c, clear", "Clear the screen"},
	{"8, h, help", "Show help list"},
	{"9, q, quit", "Quit"},
}
