package lesson_4

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Gorilla() {
	r := mux.NewRouter()

	r.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"time_from_gorilla_mux": "%s"}`, time.Now().Format(time.RFC3339))
	})

	http.ListenAndServe(":8080", r)
}
