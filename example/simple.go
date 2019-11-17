package main

import (
	"github.com/mkusaka/dgos"
	"net/http"
)

func main() {
	http.Handle("/", dgos.Handler(okHandler))
	http.ListenAndServe(":3000", nil)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("OK"))
}