package router

import (
	"sp1/app/api/healthy"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.AddSearchPath("./public")
	s.AddStaticPath("/sp1/healthy", "./public")
	s.Group("/sp1/healthy", func(group *ghttp.RouterGroup) {
		group.GET("/", healthy.Healthy)
	})
}
