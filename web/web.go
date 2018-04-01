package web

import (
	"document/config"
	"log"
	"net/http"
)

// Run 启动web
func Run() {
	bind, err := config.GetString("bind")
	port, err := config.GetString("port")
	if err != nil {
		log.Printf("obtain port error:%v\n", err)
		return
	}
	createRoute()
	err = http.ListenAndServe(bind+":"+port, nil) //设置监听的端口
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
