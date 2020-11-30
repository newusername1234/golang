package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	// http.HandleFunc("/", handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			if k == "cycles" {
				cycles, err := strconv.Atoi(v[0])
				if err != nil {
					log.Print(err)
					return
				}
				lissajous(w, cycles)
			}
		}
		// send res with 5 cycles if no param
		lissajous(w, 5)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	lissajous(w)
// }
