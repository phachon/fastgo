package fastgo

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"reflect"
	"strings"
)

type Router struct {
	fastHttpRouter *fasthttprouter.Router
}

func NewRouter() *Router {
	return &Router{
		fastHttpRouter: fasthttprouter.New(),
	}
}

func (r *Router) Add(requestMethod, path string, controller ControllerInterface, action string)  {
	r.fastHttpRouter.Handle(requestMethod, path, r.controllerHandle(controller, action))
}

// GET Request
func (r *Router) GET(path string, controller ControllerInterface, action string) {
	r.Add("GET", path, controller, action)
}

// HEAD Request
func (r *Router) HEAD(path string, controller ControllerInterface, action string) {
	r.Add("HEAD", path, controller, action)
}

// OPTIONS Request
func (r *Router) OPTIONS(path string, controller ControllerInterface, action string) {
	r.Add("OPTIONS", path, controller, action)
}

// POST Request
func (r *Router) POST(path string, controller ControllerInterface, action string) {
	r.Add("POST", path, controller, action)
}

// PUT Request
func (r *Router) PUT(path string, controller ControllerInterface, action string) {
	r.Add("PUT", path, controller, action)
}

// PATCH Request
func (r *Router) PATCH(path string, controller ControllerInterface, action string) {
	r.Add("PATCH", path, controller, action)
}

// DELETE Request
func (r *Router) DELETE(path string, controller ControllerInterface, action string) {
	r.Add("DELETE", path, controller, action)
}

// GET and POST
func (r *Router) Request(requestMethod, path string, controller ControllerInterface, action string)  {
	r.Add("GET", path, controller, action)
	r.Add("POST", path, controller, action)
}

func (r *Router) controllerHandle(c ControllerInterface, action string) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		vc := reflect.ValueOf(c)
		vt := reflect.TypeOf(c)
		if !vc.IsValid() || vt == nil {
			ctx.SetStatusCode(404)
			ctx.SetBodyString("404 Not found")
			return
		}
		controllerName := ""
		list := strings.Split(vt.String(), ".")
		if len(list) >= 1 {
			controllerName = list[1]
		}else {
			controllerName = list[0]
		}
		c.Init(ctx, controllerName, action)
		method := vc.MethodByName(action)
		if !method.IsValid() {
			ctx.SetStatusCode(404)
			ctx.SetBodyString("404 Not found")
			return
		}
		c.Before()
		method.Call(nil)
		c.After()
	})
}