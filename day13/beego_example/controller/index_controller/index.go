package index_controler

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	//c.TplName="index.html"
	map1:=make(map[string]interface{})
	map1["name"]="alan"
	map1["message"]="helloworld"
	c.Data["json"]=&map1
	c.ServeJSON(true)
}
