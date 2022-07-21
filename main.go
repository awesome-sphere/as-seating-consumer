package main

import (
	"log"

	"github.com/awesome-sphere/as-seating-consumer/kafka"
	"github.com/awesome-sphere/as-seating-consumer/redis"
)

func main() {
	redis.InitializeRedisConn()
	log.Println("Starting kafka...")
	kafka.InitKafkaTopic()
}
