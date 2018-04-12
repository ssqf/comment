package web

import (
	"document/model"
	"encoding/json"
	"net/http"
	"strings"
)

var userRex = "/user(/[0-9a-zA-z_-]+)?"

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		data, err := getBody(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(respError(err.Error()))
			return
		}

		u := model.User{}
		err = json.Unmarshal(data, &u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(respError(err.Error()))
			return
		}

		err = addUser(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(respError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == http.MethodDelete {
		id := strings.TrimPrefix(r.RequestURI, "/user/")
		err := delUser(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(respError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.RequestURI, "/user/")
		user, err := getUserByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(respError(err.Error()))
			return
		}

		js, err := json.MarshalIndent(user, "", "    ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(respError(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}

}

func addUser(user model.User) error {
	o := model.GetOperater()
	return o.AddUser(user)
}

func delUser(id string) error {
	o := model.GetOperater()
	return o.DelUserByID(id)
}

func getUserByID(id string) (model.User, error) {
	o := model.GetOperater()
	return o.GetUserByID(id)
}
