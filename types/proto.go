package types

// proto信息
// 结构体信息
type Message struct {
	Comment      string         // 注释
	Meta         string         // 元数据
	Name         string         // 消息名
	ElementInfos []*ElementInfo // proto字段
}

// 字段信息
type ElementInfo struct {
	Meta string // 元数据
	Name string // 名称
	Type string // golang 数据类型
	Tags string // 标签信息 tag | value  json:"pid" form:"pid"
}

// 方法信息
type Method struct {
	Comment     string   // 注释
	Mata        string   // 源信息
	MethodType  string   // 方法类型
	Path        string   // 路径
	PathParams  []string // 路径参数
	MiddleWares []string // 中间件
	Name        string   // 方法名
	Request     string   // 参数
	Response    string   // 返回值
	Group       string   // 分组名
}

// 服务信息
type Service struct {
	Comment string    // 注释
	Meta    string    // 元数据
	Name    string    // 服务名
	Group   string    // 分组信息
	Prefix  string    // 路由前缀
	Methods []*Method // 方法
}
