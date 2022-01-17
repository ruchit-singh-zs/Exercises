package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func testHandler1(w http.ResponseWriter, r *http.Request) {
	names := map[string]string{"ruchit": "1", "aryan": "2", "anushi": "3"}

	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")

		if _, ok := names[id]; ok {
			w.WriteHeader(http.StatusOK)

			_, err := w.Write([]byte("Found!"))
			if err != nil {
				log.Println("error in writing response")

				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", testHandler1).Methods(http.MethodGet)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
