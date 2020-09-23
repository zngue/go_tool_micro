package michttp

import (
	"encoding/json"
	"errors"
	"github.com/zngue/go_tool_micro/db"
	"github.com/zngue/go_tool_micro/michttp/httplib"
	"strconv"
	"strings"
	"time"
)

type MicHttpRequest struct {
	ServiceId string
	EndPoint  string
	Method    string
	Url       string
	Param     interface{}
	Header    map[string]string
	Timeout   time.Duration
}
type serverhost struct {
	ServiceName string
	Host        string
}
func (r *MicHttpRequest) MicRequest() (*httplib.HTTPRequest, error) {
	var url string
	if r.ServiceId == "" {
		return nil, errors.New("please input a service")
	}
	url = Getserverurl(r.ServiceId)
	if url != "" {
		url = Formaturl(url)
		r.Url = url + strings.TrimLeft(r.EndPoint, " ")
		req := r.Request()
		r, err := req.Response()
		if err != nil {
			return nil, err
		}
		if r.StatusCode == 200 {
			return req, nil
		} else {
			return nil, errors.New("service error statuscode is " + strconv.Itoa(r.StatusCode))
		}
	} else {
		return nil, errors.New("service not find")
	}
}

func (r *MicHttpRequest) GetRequesString() (string,error)  {
	req,err:=r.MicRequest()
	if err != nil {
		return "", err
	}
	return req.String()
}
func (r *MicHttpRequest) GetRequesBytes() ([]byte,error) {
	req,err:=r.MicRequest()
	if err != nil {
		return nil, err
	}
	return  req.Bytes()
}
func Getserverurl(servername string) string {
	servercemode := db.Config.HttpRequest.ServiceMode
	var url string
	if servercemode  {
		url = getlocalserver(servername)
	} else {
		url = getregserver(servername)

	}
	return url

}
func getregserver(key string) string {
	rdb := db.HttpRedisConn
	var url string
	register, err := rdb.SMembers("register").Result()
	if err != nil {

		return ""
	}
	for _, val := range register {
		server := &serverhost{}
		json.Unmarshal([]byte(val), &server)
		if server.ServiceName == key {
			url = server.Host
			break
		}
	}
	return url
}
func Formaturl(url string) string {
	if strings.Index(url, "http") < 0 {
		url = "http://" + url
	}
	last := url[(len(url) - 1):]
	if last != "/" {
		url += "/"
	}
	return url
}
func getlocalserver(key string) string {
	serviceList :=db.Config.ServiceList
	var url string
	for _,val:=range serviceList{
		if val.Name==key {
			return val.Url
		}
	}
	return url
}
func (r *MicHttpRequest) Request() *httplib.HTTPRequest {
	var req *httplib.HTTPRequest
	switch r.Method {
	case "post":
		req = httplib.Post(r.Url)
		if r.Param != nil {
			param := r.Param.(map[string]string)
			for name, val := range param {
				req.Param(name, val)

			}
		}
	case "jsonbody":
		req = httplib.Post(r.Url)
		req.JSONBody(r.Param)
	case "put":
		req = httplib.Put(r.Url)
	case "delete":
		req = httplib.Delete(r.Url)
	default:
		req = httplib.Get(r.Url)
		if r.Param != nil {
			param := r.Param.(map[string]string)
			for name, val := range param {
				req.Param(name, val)
			}
		}
	}
	if r.Header != nil {
		for name, val := range r.Header {

			req.Header(name, val)
		}
	}

	if r.Timeout == 0 {
		r.Timeout = 500
	}
	req.SetTimeout(r.Timeout*time.Millisecond, r.Timeout*time.Millisecond)

	return req
}
