package db

func CloseDB()  {
	if Config!=nil{
		if Config.DefaultLoad.Redis {
			RedisConn.Close()
		}
		if Config.DefaultLoad.Mysql {
			MysqlConn.Close()
		}
	}
}
