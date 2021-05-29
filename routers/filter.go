package routers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web/context"
)

var success = []byte("SUPPORT OPTIONS")

func Cors(ctx *context.Context) {
	origin := ctx.Input.Header("Origin")
	ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
	ctx.Output.Header("Access-Control-Max-Age", "3600")
	ctx.Output.Header("Access-Control-Allow-Headers", "X-Custom-Header,accept,Content-Type,Access-Token")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Allow-Origin", origin)
	if ctx.Input.Method() == http.MethodOptions {
		ctx.Output.SetStatus(http.StatusOK)
		_ = ctx.Output.Body(success)
	}
}
