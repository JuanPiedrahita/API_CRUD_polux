package main

import (
	_ "github.com/JuanPiedrahita/Polux/apiPolux/routers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@127.0.0.1:5432/Polux?sslmode=disable&search_path=polux")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
        AllowOrigins: []string{"http://localhost*"},
        AllowMethods: []string{"PUT", "PATCH"},
        AllowHeaders: []string{"Origin"},
        ExposeHeaders: []string{"Content-Length"},
        AllowCredentials: true,
 }))

	beego.Run()
}
