package fastgo

import (
	"github.com/spf13/viper"
	"os"
	"github.com/phachon/go-logger"
	"strings"
	"github.com/valyala/fasthttp"
	"path"
	"github.com/phachon/fasthttpsession"
)

var (
	Version = "v0.1"

	Conf = viper.New()

	Log = go_logger.NewLogger()

	AppPath = ""

	RootPath = ""

	ViewPath = ""

	StaticPath = ""

	TemplateSuffix = ".html"

	Session = fasthttpsession.NewSession(fasthttpsession.NewDefaultConfig())

	Route = NewRouter()
)

// start init
func init()  {
	initPath()
	initDefaultConf()
	initSession()
}

// init dir and path
func initPath() {
	AppPath, _ = os.Getwd()
	RootPath = strings.Replace(AppPath, "app", "", 1)
	SetViewsPath("views")
	SetStaticPath("static")
}

// init config
func initDefaultConf()  {
	Conf.SetConfigType("toml")
	Conf.SetConfigFile("config.toml")
}

// init log
func initLog() {

	Log.Detach("console")

	// console adapter config
	consoleLevelStr := Conf.GetString("log.console.level")
	consoleConfig := &go_logger.ConsoleConfig{
		Color: Conf.GetBool("log.console.color"), // show text color
		JsonFormat: Conf.GetBool("log.console.jsonFormat"), // text json format
		Format: Conf.GetString("log.console.format"),
	}
	Log.Attach("console", Log.LoggerLevel(consoleLevelStr), consoleConfig)

	// file adapter config
	fileLevelStr := Conf.GetString("log.file.level")
	levelFilenameConf := Conf.GetStringMapString("log.file.levelFilename")
	levelFilename := map[int]string{}
	if len(levelFilenameConf) > 0 {
		for levelStr, levelFile := range levelFilenameConf {
			level := Log.LoggerLevel(levelStr)
			levelFilename[level] = levelFile
		}
	}
	fileConfig := &go_logger.FileConfig{
		Filename:  Conf.GetString("log.file.filename"),
		LevelFileName : levelFilename,
		MaxSize: Conf.GetInt64("log.file.maxSize"),
		MaxLine: Conf.GetInt64("log.file.maxLine"),
		DateSlice: Conf.GetString("log.file.dateSlice"),
		JsonFormat: Conf.GetBool("log.file.jsonFormat"),
		Format: Conf.GetString("log.file.format"),
	}
	Log.Attach("file", Log.LoggerLevel(fileLevelStr), fileConfig)
}

func initSession()  {}

func SetViewsPath(view string) {
	ViewPath = RootPath + "/" +view
}

func SetStaticPath(static string)  {
	StaticPath = RootPath + "/" +static
}

func SetTemplateSuffix(suffix string)  {
	TemplateSuffix = suffix
}

func ListenAndServe(addr string, route *Router)  {
	route.Add("GET", "/"+path.Base(StaticPath)+"/*path", NewController(), "Static")
	Log.Infof("start listen server %s", addr)
	err := fasthttp.ListenAndServe(addr, route.fastHttpRouter.Handler)
	if err != nil {
		Log.Infof("listen server %s error: %s", addr, err.Error())
	}
}

func checkConf()  {
	err := Conf.ReadInConfig()
	if err != nil {
		Log.Error("Fatal error config file: "+err.Error())
		os.Exit(1)
	}

	file := Conf.ConfigFileUsed()
	if file != "" {
		Log.Info("Use config file: " + file)
	}
}
func Run() {
	checkConf()
	initLog()
	// start listen server
	server := Conf.GetString("listen.server")
	ListenAndServe(server, Route)
}