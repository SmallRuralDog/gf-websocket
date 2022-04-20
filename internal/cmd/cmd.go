package cmd

import (
	"context"
	"gf-websocket/internal/packed/websocket"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf-websocket/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			//启动服务
			websocket.StartWebSocket(ctx)
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
				//注册路由
				group.ALL("/socket", websocket.WsPage)
			})
			s.Run()
			return nil
		},
	}
)
