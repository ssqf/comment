package web

import (
	"log"
	"net/http"
	"regexp"
)

var actRegStr = "/acticle(/[0-9a-zA-z_-]+)?"

func acticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	uri := r.RequestURI
	log.Println("Acticle URI:", uri)

	actRegx := regexp.MustCompile(actRegStr)
	if !actRegx.MatchString(uri) {
		log.Println("Acticle URI don't match regexp:", "/acticle(/[0-9a-zA-z_-])$")
	} else {
		uriList := actRegx.FindStringSubmatch(uri)
		lLen := len(uriList)
		if lLen == 1 {
			w.Write([]byte(uriList[0] + " Get ALL acticle"))
		} else if lLen == 2 {
			w.Write([]byte(uriList[0] + " Get ID of acticle is " + uriList[1]))
		} else {
			log.Println("Acticle URI:", uri, "regexp wrong match")
		}
	}

}
