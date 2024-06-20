package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"go-hugging-api/api"
	hugging2 "go-hugging-api/hugging"
	"log"
	"os"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	// Define a limit rate to 5 requests per second.
	rate, err := limiter.NewRateFromFormatted("5-M")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a store with the redis client.
	store, err := sredis.NewStore(rdb)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a new middleware with the limiter instance.
	limiterMiddleware := mgin.NewMiddleware(limiter.New(store, rate))

	hugging := hugging2.Hugging{
		Token: os.Getenv("HF_API_TOKEN"),
		Rdb:   rdb,
	}

	a := api.Api{
		Hugging: hugging,
	}

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(limiterMiddleware)
	router.POST("/text-classification", a.TextClassificationHandler)
	router.Run(":8080")
}
