package constant

const (
	TplCurd      = "assets/templates/tpl/curd.tpl"               // 生成CRUD2模板
	TplStructure = "assets/templates/tpl/structure.tpl"          // 结构体模板
	TplEntity    = "assets/templates/tpl/entity.tpl"             // 结构实体模板
	TplTables    = "assets/templates/tpl/tables.tpl"             // 表结构模板
	TplInit      = "assets/templates/tpl/init.tpl"               // init模板
	TplMarkdown  = "assets/templates/templates/tpl/markdown.tpl" // markdown模板
	TplExample   = "assets/templates/tpl/example.tpl"            // example模板
	TplError     = "assets/templates/tpl/e.tpl"                  // error模板
)

const (
	TplFiberTypes       = "assets/templates/fiber/cmd/internal/types/types.go.tpl"
	TplFiberInHandler   = "assets/templates/fiber/cmd/internal/handler/handler.go.tpl"
	TplFiberInHanRoutes = "assets/templates/fiber/cmd/internal/handler/routes.go.tpl"
	TplFiberInLogic     = "assets/templates/fiber/cmd/internal/logic/logic.go.tpl"
	TplFiberMiddleware  = "assets/templates/fiber/cmd/internal/middleware/middleware.go.tpl"

	TplFiberEConfig       = "assets/templates/fiber/cmd/etc/config.yaml.tpl"
	TplFiberGGlobal       = "assets/templates/fiber/cmd/global/global.go.tpl"
	TplFiberGInit         = "assets/templates/fiber/cmd/global/init.go.tpl"
	TplFiberGMysql        = "assets/templates/fiber/cmd/global/mysql.go.tpl"
	TplFiberGRedis        = "assets/templates/fiber/cmd/global/redis.go.tpl"
	TplFiberInConfig      = "assets/templates/fiber/cmd/internal/config/config.go.tpl"
	TplFiberInConstant    = "assets/templates/fiber/cmd/internal/constant/constant.go.tpl"
	TplFiberMiddleWrapCtx = "assets/templates/fiber/cmd/internal/middleware/wrapCtx_middleware.go.tpl"
	TplFiberMiddleJwt     = "assets/templates/fiber/cmd/internal/middleware/jwt_middleware.go.tpl"
	TplFiberResponse      = "assets/templates/fiber/cmd/internal/response/response.go.tpl"
	TplFiberToolsAes      = "assets/templates/fiber/cmd/internal/tools/cypher/aes.go.tpl"
	TplFiberToolsMd5      = "assets/templates/fiber/cmd/internal/tools/cypher/md5.go.tpl"
	TplFiberToolsJwt      = "assets/templates/fiber/cmd/internal/tools/jwt/jwt.go.tpl"
	TplFiberToolsLogTool  = "assets/templates/fiber/cmd/internal/tools/logger/log_tool.go.tpl"
	TplFiberToolsLogger   = "assets/templates/fiber/cmd/internal/tools/logger/logger.go.tpl"
	TplFiberToolsStr      = "assets/templates/fiber/cmd/internal/tools/str/str.go.tpl"
	TplFiberToolsValidate = "assets/templates/fiber/cmd/internal/tools/validate/validate.go.tpl"
	TplFiberModelBase     = "assets/templates/fiber/cmd/model/base.go.tpl"
	TplFiberModelUser     = "assets/templates/fiber/cmd/model/user.go.tpl"
	TplFiberCore          = "assets/templates/fiber/cmd/core.go.tpl"
	TplFiberGoMod         = "assets/templates/fiber/go.mod.tpl"

	TplFiberPrefix = "assets/templates/fiber"
)

var TplFixedList = []string{
	TplFiberEConfig,
	TplFiberGGlobal,
	TplFiberGInit,
	TplFiberGMysql,
	TplFiberGRedis,
	TplFiberInConfig,
	TplFiberInConstant,
	TplFiberMiddleWrapCtx,
	TplFiberMiddleJwt,
	TplFiberResponse,
	TplFiberToolsAes,
	TplFiberToolsMd5,
	TplFiberToolsJwt,
	TplFiberToolsLogTool,
	TplFiberToolsLogger,
	TplFiberToolsStr,
	TplFiberToolsValidate,
	TplFiberModelBase,
	TplFiberModelUser,
	TplFiberCore,
	TplFiberGoMod,
}
