package utils

import (
	"time"
	"database/sql"
	"fmt"
	"math/rand"
	_"github.com/go-sql-driver/mysql"
	"myframe/config"
	"realgo/lib/logger"
)

type DBHandler struct {
	Info    config.DBInfoConf
	PoolCfg config.DBPoolConf
	Handler []*sql.DB
	Logger *logger.Logger
}
func (db *DBHandler) getHandler(user, pwd, host, dbname string) (*sql.DB, error) {
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?readTimeout=%dms&writeTimeout=%dms",
		user, pwd, host, dbname, db.PoolCfg.ReadTimeout, db.PoolCfg.WriteTimeout)
	handler, err := sql.Open("mysql", str)
	if err != nil {
		db.Logger.Warn("Open " + str + " failed, err: " + err.Error())
		fmt.Println("Open " + str + " failed, err: " + err.Error())
		return nil, err
	}
	handler.SetMaxOpenConns(db.PoolCfg.MaxOpenConn)
	handler.SetMaxIdleConns(db.PoolCfg.MaxIdleConn)
	handler.SetConnMaxLifetime(time.Duration(db.PoolCfg.MaxLifeTime) * time.Second)
	db.Logger.Info4("Open DB: %s with pool [MaxOpenConns:%d] [MaxIdleConns:%d] [MaxLifeTime:%d]",
		str, db.PoolCfg.MaxOpenConn, db.PoolCfg.MaxIdleConn, db.PoolCfg.MaxLifeTime)
	//测试是否能连上DB
	if err := handler.Ping(); err != nil {
		return nil, err
	}
	return handler, err
}
func (db *DBHandler) GetInstance() (*sql.DB, error) {
	if len(db.Handler) == 0 {
		return nil, fmt.Errorf("no db handler, maybe forgot init DBHandler")
	}
	size := len(db.Handler)
	index := rand.Intn(size)
	var err error
	for i := 0; i < size; i++ {
		if db.Handler[index] != nil {
			return db.Handler[index], nil
		}
		//db Handle init failed at first, init Handle
		db.Handler[index], err = db.getHandler(db.Info.User, db.Info.Pwd, db.Info.Host[index], db.Info.DBName)
		if err != nil {
			index = (index + 1) % size
		} else {
			return db.Handler[index], nil
		}
	}
	return nil, fmt.Errorf("can not find any host connect success")
}
func NewDBHandler(poolCfg config.DBPoolConf, info config.DBInfoConf, log *logger.Logger) *DBHandler {
	db := new(DBHandler)
	db.Info = info
	db.PoolCfg = poolCfg
	db.Logger = log
	db.Handler = make([]*sql.DB, len(db.Info.Host))
	for i := range db.Handler {
		db.Handler[i], _ = db.getHandler(db.Info.User, db.Info.Pwd, db.Info.Host[i], db.Info.DBName)
	}
	return db
}
