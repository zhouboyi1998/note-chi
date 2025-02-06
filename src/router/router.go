package router

import (
	"github.com/go-chi/chi/v5"
	"note-chi/src/controller"
)

// RegisterRouter 注册路由
func RegisterRouter(c *chi.Mux) {
	// 新建命令路由组
	c.Route("/command", func(r chi.Router) {
		// 添加命令相关路由
		r.Get("/:commandId", controller.One)
		r.Get("/", controller.List)
		r.Post("/", controller.Insert)
		r.Post("/batch", controller.InsertBatch)
		r.Put("/", controller.Update)
		r.Put("/batch", controller.UpdateBatch)
		r.Delete("/:commandId", controller.Delete)
		r.Delete("/batch", controller.DeleteBatch)
		r.Get("/select/:commandName", controller.Select)
		r.Get("/name-list", controller.NameList)
	})
}
