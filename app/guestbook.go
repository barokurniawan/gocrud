package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"text/template"
	"time"

	"github.com/barokurniawan/gocrud/model"
	"github.com/barokurniawan/gocrud/service"
)

type Guestbook struct {
	rsp   *service.RouteServiceProvider
	model *model.Guestbook
}

func (gb *Guestbook) SetModel(model *model.Guestbook) {
	gb.model = model
}

func (gb *Guestbook) SetRouteService(rsp *service.RouteServiceProvider) {
	gb.rsp = rsp
}

func (gb Guestbook) Route() {
	gb.rsp.RegisterRoute("/api/list-message", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		list, err := gb.model.AdvanceShowList()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		byteJson, err := json.Marshal(list)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(byteJson)
	}))

	gb.rsp.RegisterRoute("/api/create-message", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := req.ParseForm()
		if err != nil {
			fmt.Fprint(w, false)
			return
		}

		Name := req.FormValue("name")
		Message := req.FormValue("message")
		CreatedAT := time.Now()

		_, err = gb.model.CreateNew(Name, Message, CreatedAT)
		if err != nil {
			fmt.Fprint(w, false)
			return
		}

		fmt.Fprint(w, true)
	}))

	gb.rsp.RegisterRoute("/api/delete-message", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var query = req.URL.Query()
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, err.Error())
			return
		}

		_, err = gb.model.Delete(int64(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, true)
	}))

	gb.rsp.RegisterRoute("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var filepath = path.Join("views", "index.html")
		tpl, err := template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
}
