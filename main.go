package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	_ "idyllim.com/blog/routers"
	"net/http"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles("/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func main() {

	fmt.Print("start web ...")
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateRight = "<<<"
	beego.Run()

}
