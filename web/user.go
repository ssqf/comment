package web

import (
	"document/model"
	"encoding/json"
	"net/http"
)

var userRex = "/user(/[0-9a-zA-z_-]+)?"

func user(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/user" && r.Method == http.MethodPost {
		data, err := getBody(r)
		if err != nil {
			w.Write([]byte("get request body error"))
			return
		}

		u := model.User{}
		err = json.Unmarshal(data, &u)
		if err != nil {
			w.Write([]byte("Unmarshal request body error"))
			return
		}

		err = addUser(u)
		if err != nil {
			w.Write([]byte("add user error"))
			return
		}
		w.Write([]byte("add user ok"))
	}
}

func addUser(user model.User) error {
	o := model.GetOperater()
	return o.AddUser(user)
}
