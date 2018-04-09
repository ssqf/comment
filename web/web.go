package web

import (
	"document/config"
	"document/web/route"
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
	//createRoute
	c := route.NewRoute()
	c.Add(actRegStr, []string{http.MethodGet, http.MethodPost, http.MethodDelete}, acticle)
	c.Add("/", []string{http.MethodGet}, index)
	c.Add("/p", []string{http.MethodPost}, post)
	c.Add(userRex, []string{http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodPatch}, user)

	err = http.ListenAndServe(bind+":"+port, c) //设置监听的端口
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
