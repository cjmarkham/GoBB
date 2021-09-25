package main

import (
	"fmt"
	"net/http"
)

type App struct {
	server *http.Server
}

func ProvideApp(server *http.Server) *App {
	return &App{
		server: server,
	}
}

func (a *App) start() {
	if err := a.server.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %+v", err)
	}

	fmt.Printf("Server listening")
}

func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	app.start()
}
