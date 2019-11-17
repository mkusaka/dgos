package main

import (
	"github.com/mkusaka/dgos"
	"net/http"
)

func main() {
	http.Handle("/", dgos.Handler(okHandler))
	http.ListendAndServe(":3000", nil)
}

func okHanlder(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("OK"))
}