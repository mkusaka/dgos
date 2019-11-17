package main

import (
	"net/http"
	"time"

	"github.com/mkusaka/dgos"
)

func main() {
	http.Handle("/", dgos.Handler(http.HandlerFunc(okHandler), time.Duration(60)*time.Second, time.Duration(120)*time.Second, 5))
	http.ListenAndServe(":3000", nil)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
