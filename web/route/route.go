package route

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

// func createRoute() {
// 	http.HandleFunc("/", index) //设置访问的路由
// 	http.HandleFunc("/acticle", getActicle)
// }

// NewRoute 创建一个route
func NewRoute() *Route {
	return &Route{make(map[string]routeHandler)}
}

var muxMutex sync.RWMutex

type routeHandler struct {
	method methodSlice
	f      http.HandlerFunc
}

// Route 控制器
type Route struct {
	mux map[string]routeHandler
}

type methodSlice []string

func (m methodSlice) isSupportMethod(method string) bool {
	for _, s := range m {
		if s == method {
			return true
		}
	}
	return false
}

// methodNotSupport 请求方法不支持返回错误
func methodNotSupport(w http.ResponseWriter, r *http.Request) {
	out := fmt.Sprintf("%s isn't support method:%s", r.RequestURI, r.Method)
	w.Write([]byte(out))
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// HttpServer 请求处理
func (c Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//rStr := fmt.Sprintf("Method:%s,ReqURI:%s,URL:{%v},RemoteAddr:%s", r.Method, r.RequestURI, r.URL, r.RemoteAddr)
	//fmt.Println(rStr)

	//w.Write([]byte(rStr))
	method := r.Method
	uri := r.RequestURI

	muxMutex.RLock()
	for patten, handler := range c.mux {
		re := regexp.MustCompile(patten)
		if re.MatchString(uri) {
			if handler.method.isSupportMethod(method) {
				handler.f(w, r)
			} else {
				methodNotSupport(w, r)
			}
			return
		}
	}
	muxMutex.RUnlock()
	http.NotFound(w, r)
}

// Add 添加路由和处理函数
func (c *Route) Add(patten string, method []string, f http.HandlerFunc) {
	h := routeHandler{method: method, f: f}
	patten = "^" + patten + "$" //正则完全匹配
	muxMutex.Lock()
	c.mux[patten] = h
	muxMutex.Unlock()
}
