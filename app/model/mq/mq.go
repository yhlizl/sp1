package mq

import (
	"log"
	"time"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var (
	db gdb.DB
)

func init() {
	//TODO: rememmber to modify cfg
	g.Cfg().AddPath("/Users/royale/go/src/sp1/config/")
	db = g.DB("mq")
}

//GetStatus is to get status of mq
func GetStatus(item, params string) (string, string) {
	type Result struct {
	}

	subQuery := db.Model("mq_table").Fields("MAX(starttime)").Where("item", item).Where("params", params)
	res, err := db.Model("mq_table").Fields("status,starttime").Where("item", item).Where("starttime=?", subQuery).Where("params", params).All()
	if err != nil {
		log.Println(err)
	}
	status := "1"
	starttime := ""
	if !(res.IsEmpty()) {
		status = res.Array("status")[0].String()

		starttime = res.Array("starttime")[0].String()
	}
	return status, starttime

}

//EndJOB is to end job of mq
func EndJOB(item, params string) {
	now := time.Now().Format(time.RFC3339)
	starttime, err := db.Model("mq_table").Fields("MAX(starttime)").Where("item", item).Value()
	if err != nil {
		log.Println(err)
	}
	res, err := db.Model("mq_table").Data(g.Map{"item": item, "status": "1", "starttime": starttime, "params": params, "endtime": now}).Save()
	if err != nil {
		log.Println(err)
	}
	log.Println("mq end finish:", item, res)
}

//StartJOB is to start job of mq
func StartJOB(item, params string) {
	now := time.Now().Format(time.RFC3339)

	res, err := db.Model("mq_table").Data(g.Map{"item": item, "status": "0", "params": params, "starttime": now}).Save()
	if err != nil {
		log.Println(err)
	}
	log.Println("mq start finish:", item, res)
}
