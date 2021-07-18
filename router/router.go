package router

import (
	"sp1/app/api/healthy"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.SetSessionCookieMaxAge(1 * time.Hour)
	s.AddSearchPath("./public")
	s.AddStaticPath("/sp1/healthy", "./public")
	s.Group("/sp1/healthy", func(group *ghttp.RouterGroup) {
		group.ALL("/", healthy.Healthy)
	})
}
