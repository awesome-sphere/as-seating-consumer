package redis

import (
	"fmt"
	"log"
	"strconv"

	"github.com/awesome-sphere/as-seating-consumer/utils"
	"github.com/go-redis/redis"
)

var CLIENT *redis.Client

/*
REDIS KEY FORMAT: (theater_id)-(timeslot_id)-(seat_id)
REDIS VALUE FORMAT: seat_status
*/

func InitializeRedisConn() {
	redisHost := utils.GetenvOr("REDIS_HOST", "localhost")
	redisPort := utils.GetenvOr("REDIS_PORT", "6379")
	redisPassword := utils.GetenvOr("REDIS_PASSWORD", "")
	redisDB, err := strconv.Atoi(utils.GetenvOr("REDIS_DB", "0"))

	if err != nil {
		log.Fatalf("Failed to convert REDIS_DB to int: %v", err)
		return
	}

	CLIENT = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       redisDB,
	})
}
