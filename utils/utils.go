package utils

import (
	"crypto/md5"
	"fmt"
	"net/http"
)

// Hash is function which generate uniq key logic from given string
func Hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Key returns uniq key string from given request
func Key(r *http.Request) string {
	// key: urlstring/method/remote ip address
	keys := []string{r.URL.String(), r.Method, r.RemoteAddr}
	return Hash(strMap(keys))
}

func strMap(ss []string) string {
	r := ""
	for _, s := range ss {
		r = r + s
	}
	return r
}
