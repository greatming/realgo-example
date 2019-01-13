package model

import (
	"myframe/config"
	"myframe/utils"
	"realgo/lib/logger"
	"sync"
)

var (
	BackUpDB *utils.DBHandler
)
func GetBackUpDBHandler() *utils.DBHandler {
	if BackUpDB != nil{
		return  BackUpDB
	}
	m := new(sync.RWMutex)

	m.Lock()
	log := logger.New()
	conf := config.DBCfg.Resourcesbackup
	log.SetBaseInfo("model", "ResourcesbackupDB")
	BackUpDB = utils.NewDBHandler(conf.Pool, conf.Info, log)
	m.Unlock()
	return BackUpDB
}
