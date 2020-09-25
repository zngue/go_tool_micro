package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/sign_chan"
	"github.com/zngue/go_tool_micro/config"
	"sync"
)
var (
	Config        *config.Config
	MysqlConn     *gorm.DB
	RedisConn     *redis.Client
	HttpRedisConn *redis.Client
)
func init() {
	if Config == nil {
		config.MicroConfig()
		if config.MicroConf == nil {
			sign_chan.SignLog(" micro.yaml 配置文件加载失败... miro")
		}
		sd:=config.MicroConf
		fmt.Println(sd)
		if config.MicroConf.IsLocal {
			fmt.Println(1111)
			Config=config.YamlToStruck()
		}else{
			Config = config.MicroHttpRequest()
		}
		if Config == nil {
			sign_chan.SignLog("配置文件加载失败... config")
		}
	}
}
type AutoDB func(db *gorm.DB)
func InitDB(mysqlDbd ...AutoDB) {
	if Config == nil {
		sign_chan.SignLog("配置文件加载失败...")
		return
	}
	load := Config.DefaultLoad
	dbConn := DB{
		Redis:     Config.Redis,
		HttpRedis: Config.HttpRedis,
		Mysql:     Config.Mysql,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if load.Redis {
			dbConn.RedisConn()
			dbConn.HttpRedisConn()
		}
		RedisConn = dbConn.RedisCon
		HttpRedisConn = dbConn.HttpRedisCon
	}()
	go func() {
		defer wg.Done()
		if load.Mysql {
			dbConn.MysqlConn(mysqlDbd...)
		}
		MysqlConn = dbConn.MysqlCon
	}()
	wg.Wait()
}

//关闭连接池
func ConnClose() {
	if Config != nil {
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			defer wg.Done()
			if Config.DefaultLoad.Mysql && MysqlConn != nil { //关闭数据库连接池
				MysqlConn.Close()
			}
		}()
		go func() {
			defer wg.Done()
			if Config.DefaultLoad.Redis && RedisConn != nil {
				RedisConn.Close()
			}
		}()
		go func() {
			defer wg.Done()
			if Config.DefaultLoad.Redis && HttpRedisConn != nil {
				HttpRedisConn.Close()
			}
		}()
		wg.Wait()
	}
}
