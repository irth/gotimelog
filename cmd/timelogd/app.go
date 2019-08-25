package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/irth/gotimelog"
)

type App struct {
	File *gotimelog.TimelogFile
}

func (a *App) Routes(m *mux.Router) {
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	m.HandleFunc("/range", a.GetByRange)
}

func (a *App) GetByRange(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	start, end := q.Get("start"), q.Get("end")

	startTime, err := time.Parse(gotimelog.EntryDateFormat, start)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "unable to parse range start time: %v", err)
		return
	}
	endTime, err := time.Parse(gotimelog.EntryDateFormat, end)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "unable to parse range end time: %v", err)
		return
	}

	a.File.RLock()
	defer a.File.RUnlock()
	entries := a.File.Lines.EntriesByRange(startTime, endTime)
	json.NewEncoder(w).Encode(entries)
}
