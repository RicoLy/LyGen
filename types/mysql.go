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

// EntityReq 生成实体的请求结构
type EntityReq struct {
	Index        int    // 序列号
	TableName    string // 表名称
	TableComment string // 表注释
	Path         string // 文件路径
	EntityPath   string // 实体路径
	Pkg          string // 命名空间名称
	EntityPkg    string // entity实体的空间名称
	FormatList   []string
	TableDesc    []*TableDesc // 表详情
}

// FieldsInfo 表字段信息
type FieldsInfo struct {
	Name         string
	Type         string
	NullType     string
	DbName       string
	DbOriField   string // 数据库的原生字段名称
	FormatFields string
	Remark       string
}

// TableInfo 表信息结构
type TableInfo struct {
	Table            string        // 表名驼峰
	TableName        string        // 表名
	NullTable        string        // 空表名称
	TableComment     string        // 表注释
	TableCommentNull string        // 表注释
	Fields           []*FieldsInfo // 表字段
	PackageName      string        //包名
}

// SqlFieldInfo 查询使用的字段结构信息
type SqlFieldInfo struct {
	HumpName string // 驼峰字段名称
	Comment  string // 字段注释
}

// NullSqlFieldInfo 查询使用的空字段结构信息
type NullSqlFieldInfo struct {
	GoType       string // golang类型
	HumpName     string // 驼峰字段名称
	OriFieldType string // 原数据库类型
	Comment      string // 字段注释
}

// SqlInfo 生成select,update,insert,delete所需信息
type SqlInfo struct {
	TableName           string              // 表名
	PrimaryKey          string              // 主键字段
	PrimaryType         string              // 主键类型
	StructTableName     string              // 结构表名称
	NullStructTableName string              // 判断为空的表名
	PkgEntity           string              // 实体空间名称
	PkgTable            string              // 表的空间名称
	UpperTableName      string              // 大写的表名
	AllFieldList        string              // 所有字段列表,如: id,name
	InsertFieldList     string              // 插入字段列表,如:id,name
	InsertMark          string              // 插入字段使用多少个?,如: ?,?,?
	UpdateFieldList     string              // 更新字段列表
	SecondField         string              // 存放第二个字段
	UpdateListField     []string            // 更新字段列表
	FieldsInfo          []*SqlFieldInfo     // 字段信息
	NullFieldsInfo      []*NullSqlFieldInfo // 判断为空时
	InsertInfo          []*SqlFieldInfo     // 插入信息
	UpdateInfo          []*SqlFieldInfo     // 跟新信息
	PackageName         string              // 包名
}