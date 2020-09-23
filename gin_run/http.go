package gin_run

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/sign_chan"
	"log"
	"net/http"
	"time"
)

type httpFunc func(server *http.Server)
func Http(engine *gin.Engine,port string ) *http.Server {
	server := &http.Server{
		Addr:         ":"+port,
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("|-----------------------------------|")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|----------Port:" + port + "----------------|")
	fmt.Println("|-zngue微服务博客地址blog.zngue.com-|")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|-----------------------------------|")
	return server
}
func HttpRun(fns func() error )  {
	if err :=fns();err!=nil{
		sign_chan.SignLog(err)
		log.Fatalln(fmt.Sprintf("服务启动失败"))
	}else{
		log.Println(fmt.Sprintf("服务启动成功"))
	}
}
