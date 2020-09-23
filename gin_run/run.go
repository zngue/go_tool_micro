package gin_run

import (
	"context"
	"github.com/zngue/go_tool/src/sign_chan"
	"github.com/zngue/go_tool_micro/db"
)

func GinRun(fnRouter RouterFun,mysqlDB ...db.AutoDB)  {
	port := db.Config.System.Port
	if port=="" {
		sign_chan.SignLog("服务器端口不存在请配置端口")
		return
	}
	http,_:=HttpRouterServe(port,fnRouter)
	go HttpRun(func() error {
		return http.ListenAndServe()
	})
	go db.InitDB(mysqlDB...)
	sign_chan.ListClose(func(ctx context.Context) error {
		return http.Shutdown(ctx)
	})
}
