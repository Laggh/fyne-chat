package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var opt, _ = redis.ParseURL(os.Getenv("REDIS_URL"))
var client *redis.Client

func Handler(w http.ResponseWriter, r *http.Request) {
	client = redis.NewClient(opt)
	val := client.Get(ctx, "foo").Val()
	fmt.Fprintf(w, "<h1>Hello</h1><p>%s</p>", val)
}
