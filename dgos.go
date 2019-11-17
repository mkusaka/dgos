package dgos

import (
	"net/http"

	"dgos/db"
	"dgos/utils"

	"log"
	"time"
)

var currentRedis = db.Redis{}
var client = currentRedis.Start("localhost:6379", "", 0)

func Handler(next http.Handler, bantime time.Duration, findtime time.Duration, maxretry int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := utils.Key(r)
		ok, err := checkCount(key, bantime, findtime, maxretry)
		if err != nil {
			log.Fatalf("check failed: %v", err)
		}
		if !ok {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
		}
		next.ServeHTTP(w, r)
	})
}

func checkCount(key string, bantime time.Duration, findtime time.Duration, maxretry int) (bool, error) {
	if isBanned(key) {
		return false, nil
	}

	count, err := client.Count(key)

	if err != nil {
		log.Fatal("something wrong with count")
	}

	if count > maxretry {
		ban(key, bantime)
		return false, nil
	}

	update(key)
	return true, nil
}

func isBanned(key string) bool {
	banned, err := client.Exists(key)
	if err != nil {
		log.Fatal(err)
	}
	return banned
}

func ban(key string, bantime time.Duration) {
	client.Set(key, 1, bantime)
}

func update(key string) {
	_, err := client.Inc(key)

	if err != nil {
		log.Fatalf("something wrong with inc count key: %s", key)
	}
}
