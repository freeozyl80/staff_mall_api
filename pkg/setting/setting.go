package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	AppVer    string
	JwtSecret string
}

type Server struct {
	Host         string
	DevMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Xlsx struct {
	UserXlsx string
}

type Crypto struct {
	Seed1 string
	Seed2 string
}

var ServerSetting = &Server{}
var AppSetting = &App{}
var DatabaseSetting = &Database{}
var XlsxSetting = &Xlsx{}
var CryptoSetting = &Crypto{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/app.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)
	mapTo("xlsx", XlsxSetting)
	mapTo("crypto", CryptoSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo Setting err: %v", err)
	}
}
