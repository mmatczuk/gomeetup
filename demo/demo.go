package main

import (
	"net/http"
	"net/http/httputil"
	"log"
	"os"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Print(err)
			return
		}
		os.Stdout.Write(b)

		w.Write([]byte("Hello"))
	}))
}
