package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zngue/go_tool/src/fun/file"
	"github.com/zngue/go_tool/src/log"
	"github.com/zngue/go_tool/src/sign_chan"
	"github.com/zngue/go_tool_micro/michttp/httplib"
	"io/ioutil"
)
const (
	ConfigJson = ""
	ConfigYaml = "config.yaml"
	MicroYaml = "micro.yaml"
)
var MicroConf *Micro
func JsonToStruck() *Config  {
	var config Config
	data, err := ioutil.ReadFile(ConfigJson)
	if err != nil {
		return nil
	}
	jerr:=json.Unmarshal(data,&config)
	if jerr!=nil {
		log.LogInfo(jerr)
		return nil
	}
	return  &config
}
func YamlToStruck()(configinfo *Config) {
	var config Config
	if !file.FileExist(ConfigYaml){
		sign_chan.SignLog(errors.New("config.yaml is not Exist "))
		return
	}
	v := viper.New()
	v.SetConfigFile(ConfigYaml)
	err := v.ReadInConfig()
	if err != nil {
		sign_chan.SignLog(err)
		return
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&config); err != nil {
			sign_chan.SignLog(err)
		}
		configinfo = &config
	})
	if err := v.Unmarshal(&config); err != nil {
		sign_chan.SignLog(err)
		return
	}
	configinfo = &config
	return
}
func MicroConfig(){
	if !file.FileExist(MicroYaml){
		sign_chan.SignLog(errors.New("micro.yaml is not Exist "))
		return
	}
	v := viper.New()
	v.SetConfigFile(MicroYaml)
	err := v.ReadInConfig()
	if err != nil {
		sign_chan.SignLog(err)
		return
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&MicroConf); err != nil {
			sign_chan.SignLog(err)
		}
	})
	if err := v.Unmarshal(&MicroConf); err != nil {
		sign_chan.SignLog(err)
		return
	}
	return
}
func  MicroHttpRequest() *Config  {
	req:=httplib.Get(MicroConf.Host+MicroConf.EndPoint)
	req.Param("id",MicroConf.ID)
	s,rErr:=req.Bytes()
	if rErr != nil {
		sign_chan.SignLog(rErr)
		return nil
	}
	var m MicroResponse
	err:=json.Unmarshal(s,&m)
	if err != nil {
		sign_chan.SignLog(err)
		return nil
	}
	if m.StatusCode==200 {
		return &m.Data
	}
	sign_chan.SignLog(m.Message,m.Data)
	return nil
}