package healthy

import (
	"context"
	"fmt"
	"log"
	"path"
	"sp1/app/model/mq"
	"sp1/app/service/sp1compare"
	"time"

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

	res := []sp1compare.Configdata{}
	rstatus := ""
	item := r.GetString("item")
	fab := r.GetArray("fab")
	phase := r.GetArray("phase")
	params := ""
	for _, f := range fab {
		params += f
	}
	for _, p := range phase {
		params += p
	}

	log.Println("get params:", item, fab, phase)
	r.Session.Set(params, "start")
	timeLimit := g.Cfg().GetDuration("mq.timeout")
	status, starttime := mq.GetStatus(item, params)
	//startTime, err := time.Parse("2006-01-02 15:04:05", starttime)
	startTime := time.Now()
	if starttime != "" {
		err := *new(error)
		startTime, err = time.ParseInLocation("2006-01-02 15:04:05", starttime, time.Local)
		if err != nil {
			log.Println(err)
		}
	} else {
		status = "1"
	}

	if (status == "0") && (time.Now().Sub(startTime) > timeLimit*time.Minute) {
		mq.StartJOB(item, params)
		_ = sp1compare.Compareini(fab, phase)
		dir := path.Join(sp1compare.Config.Configpath, params)
		res = sp1compare.ReadFromConfig(dir)
		mq.EndJOB(item, params)
		rstatus = "timeout rerun"
		err := r.Session.Remove(params)
		if err != nil {
			log.Println(err)
		}

	} else if status == "0" {
		for {
			statusNew, _ := mq.GetStatus(item, params)
			if statusNew == "1" {
				break
			}
			time.Sleep(5 * time.Second)
		}
		dir := path.Join(sp1compare.Config.Configpath, params)
		res = sp1compare.ReadFromConfig(dir)
		rstatus = "last run is still running"

	} else if status == "1" {

		mq.StartJOB(item, params)
		_ = sp1compare.Compareini(fab, phase)
		dir := path.Join(sp1compare.Config.Configpath, params)
		res = sp1compare.ReadFromConfig(dir)
		mq.EndJOB(item, params)

		rstatus = "start running"
		err := r.Session.Remove(params)
		if err != nil {
			log.Println(err)
		}
	}

	result := g.Map{
		"status": rstatus,
		"data":   res,
	}
	r.Response.WriteJson(result)

}

// ForceGetConfig is forcing to get config
func (*healthyAPI) ForceGetConfig(r *ghttp.Request) {
	res := []sp1compare.Configdata{}
	rstatus := ""
	item := r.GetString("item")
	fab := r.GetArray("fab")
	phase := r.GetArray("phase")
	log.Println("get params:", item, fab, phase)

	// timeLimit := g.Cfg().GetDuration("mq.timeout")
	// status, starttime := mq.GetStatus(item)
	// //startTime, err := time.Parse("2006-01-02 15:04:05", starttime)
	// startTime, err := time.ParseInLocation("2006-01-02 15:04:05", starttime, time.Local)
	// if err != nil {
	// 	log.Println(err)
	// }
	params := ""
	for _, f := range fab {
		params += f
	}
	for _, p := range phase {
		params += p
	}
	mq.StartJOB(item, params)
	_ = sp1compare.Compareini(fab, phase)
	dir := path.Join(sp1compare.Config.Configpath, params)
	res = sp1compare.ReadFromConfig(dir)
	fmt.Println(res)
	mq.EndJOB(item, params)
	rstatus = "force run"

	//xfer data to html
	result := g.Map{
		"status": rstatus,
		"data":   res,
	}
	r.Response.WriteJson(result)

}
