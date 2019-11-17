# dgos
add request limitation middleware to protect ddos attack.

# useage

```bash
go get -u github.com/mkusaka/dgos
```

```go
package main

import (
	"net/http"
	"time"

	"github.com/mkusaka/dgos"
)

func main() {
	// After 5 requests in 2 minutes, block all requests from that IP for 1 minute
	http.Handle("/", dgos.Handler(http.HandlerFunc(okHandler), time.Duration(60)*time.Second, time.Duration(120)*time.Second, 5))
	http.ListenAndServe(":3000", nil)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
```


# memo
- request handle middleware
- request count data in redis (key: hashed url/method/request_ip)
- update request count

# TODO
- [ ] use docker
- [ ] refactoring
- [ ] worker for log
- [ ] add [Retry-After header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After) for request if it reached limit
