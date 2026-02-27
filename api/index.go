package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var client *redis.Client
var once sync.Once

func GetRedis() (*redis.Client, error) {
	var err error
	once.Do(func() {
		redisURL := os.Getenv("REDIS_URL")
		if redisURL == "" {
			err = fmt.Errorf("REDIS_URL vazio")
			return
		}

		opt, parseErr := redis.ParseURL(redisURL)
		if parseErr != nil {
			err = fmt.Errorf("ParseURL falhou: %v", parseErr)
			return
		}

		client = redis.NewClient(opt)
	})
	return client, err
}

func Handler(w http.ResponseWriter, r *http.Request) {
	c, err := GetRedis()
	if err != nil {
		fmt.Fprintf(w, "Redis error: %v", err)
		return
	}

	val := c.Get(Ctx, "foo").Val()
	fmt.Fprintf(w, "<h1>Hello!</h1><p>%s</p>", val)
}
