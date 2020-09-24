package db

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zngue/go_tool/src/sign_chan"
	"github.com/zngue/go_tool_micro/config"
	"time"
)

type DB struct {
	HttpRedis config.Redis
	Redis config.Redis
	Mysql config.Mysql
	RedisCon *redis.Client
	HttpRedisCon *redis.Client
	MysqlCon *gorm.DB
}
func (d *DB) RedisConn()  {
	redisC:=d.Redis
	if &redisC!=nil {
		d.RedisCon = redis.NewClient(&redis.Options{
			Addr:         redisC.Host + ":"+redisC.Port,
			Password:     redisC.Password,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     30,
			PoolTimeout:  30 * time.Second,
			MinIdleConns: 10,
			DB:           redisC.DBNum,
		})
		pong, err := d.RedisCon.Ping().Result()
		if err != nil {
			sign_chan.SignLog("redis:错误",pong,err)
		}
	}
}
func (d *DB)HttpRedisConn()  {
	if !config.MicroConf.IsLocal {
		redisC:=d.HttpRedis
		d.HttpRedisCon = redis.NewClient(&redis.Options{
			Addr:         redisC.Host + ":"+redisC.Port,
			Password:     redisC.Password,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     30,
			PoolTimeout:  30 * time.Second,
			MinIdleConns: 10,
			DB:           redisC.DBNum,
		})
		pong, err := d.HttpRedisCon.Ping().Result()
		if err != nil {
			sign_chan.SignLog("http redis:错误",pong,err)
		}
	}
}
func (d *DB)RedisConnClose() (err error)  {
	defer func() {
		err =d.RedisCon.Close()
	}()
	return
}
func (d *DB) HttpRedisConnClose() (err error)  {
	defer func() {
		err =d.HttpRedisCon.Close()
	}()
	return
}
func (d *DB) MysqlConn(mysqlDB ...AutoDB)  {
	mysql :=d.Mysql
	dns:=mysql.Username+":"+mysql.Password+"@tcp("+mysql.Host+":"+mysql.Port+")/"+mysql.DBName+"?charset=utf8&parseTime=True&loc=Asia%2FShanghai"
	db, errDb := gorm.Open("mysql", dns)
	if errDb !=nil {
		sign_chan.SignLog(errDb)
		return
	}
	if mysql.LogMode {
		db = db.LogMode(mysql.LogMode)
	}
	if len(mysqlDB)>0 {
		for _,fn:=range mysqlDB{
			fn(db)
		}
	}
	db.DB().SetMaxIdleConns(mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysql.MaxOpenConns)
	if mysql.TimeStamp {
		db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
		db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	}
	if mysql.Prefix!="" {
		gorm.DefaultTableNameHandler= func(db *gorm.DB,defaultTableName string) string{
			return mysql.Prefix+defaultTableName
		}
	}
	d.MysqlCon = db
	return
}
func (d *DB) MysqlConnClose()  (err error)  {
	defer func() {
		err=d.MysqlCon.Close()
	}()
	return
}
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}
// // 注册新建钩子在持久化之前
func updateTimeStampForCreateCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}
