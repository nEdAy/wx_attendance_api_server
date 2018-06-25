package config

import (
	"time"

	"github.com/go-ini/ini"
	"github.com/rs/zerolog/log"
)

var (
	Cfg *ini.File

	App      app
	Server   server
	Path     path
	Database database
	WeChat   weChat
)

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	Cfg.BlockMode = false // if false, only reading, speed up read operations about 50-70% faster
	if err != nil {
		log.Fatal().Msgf("Fail to parse 'conf/app.ini': %v", err)
	}
	loadApp()
	loadServer()
	loadPath()
	loadDatabase()
	loadWeChat()
}

func loadApp() {
	sec, err := Cfg.GetSection("App")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'App': %v", err)
	}
	App.RunMode = sec.Key("RUN_MODE").In("debug", []string{"debug", "release"})
	App.JwtSecret = sec.Key("Jwt_Secret").MustString("123")
}

func loadServer() {
	sec, err := Cfg.GetSection("Server")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Server': %v", err)
	}
	Server.Protocol = sec.Key("PROTOCOL").In("http", []string{"http", "https"})
	Server.Host = sec.Key("HOST").MustString("127.0.0.1")
	Server.Port = sec.Key("PORT").MustInt(80)
	Server.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	Server.WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadPath() {
	sec, err := Cfg.GetSection("Paths")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Paths': %v", err)
	}
	Path.DataPath = sec.Key("DATA_PATH").MustString("./runtime")
	Path.LogPath = sec.Key("LOG_PATH").MustString("./runtime/log")
	Path.CachePath = sec.Key("CACHE_PATH").MustString("./runtime/cache")
	Path.CertFilePath = sec.Key("CERT_FILE_PATH").MustString("./data/ssl/ssl.cer")
	Path.KeyFilePath = sec.Key("KEY_FILE_PATH").MustString("./data/ssl/ssl.key")
}

func loadDatabase() {
	sec, err := Cfg.GetSection("Database")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Database': %v", err)
	}
	Database.Debug = sec.Key("DEBUG").MustBool(false)
	Database.Type = sec.Key("TYPE").MustString("mysql")
	Database.User = sec.Key("USER").MustString("root")
	Database.Password = sec.Key("PASSWORD").String()
	Database.Host = sec.Key("HOST").String()
	Database.Port = sec.Key("PORT").MustInt(3306)
	Database.Name = sec.Key("NAME").String()
	Database.TablePrefix = sec.Key("TABLE_PREFIX").String()
	Database.MaxIdleConns = sec.Key("MAX_IDLE_CONNS").MustInt(64)
	Database.MaxOpenConns = sec.Key("MAX_OPEN_CONNS").MustInt(24)
	Database.PingInterval = time.Duration(sec.Key("PING_INTERVAL").MustInt(30)) * time.Second
}

func loadWeChat() {
	sec, err := Cfg.GetSection("WeChat")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'WeChat': %v", err)
	}
	WeChat.CodeToSessionUrl = sec.Key("CODE_TO_SESSION_URL").String()
	WeChat.AppID = sec.Key("APP_ID").String()
	WeChat.AppSecret = sec.Key("APP_SECRET").String()
	WeChat.SessionMagicID = sec.Key("SESSION_MAGIC_ID").String()

}

type app struct {
	RunMode   string
	JwtSecret string
}

type server struct {
	Protocol     string
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type path struct {
	DataPath     string
	LogPath      string
	CachePath    string
	CertFilePath string
	KeyFilePath  string
}

type database struct {
	Debug        bool
	Type         string
	User         string
	Password     string
	Host         string
	Port         int
	Name         string
	TablePrefix  string
	MaxIdleConns int
	MaxOpenConns int
	PingInterval time.Duration
}

type weChat struct {
	CodeToSessionUrl string
	AppID            string
	AppSecret        string
	SessionMagicID   string
}
