package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/barokurniawan/gocrud/app"
	"github.com/barokurniawan/gocrud/service"
	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()
	routeServiceProvider := new(service.RouteServiceProvider)
	ServiceInitiator := new(service.InitService)
	srv := ServiceInitiator.Init()
	routeServiceProvider.SetRouter(router)

	guestbook := new(app.Guestbook)
	guestbook.SetRouteService(routeServiceProvider)
	guestbook.SetModel(srv.GB)
	guestbook.Route()

	router.PathPrefix("/js/").
		Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./assets/"))))

	routeServiceProvider.InitRoute()
	server := &http.Server{
		Addr:         "127.0.0.1:3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)

	defer cancel()
	defer srv.DB.Connection.Close()
	server.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
