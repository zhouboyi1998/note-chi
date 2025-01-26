package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	// 新建 Chi 实例
	app := chi.NewRouter()

	// 启用日志中间件 (记录请求日志)
	app.Use(middleware.Logger)
	// 启用崩溃恢复中间件 (捕获程序中的 panic, 并恢复程序的正常运行, 防止程序因未处理的 panic 而崩溃)
	app.Use(middleware.Recoverer)

	// Hello World
	app.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Hello, " + name,
		})
	})

	// 启动服务
	log.Println("Server listen on :18082")
	http.ListenAndServe(":18082", app)
}
