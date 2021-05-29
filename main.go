package main

import (
	"object-storage/routers"
	_ "object-storage/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, routers.Cors)
	beego.Run()
}
