package app

import (
	"github.com/snail007/go-activerecord/mysql"
	"github.com/phachon/fastgo"
	"os"
)

var (
	Conf = fastgo.Conf
	Log = fastgo.Log
)

var G *mysql.DBGroup

func init()  {
	initDatabase()
}

// init database
func initDatabase() {

	G = mysql.NewDBGroup("default")

	host := Conf.GetString("db.host")
	port := Conf.GetInt("db.port")
	dbName := Conf.GetString("db.name")
	user := Conf.GetString("db.user")
	pass := Conf.GetString("db.pass")
	maxIdle := Conf.GetInt("db.conn_max_idle")
	maxConn := Conf.GetInt("db.conn_max_connection")
	dbTablePrefix := Conf.GetString("db.table_prefix")

	cfg := mysql.NewDBConfigWith(host, port, dbName, user, pass)
	cfg.SetMaxIdleConns = maxIdle
	cfg.SetMaxOpenConns = maxConn
	cfg.TablePrefix = dbTablePrefix
	cfg.TablePrefixSqlIdentifier = "__PREFIX__"
	err := G.Regist("default", cfg)
	if err != nil {
		Log.Error("register db error: " + err.Error())
		os.Exit(100)
	}

	Log.Info("database conn success")
}