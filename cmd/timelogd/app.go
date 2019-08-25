package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irth/gotimelog"
)

type App struct {
	File gotimelog.TimelogFile
}

func (a *App) Routes(m *mux.Router) {
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
}
