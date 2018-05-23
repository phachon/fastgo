package fastgo

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"encoding/json"
	"html/template"
	"mime/multipart"
	"path"
	"strings"
	"bytes"
	"github.com/phachon/fasthttpsession"
	"time"
)

type ControllerInterface interface {
	Init(ctx *fasthttp.RequestCtx, controllerName string, actionName string)
	Before()
	After()
}

func NewController() *Controller {
	return &Controller{}
}

type Controller struct {
	Ctx *fasthttp.RequestCtx
	ControllerName string
	ActionName string
	Data map[string]interface{}
	Session fasthttpsession.SessionStore
}

// init controller
func (controller *Controller) Init(ctx *fasthttp.RequestCtx, controllerName string, actionName string) {
	controller.Ctx = ctx
	controller.ControllerName = controllerName
	controller.ActionName = actionName
	controller.Data = map[string]interface{}{}

	var err error
	controller.Session, err = Session.Start(ctx)
	if err != nil {
		panic("session start error, "+err.Error())
	}
}

// controller before
func (controller *Controller) Before() {
	//todo controller before
}

// controller after
func (controller *Controller) After() {
	// todo controller after
	err := controller.Session.Save(controller.Ctx)
	if err != nil {
		panic("session save error, "+err.Error())
	}
}

// render view response
func (controller *Controller) Render(tpl string) {
	controller.Ctx.SetContentType("text/html;charset=utf-8")
	t, err := template.ParseFiles(ViewPath +"/"+ tpl + TemplateSuffix)
	if err != nil {
		Log.Errorf("template %s parese error %s", ViewPath+"/"+tpl, err.Error())
		controller.Ctx.SetBodyString("template parese error, "+err.Error())
		return
	}
	t.Execute(controller.Ctx.Response.BodyWriter(), controller.Data)
}

// layout render view response
func (controller *Controller) LayoutRender(layout string, tpl string) {
	controller.Ctx.SetContentType("text/html;charset=utf-8")

	var buf bytes.Buffer
	t, err := template.ParseFiles(ViewPath +"/"+ tpl + TemplateSuffix)
	if err != nil {
		Log.Errorf("template %s parese error %s", ViewPath+"/"+tpl, err.Error())
		controller.Ctx.SetBodyString("template parese error, "+err.Error())
		return
	}
	t.Execute(&buf, controller.Data)
	controller.Data["LayoutContent"] = template.HTML(buf.String())

	t, err = template.ParseFiles(ViewPath +"/"+ layout + TemplateSuffix)
	if err != nil {
		Log.Errorf("template %s parese error %s", ViewPath+"/"+layout, err.Error())
		controller.Ctx.SetBodyString("template parese error, " +err.Error())
		return
	}
	t.Execute(controller.Ctx.Response.BodyWriter(), controller.Data)
}

// return json
func (controller *Controller) ReturnJson(body interface{}) {
	controller.Ctx.SetContentType("application/json;charset=utf-8")
	jsonByte, err := json.Marshal(body)
	if err != nil {
		controller.Ctx.SetBodyString(err.Error())
	} else {
		controller.Ctx.SetBody(jsonByte)
	}
}

// return error
func (controller *Controller) ReturnError(status int, message string) {
	controller.Ctx.SetStatusCode(status)
	controller.Ctx.SetBodyString(message)
}

// get request ctx string
func (controller *Controller) GetString(key string, def string) string {
	if string(controller.Ctx.FormValue(key)) == "" {
		return def
	}else {
		return string(controller.Ctx.FormValue(key))
	}
}

// get request ctx bool
func (controller *Controller) GetBool(key string) bool {
	str := string(controller.Ctx.FormValue(key))
	if str == "1" {
		return true
	}
	return false
}

// get request ctx int
func (controller *Controller) GetInt(key string, def int) (int, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i, err := strconv.Atoi(str)
	return i, err
}

// get request ctx int8
func (controller *Controller) GetInt8(key string, def int8) (int8, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(str, 10, 8)
	return int8(i64), err
}

// get request ctx uint8
func (controller *Controller) GetUInt8(key string, def uint8) (uint8, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(str, 10, 8)
	return uint8(i64), err
}

// get request ctx int16
func (controller *Controller) GetInt16(key string, def int16) (int16, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(str, 10, 8)
	return int16(i64), err
}

// get request ctx uint16
func (controller *Controller) GetUInt16(key string, def uint16) (uint16, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(str, 10, 8)
	return uint16(i64), err
}

// get request ctx int32
func (controller *Controller) GetInt32(key string, def int32) (int32, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(str, 10, 8)
	return int32(i64), err
}

// get request ctx uint32
func (controller *Controller) GetUInt32(key string, def uint32) (uint32, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(str, 10, 8)
	return uint32(i64), err
}

// get request ctx int64
func (controller *Controller) GetInt64(key string, def int64) (int64, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(str, 10, 8)
	return i64, err
}

// get request ctx uint64
func (controller *Controller) GetUInt64(key string, def uint64) (uint64, error) {
	str := string(controller.Ctx.FormValue(key))
	if str == "" {
		return def, nil
	}
	i64, err := strconv.ParseInt(str, 10, 8)
	return uint64(i64), err
}

// get request content text float64
func (controller *Controller) GetCtxFloat64(key string, def float64) (float64, error) {
	str := string(controller.Ctx.FormValue(key))
	i, err := strconv.Atoi(str)
	return float64(i), err
}

// get form file by key
func (controller *Controller) GetFile(key string) (*multipart.FileHeader, error) {
	return controller.Ctx.FormFile(key)
}

// request is ajax
func (controller *Controller) IsAjax() bool {
	return string(controller.Ctx.Request.Header.Peek("X-Requested-With")) == "XMLHttpRequest"
}

func (controller *Controller) SetCookie(name string, value string, others ...interface{}) {
	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)

	cookie.SetKey(name)
	cookie.SetHTTPOnly(true)
	// 1. expires
	if len(others) > 0 {
		var expires int64
		switch v := others[0].(type) {
		case int:
			expires = int64(v)
		case int32:
			expires = int64(v)
		case int64:
			expires = v
		}
		switch {
		case expires > 0:
			cookie.SetExpire(time.Now().Add(time.Duration(expires) * time.Second))
		case expires < 0:
			cookie.SetExpire(fasthttp.CookieExpireUnlimited)
		}
	}
	// 2. path
	if len(others) > 1 {
		if v, ok := others[1].(string); ok && len(v) > 0 {
			cookie.SetPath(v)
		}
	} else {
		cookie.SetPath("/")
	}
	// 3. domain
	if len(others) > 2 {
		if v, ok := others[2].(string); ok && len(v) > 0 {
			cookie.SetDomain(v)
		}
	}
	// 4. secure
	if len(others) > 3 {
		if v, ok := others[2].(bool); ok && v && controller.Ctx.IsTLS() {
			controller.Ctx.IsTLS()
		}
	}

	cookie.SetValue(value)
	controller.Ctx.Response.Header.SetCookie(cookie)
}

func (controller *Controller) GetCookie(key string) []byte {
	return controller.Ctx.Request.Header.Cookie(key)
}

func (controller *Controller) GetCookieString(key string) string {
	return string(controller.Ctx.Request.Header.Cookie(key))
}

func (controller *Controller) UserAgent() string {
	return string(controller.Ctx.Request.Header.UserAgent())
}

// static file
func (controller *Controller) Static() {

	baseName := path.Base(StaticPath)
	dir := strings.TrimRight(StaticPath, baseName)
	fsHandler := fasthttp.FSHandler(dir, 0)
	fsHandler(controller.Ctx)
}