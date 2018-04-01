package web

import "net/http"

func createRoute() {
	http.HandleFunc("/", index) //设置访问的路由
	http.HandleFunc("/acticle", getActicle)
}
