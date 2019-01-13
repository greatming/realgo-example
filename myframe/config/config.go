package config

import (
	"realgo/conf"
	"fmt"
	"time"
	"realgo/lib/logger"
)

type LogConf struct {
	Path        string `toml:"path"`
	Level       int32    `toml:"level"`
	BackUpCount int    `toml:"backUpCount"`
}


type DBConf struct {
	Resourcesbackup  DBItemConf `toml:"resourcesbackup"`
}
type DBItemConf struct {
	Pool DBPoolConf `toml:"pool"`
	Info DBInfoConf `toml:"info"`
}
type DBPoolConf struct {
	MaxOpenConn int `toml:"MaxOpenConn"`
	MaxIdleConn int `toml:"MaxIdleConn"`
	MaxLifeTime int `toml:"MaxLifeTime"`
	ReadTimeout  int `toml:"ReadTimeout"`
	WriteTimeout int `toml:"WriteTimeout"`
}
type DBInfoConf struct {
	Host   []string `toml:"host"`
	User   string `toml:"user"`
	Pwd    string `toml:"pwd"`
	DBName string `toml:"dbname"`
}


var LogCfg LogConf
var DBCfg DBConf

func Init()  {
	initLog()
	InitDBConf()

}

func initLog()  {
	conf.ReadAppConfFile("log.toml", &LogCfg)
	logger.SetLevel(LogCfg.Level)
	logger.SetFile(LogCfg.Path)
	fmt.Println(LogCfg.Path)
	//日志切分，推荐使用下面这种EnableRotate，设置按小时切分，不推荐使用外部发信号处理日志切分
	logger.EnableRotate(time.Hour)
	//日志保留48个最新的日志文件，不设置为0，表示不删除过期的文件
	logger.SetbackupCount(LogCfg.BackUpCount)

}

func InitDBConf() error {
	if err := conf.ReadAppConfFile("db.toml", &DBCfg); err != nil {
		return err
	}
	fmt.Println(DBCfg)
	return  nil
}