package healthy

import (
	"context"
	"log"
	"sp1/app/service/sp1compare"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//Healthy is component
var Healthy = healthyAPI{}

type healthyAPI struct{}

// Index is a demonstration route handler for sp1 healthy check
func (*healthyAPI) Index(r *ghttp.Request) {
	v := g.View()
	v.AddPath("./template")
	v.SetDelimiters("${{", "}}")

	_, _ = v.Parse(context.TODO(), "header.html", g.Map{
		"action": "healthy",
	})

	res, err := v.Parse(context.TODO(), "./healthy/healthy_index.html", g.Map{
		"action": "healthy",
	})
	if err != nil {
		log.Println(err)
	}
	r.Response.Write(res)
}

// Healthyconfig is to compare config web
func (*healthyAPI) Healthyconfig(r *ghttp.Request) {
	v := g.View()
	v.AddPath("./template")
	v.SetDelimiters("${{", "}}")

	_, _ = v.Parse(context.TODO(), "header.html", g.Map{
		"action": "healthy",
	})

	res, err := v.Parse(context.TODO(), "./healthy/healthyconfig.html", g.Map{
		"action": "healthy",
	})
	if err != nil {
		log.Println(err)
	}
	r.Response.Write(res)
}

// GetConfig is to get config
func (*healthyAPI) GetConfig(r *ghttp.Request) {

	res := sp1compare.Compareini()

	r.Response.WriteJson(res)

}
