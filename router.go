package fastgo

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"reflect"
)

var router = fasthttprouter.New()

func AddRouter(requestMethod, path string, controller ControllerInterface, action string)  {
	router.Handle(requestMethod, path, controllerHandle(controller, action))
}

func controllerHandle(c ControllerInterface, action string) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		vc := reflect.ValueOf(c)
		vt := reflect.TypeOf(c)
		method := vc.MethodByName(action)
		c.Init(ctx, vt.Name(), action)
		c.Before()
		method.Call(nil)
		c.After()
	})
}