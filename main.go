package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"note-chi/src/application"
	"note-chi/src/router"
)

func main() {
	// 新建 Chi 实例
	app := chi.NewRouter()

	// 启用日志中间件 (记录请求日志)
	app.Use(middleware.Logger)
	// 启用崩溃恢复中间件 (捕获程序中的 panic, 并恢复程序的正常运行, 防止程序因未处理的 panic 而崩溃)
	app.Use(middleware.Recoverer)

	// 注册路由
	router.RegisterRouter(app)

	// 启动服务
	addr := application.App.Server.Host + ":" + application.App.Server.Port
	log.Println("Server listen on " + addr)
	http.ListenAndServe(addr, app)
}
