package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irth/gotimelog"
)

func main() {
	f := gotimelog.TimelogFile{Path: "/home/me/.local/share/gtimelog/timelog.txt"}
	err := f.Load()
	if err != nil {
		panic(err)
	}

	app := App{
		File: f,
	}

	m := mux.NewRouter()
	app.Routes(m)
	http.ListenAndServe(":2137", m)
}
