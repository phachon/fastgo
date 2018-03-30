package fastgo

import (
	"github.com/spf13/viper"
	"os"
	"flag"
	"github.com/phachon/go-logger"
	"strings"
	"github.com/kataras/go-sessions"
	"time"
	"github.com/valyala/fasthttp"
)

var (
	flagConf = flag.String("conf", "config.toml", "please input conf path")
)

var (
	AppVersion = "v1.0"

	Conf = viper.New()

	Log = go_logger.NewLogger()

	AppPath = ""

	RootPath = ""

	ViewPath = ""

	StaticPath = ""

	TemplateSuffix = ".html"

	Session = &sessions.Sessions{}
)

// 启动初始化
func init()  {
	initFlag()
	initConfig()
	initLog()
	initPath()
	initSession()
}

// init flag
func initFlag() {
	flag.Parse()
}

// init config
func initConfig()  {

	if *flagConf == "" {
		Log.Error("config file not found!")
		os.Exit(1)
	}

	Conf.SetConfigType("toml")
	Conf.SetConfigFile(*flagConf)
	err := Conf.ReadInConfig()
	if err != nil {
		Log.Error("Fatal error config file: "+err.Error())
		os.Exit(1)
	}

	file := Conf.ConfigFileUsed()
	if(file != "") {
		Log.Info("Use config file: " + file)
	}
}

// init log
func initLog() {

	Log.Detach("console")

	consoleConfig := &go_logger.ConsoleConfig{
		Color: true,
	}
	Log.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, go_logger.NewConfigConsole(consoleConfig))
}

// init dir and path
func initPath() {
	AppPath, _ = os.Getwd()
	RootPath = strings.Replace(AppPath, "app", "", 1)
	SetViewsPath("views")
	SetStaticPath("static")
}

func initSession()  {
	Session = sessions.New(sessions.Config{
		Cookie: "cardssionid",
		Expires: time.Hour * 2,
		DisableSubdomainPersistence: false,
	})
}

func SetViewsPath(view string) {
	ViewPath = RootPath + "/" +view
}

func SetStaticPath(static string)  {
	StaticPath = RootPath + "/" +static
}

func SetTemplateSuffix(suffix string)  {
	TemplateSuffix = suffix
}

func Run() {

	AddRouter("GET", "/static/*path", NewController, "Static")
	// start listen server
	server := Conf.GetString("listen.server")
	Log.Info("start listen server "+server)
	err := fasthttp.ListenAndServe(server, router.Handler)
	if err != nil {
		Log.Info("listen server "+server+" error :"+err.Error())
	}
}