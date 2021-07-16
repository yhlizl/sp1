package mq

import (
	"fmt"
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

//ToggleMQ is to lock submit buttom for q
func ToggleMQ(item string) {
	timeLimit := g.Cfg().GetDuration("mq.timeout")
	status, starttime := GetStatus(item)
	//startTime, err := time.Parse("2006-01-02 15:04:05", starttime)
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", starttime, time.Local)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(time.Now().Sub(startTime), startTime)
	if (status == "0") && (time.Now().Sub(startTime) < timeLimit*time.Minute) {

	} else if status == "0" {
		EndJOB(item)
	} else if status == "1" {
		StartJOB(item)
	}
}

//GetStatus is to get status of mq
func GetStatus(item string) (string, string) {
	type Result struct {
	}

	subQuery := db.Model("mq_table").Fields("MAX(starttime)").Where("item", item)
	res, err := db.Model("mq_table").Fields("status,starttime").Where("item", item).Where("starttime=?", subQuery).All()
	if err != nil {
		log.Println(err)
	}
	status := res.Array("status")[0].String()

	starttime := res.Array("starttime")[0].String()

	return status, starttime

}

//EndJOB is to end job of mq
func EndJOB(item string) {
	now := time.Now().Format(time.RFC3339)
	starttime, err := db.Model("mq_table").Fields("MAX(starttime)").Where("item", item).Value()
	if err != nil {
		log.Println(err)
	}
	res, err := db.Model("mq_table").Data(g.Map{"item": item, "status": "1", "starttime": starttime, "endtime": now}).Save()
	if err != nil {
		log.Println(err)
	}
	log.Println("mq end finish:", item, res)
}

//StartJOB is to start job of mq
func StartJOB(item string) {
	now := time.Now().Format(time.RFC3339)

	res, err := db.Model("mq_table").Data(g.Map{"item": item, "status": "0", "starttime": now}).Save()
	if err != nil {
		log.Println(err)
	}
	log.Println("mq start finish:", item, res)
}
