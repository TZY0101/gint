package cmd_gen

type ApiTmpl struct {
	PackageName string
	Imports     []string
	ApiName     string
	Req         string
	Resp        string
	HandlerName string
	LogicName   string
}

type RoutesTmplData struct {
	Imports            []string
	ModuleToUpperCamel string
	ModuleToLowerCamel string
	Prefix             string
	Apis               []Route
}

type Route struct {
	MethodToUpper string
	Path           string
	ApiNameToTitle string
}

// 路由模板参数
type RouterTmplData struct {
	Imports     []string
	Module      string
	FuncName    string
	RoutePrefix string
	Routes      []RouteTmplParams
}

type RouteTmplParams struct {
	Method      string
	Path        string
	HandlerName string
}
