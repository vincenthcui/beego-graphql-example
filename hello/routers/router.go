package routers

import (
	"github.com/vincentcui/beego-graphql-example/hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.HelloController{})
}
