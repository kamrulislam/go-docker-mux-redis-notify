package main

import (
	"github.com/go-redis/redis"
)

//RedisConnect : create a connection to redis
func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

//AddRecord : add a record to Redis
func AddRecord(key string, data []byte) error {
	client := RedisConnect()
	defer client.Close()

	err := client.Set(key, data, 0).Err()

	return err

}

//FindRecord : find a record to Redis
func FindRecord(key string) (string, error) {
	client := RedisConnect()
	defer client.Close()

	data, err := client.Get(key).Result()

	return data, err

}
