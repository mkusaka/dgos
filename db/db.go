package db

import (
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)

func NewClient(options *redis.Options) *redis.Client {
	client := redis.NewClient(options)
	return client
}

type Redis struct {
	options redis.Options
	client  *redis.Client
}

func (r Redis) Start(addr string, password string, db int) Redis {
	options := redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
	r.options = options
	r.client = NewClient(&options)
	return r
}

func (r Redis) Inc(key string) (int64, error) {
	return r.client.Incr(key).Result()
}

func (r Redis) Count(key string) (int, error) {
	result, err := r.client.Get(key).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		log.Fatal("cannot get count")
	}
	count, err := strconv.Atoi(result)
	if err != nil {
		log.Fatal("cannot convert count string into integer")
	}
	return count, nil
}

func (r Redis) Set(key string, count int, expire_in time.Duration) {
	r.client.Set(key, count, expire_in)
}

func (r Redis) Exists(key string) (bool, error) {
	is, err := r.client.Exists(key).Result()
	if err == redis.Nil {
		return true, nil
	} else if err != nil {
		return false, err
	}
	return is > 0, nil
}
