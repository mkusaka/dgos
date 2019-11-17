# dgos
add request limitation to protect ddos attack middleware.

# useage

```bash
go get -u github.com/mkusaka/dgos
```

```go
package main
import (
  "github.com/mkusaka/dgos"

)
```


# memo
- request handle middleware
- request count data in redis (key: hashed url/method/request_ip)
- update request count

# TODO
- [ ] use docker
- [ ] refactoring
- [ ] worker for log
