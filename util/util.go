package util

import (
	"gopkg.in/redis.v5"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Del(client *redis.Client, key ...string) error {
	ic := client.Del(key...)

	return ic.Err()
}

func RPush(client *redis.Client, key string, val ...interface{}) error {
	ic := client.RPush(key, val...)

	return ic.Err()
}

func LPush(client *redis.Client, key string, val ...interface{}) error {
	ic := client.LPush(key, val...)

	return ic.Err()
}

func LRange(client *redis.Client, key string, start, stop int64) ([]string, error) {
	return client.LRange(key, start, stop).Result()
}

func LRangeAll(client *redis.Client, key string) ([]string, error) {
	return client.LRange(key, int64(0), int64(-1)).Result()
}
