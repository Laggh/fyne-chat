package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Handler(w http.ResponseWriter, r *http.Request) {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	client := redis.NewClient(opt)

	val := client.Get(ctx, "foo").Val()
	fmt.Fprintf(w, "<h1>Hello!</h1><p>%s</p>", val)
}
