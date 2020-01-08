package app

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"text/template"
	"time"

	"github.com/barokurniawan/gocrud/entity"

	"github.com/barokurniawan/gocrud/model"
	"github.com/barokurniawan/gocrud/service"
	"github.com/barokurniawan/gocrud/sys"
)

type Guestbook struct {
	rsp   *service.RouteServiceProvider
	model *model.Guestbook
}

func NewGuestbookApp(rsp *service.RouteServiceProvider, model *model.Guestbook) *Guestbook {
	var gb = &Guestbook{
		rsp:   rsp,
		model: model,
	}

	gb.Route()
	return gb
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
			res := sys.Response{
				Info:    false,
				Message: err.Error(),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		byteJson, err := json.Marshal(list)
		if err != nil {
			res := sys.Response{
				Info:    false,
				Message: err.Error(),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteJson)
	}))

	gb.rsp.RegisterRoute("/api/create-message", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			res := sys.Response{
				Info:    false,
				Message: err.Error(),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		Name := req.FormValue("name")
		Message := req.FormValue("message")
		CreatedAT := time.Now()

		_, err = gb.model.CreateNew(Name, Message, CreatedAT)
		if err != nil {
			res := sys.Response{
				Info:    false,
				Message: err.Error(),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		res := sys.Response{
			Info:    true,
			Message: "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res.Parse())
	}))

	gb.rsp.RegisterRoute("/api/delete-message", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var query = req.URL.Query()
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			res := sys.Response{
				Info:    false,
				Message: err.Error(),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		_, err = gb.model.Delete(int64(id))
		if err != nil {
			res := sys.Response{
				Info:    false,
				Message: err.Error(),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		res := sys.Response{
			Info:    true,
			Message: "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res.Parse())
	}))

	gb.rsp.RegisterRoute("/api/update-message", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			res := sys.Response{
				Info:    false,
				Message: "Invalid Method",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		i, err := strconv.Atoi(req.FormValue("id"))
		id := int64(i)
		if err != nil {
			res := sys.Response{
				Info:    false,
				Message: "Invalid ID",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		item := entity.Guestbook{
			ID:      id,
			Name:    req.FormValue("name"),
			Message: req.FormValue("message"),
		}

		info, err := gb.model.Update(id, item)
		if err != nil {
			res := sys.Response{
				Info:    false,
				Message: "Invalid ID",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res.Parse())
			return
		}

		res := sys.Response{
			Info:    info,
			Message: "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res.Parse())
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
