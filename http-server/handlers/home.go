package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func home(commit, buildTime, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := struct {
			Commit    string `json:"commit"`
			BuildTime string `json:"buildTime"`
			Release   string `json:"release"`
		}{
			commit, buildTime, release,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
