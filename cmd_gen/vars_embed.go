package cmd_gen

import _ "embed"

//go:embed template/api.template
var apiTmpl string

//go:embed template/logic.template
var logicTmpl string

//go:embed template/routes.template
var routesTmpl string

//go:embed template/router.template
var routerTmpl string

//go:embed template/main.template
var mainTmpl string

//go:embed template/config.template
var configTmpl string

//go:embed template/conf_yaml.template
var confYamlTmpl string


//go:embed template/dto.template
var dtoTmpl string

//go:embed template/ginx.template
var ginxTmpl string

//go:embed template/errorx_base.template
var errorxBaseTmpl string

//go:embed template/errorx_code.template
var errorxCodeTmpl string

//go:embed template/errorx_msg.template
var errorxMsgTmpl string