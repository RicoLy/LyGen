package types

// TableNameAndComment 表名与表注释
type TableNameAndComment struct {
	Index   int
	Name    string
	Comment string
}

// TableList 表名和表注释
type TableList struct {
	Index          int
	UpperTableName string
	TableName      string
	Comment        string
}

// TableDesc 表结构详情
type TableDesc struct {
	Index            int
	ColumnName       string // 数据库原始字段
	GoColumnName     string // go使用的字段名称
	OriMysqlType     string // 数据库原始类型
	UpperMysqlType   string // 转换大写的类型
	GolangType       string // 转换成golang类型
	MysqlNullType    string // MYSQL对应的空类型
	PrimaryKey       bool   // 是否是主键
	IsNull           string // 是否为空
	DefaultValue     string // 默认值
	ColumnTypeNumber string // 类型(长度)
	ColumnComment    string // 备注
}

// MarkDownDataChild markdown子数据
type MarkDownDataChild struct {
	Index     int    // 自增
	TableName string // 表名
	Comment   string // 表备注
	List      []*TableDesc
}

// MarkDownData markdown
type MarkDownData struct {
	TableList []*TableList
	DescList  []*MarkDownDataChild
}

