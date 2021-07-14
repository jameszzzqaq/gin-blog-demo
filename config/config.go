package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	// base config
	RunMode string

	// server config
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	// app config
	PageSize  int
	JwtSecret string
)

func InitConfig() {

	var err error
	Cfg, err = ini.Load("app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("run_mode").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("http_port").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("read_timeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	PageSize = sec.Key("page_size").MustInt(10)
	JwtSecret = sec.Key("jwt_secret").MustString("!@)*#)!@U#@*!@!)")
}
