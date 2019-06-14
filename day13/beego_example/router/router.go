package router

import (
	"github.com/astaxie/beego"
	index_controler "oldboy/day13/beego_example/controller/index_controller"
)

func init() {
	beego.Router("/index",&index_controler.IndexController{},"*:Index")
}
