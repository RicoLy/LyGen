package constant


const (
	RegExpMessage     = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s|]*)message\\s*(\\w*)\\s*{([\\s\\|\\*\\/@\\w:\",=;]*)}"
	RegExpElementInfo = "\\/\\/\\s*@\\w*\\:\\s*([\\s\\w@:\",|=*]*)\\n([\\s\\w]*)=\\s*\\d;"
	RegExpService     = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s@:\"\\/]*)service\\s*(\\w*)\\s*{([\\s/\u4e00-\u9fa5\\w@:(){}|]*)}\\s*}"
	RegExpAnoInfo     = "@(\\w*):\\s*([/:|\\w]*)"
	RegExpMethod      = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s@:|/]*)rpc\\s*(\\w*)\\s*\\((\\w*)\\)\\s*returns\\s*\\((\\w*)\\)\\s*{"
)