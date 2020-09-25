package config

import (
	"crypto/rsa"
)

type Micro struct {
	ID            string        `json:"ID"`
	Host          string        `json:"host"`
	EndPoint      string        `json:"end_point"`
	IsMicroConfig bool          `json:"is_micro_config"`
	IsLocal       bool          `json:"IsLocal"`
	ServiceList   []ServiceList `json:"serviceList" `
}
type MicroResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       Config `json:"data"`
}
type Config struct {
	Mysql          Mysql         `json:"mysql" `
	Redis          Redis         `json:"redis" `
	System         System        `json:"system" `
	HttpRequest    HttpRequest   `json:"http_request" `
	HttpRedis      Redis         `json:"http_redis"`
	ServiceList    []ServiceList `json:"serviceList" `
	DefaultLoad    DefaultLoad   `json:"defaultLoad"`
	AliyunOss      AliyunOss     `json:"aliyunoss" `
	JWT            Jwt           `json:"jwt"`
	AliAppClient   AliAppClient  `json:"ali_app_client"`
	WeChat         WeChat        `json:"weChat"`
	HttpRedisModel bool          `json:"http_redis_model"`
}

//数据库配置信息
type Mysql struct {
	DBName       string `json:"dbName" form:"db_name" gorm:"column:db_name"`
	Prefix       string `json:"prefix" form:"prefix" gorm:"column:prefix"`
	Username     string `json:"username" form:"user_name" gorm:"column:user_name"`
	Password     string `json:"password" form:"password" gorm:"column:password"`
	Config       string `json:"config" form:"config" gorm:"column:config"`
	Host         string `json:"host" form:"host" gorm:"column:host"`
	Port         string `json:"port" form:"port" gorm:"column:port"`
	LogMode      bool   `json:"logMode" form:"log_mode" gorm:"column:log_mode"`
	MaxIdleConns int    `json:"maxIdleConns" form:"max_idle_conns" gorm:"column:max_idle_conns"`
	MaxOpenConns int    `json:"maxOpenConns" form:"max_open_conns" gorm:"column:max_open_conns"`
	Charset      string `json:"charset" form:"charset" gorm:"column:charset"`
	TimeStamp    bool   `json:"timeStamp" form:"time_stamp" gorm:"column:time_stamp"`
}
type AliAppClient struct {
	PartnerID  string          `json:"partner_id"` //合作者ID
	SellerID   string          `json:"seller_id"`
	AppID      string          `json:"app_id"` // 应用ID
	PrivateKey *rsa.PrivateKey `json:"private_key"`
	PublicKey  *rsa.PublicKey  `json:"public_key"`
}

//redis配置信息
type Redis struct {
	Host     string `json:"host" gorm:"column:host"  form:"host"`
	Port     string `json:"port"  gorm:"column:port"  form:"port"`
	Password string `json:"password" gorm:"column:password"  form:"password"`
	DBNum    int    `json:"dbNum" gorm:"column:db_num"  form:"db_num"`
	PoolSize int    `json:"poolSize" gorm:"column:pool_size"  form:"pool_size"`
}

//系统配置信息
type System struct {
	Port       string `json:"port" form:"port" gorm:"column:port"`
	SystemName string `json:"system_name" form:"system_name" gorm:"column:system_name"`
}

//http请求配置信息
type HttpRequest struct {
	ServiceMode bool `json:"serviceMode" gorm:"column:redis_db_num"  form:"redis_db_num"`
	RedisDBNum  int  `json:"redisDbNum" gorm:"column:service_mode" form:"service_mode"`
}

//http 服务请求配置
type ServiceList struct {
	Name string `json:"name" gorm:"column:name" form:"name"`
	Url  string `json:"url" gorm:"column:url" form:"url"`
}

//默认加载数据
type DefaultLoad struct {
	Mysql bool `json:"mysql" gorm:"column:mysql"  form:"mysql"`
	Redis bool `json:"redis" gorm:"column:redis" form:"redis"`
	HttpRedis bool `json:"http_redis" gorm:"column:http_redis" form:"http_redis"`
}
type AliyunOss struct {
	Accessid   string `json:"accessid" gorm:"auto_increment"   form:"accessid"`
	Accesskey  string `json:"accesskey" gorm:"column:accesskey" form:"accesskey"`
	Endpoint   string `json:"endpoint" gorm:"column:endpoint" form:"endpoint"`
	Bucket     string `json:"bucket" gorm:"column:bucket" form:"bucket"`
	Uploaddir  string `json:"uploaddir" gorm:"column:upload_dir" form:"upload_dir"`
	DomainName string `json:"domainName" gorm:"column:domain_name" form:"domain_name"`
	Ssl        bool   `json:"ssl" gorm:"column:ssl" form:"ssl"`
}
type Jwt struct {
	Secret     string `json:"secret" gorm:"column:secret" form:"secret"`
	ExpireTime int    `json:"expireTime" gorm:"column:expire_time" form:"expire_time"`
	Issuer     string `json:"issuer" gorm:"column:issuer" form:"issuer"`
	Subject    string `json:"subject" gorm:"column:subject" form:"subject"`
}
type WeChat struct {
	AppID          string `json:"appID" gorm:"column:app_id" form:"app_id"`
	AppSecret      string `json:"appSecret" gorm:"column:app_secret" form:"app_secret"`
	Token          string `json:"token" gorm:"column:token" form:"token"`
	EncodingAESKey string `json:"encodingAESKey" gorm:"column:encoding_aes_key" form:"encoding_aes_key"`
	PayMchID       string `json:"payMchID"  gorm:"column:pay_mch_id" form:"pay_mch_id"`
	PayNotifyURL   string `json:"payNotifyURL"  gorm:"column:pay_notify_url" form:"pay_notify_url"`
	PayKey         string `json:"payKey"  gorm:"column:pay_key" form:"pay_key"`
	RedisNum       int    `json:"redisNum" gorm:"column:redis_num" form:"redis_num"`
}

