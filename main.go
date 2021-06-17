package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	opt, err0 := redis.ParseURL(os.Getenv("redis"))
	if err0 != nil {
		panic(err0)
	}
	client = redis.NewClient(opt)
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		typeId := strings.TrimLeft(r.URL.Path, "/")
		hour := time.Now().Unix() / 3600
		key := fmt.Sprintf("%s:%d", typeId, hour)
		id := client.Incr(key).Val()
		go client.Expire(key, time.Hour)
		val := fmt.Sprintf("%s%d%07d", typeId, hour, id)
		w.Write([]byte(val))
	})
	http.ListenAndServe(":80", nil)
}
