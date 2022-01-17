package main

import (
	"log"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	names := map[string]string{"ruchit": "1", "aryan": "2"}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		id := r.URL.Query().Get("id")

		if _, ok := names[id]; ok {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("Found!"))
			if err != nil {
				log.Println("error in writing response")

				return
			}
		}
		//w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/test", testHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
