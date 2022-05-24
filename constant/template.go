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
	TplFiberTypes       = "assets/templates/fiber/cmd/internal/types/types.tpl"
	TplFiberInHandler   = "assets/templates/fiber/cmd/internal/handler/handler.tpl"
	TplFiberInHanRoutes = "assets/templates/fiber/cmd/internal/handler/routes.tpl"
	TplFiberInLogic     = "assets/templates/fiber/cmd/internal/logic/logic.tpl"
	TplFiberMiddleware  = "assets/templates/fiber/cmd/internal/middleware/middleware.tpl"

	TplFiberEConfig       = "assets/templates/fiber/cmd/etc/config.yaml"
	TplFiberGGlobal       = "assets/templates/fiber/cmd/global/global.tpl"
	TplFiberGInit         = "assets/templates/fiber/cmd/global/init.tpl"
	TplFiberGMysql        = "assets/templates/fiber/cmd/global/mysql.tpl"
	TplFiberGRedis        = "assets/templates/fiber/cmd/global/redis.tpl"
	TplFiberInConfig      = "assets/templates/fiber/cmd/internal/config/config.tpl"
	TplFiberInConstant    = "assets/templates/fiber/cmd/internal/constant/constant.tpl"
	TplFiberMiddleWrapCtx = "assets/templates/fiber/cmd/internal/middleware/wrapCtx_middleware.tpl"
	TplFiberMiddleJwt     = "assets/templates/fiber/cmd/internal/middleware/jwt_middleware.tpl"
	TplFiberResponse      = "assets/templates/fiber/cmd/internal/response/response.tpl"
	TplFiberToolsAes      = "assets/templates/fiber/cmd/internal/tools/cypher/aes.tpl"
	TplFiberToolsMd5      = "assets/templates/fiber/cmd/internal/tools/cypher/md5.tpl"
	TplFiberToolsJwt      = "assets/templates/fiber/cmd/internal/tools/jwt/jwt.tpl"
	TplFiberToolsLogTool  = "assets/templates/fiber/cmd/internal/tools/logger/log_tool.tpl"
	TplFiberToolsLogger   = "assets/templates/fiber/cmd/internal/tools/logger/logger.tpl"
	TplFiberToolsStr      = "assets/templates/fiber/cmd/internal/tools/str/str.tpl"
	TplFiberToolsValidate = "assets/templates/fiber/cmd/internal/tools/validate/validate.tpl"
	TplFiberModelBase     = "assets/templates/fiber/cmd/model/base.tpl"
	TplFiberModelUser     = "assets/templates/fiber/cmd/model/user.tpl"
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
}
