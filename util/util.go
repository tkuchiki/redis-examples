package util

import (
	"gopkg.in/redis.v5"
	"time"
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

func LPop(client *redis.Client, key string) (string, error) {
	scmd := client.LPop(key)
	val, err := scmd.Result()
	return val, err
}

func RPop(client *redis.Client, key string) (string, error) {
	scmd := client.RPop(key)
	val, err := scmd.Result()
	return val, err
}

func RPopLPush(client *redis.Client, src, dest string) (string, error) {
	scmd := client.RPopLPush(src, dest)
	val, err := scmd.Result()
	return val, err
}

func LIndex(client *redis.Client, key string, index int64) (string, error) {
	scmd := client.LIndex(key, index)
	val, err := scmd.Result()
	return val, err
}

func LInsert(client *redis.Client, key, op string, pivot, value interface{}) error {
	icmd := client.LInsert(key, op, pivot, value)
	return icmd.Err()
}

func LLen(client *redis.Client, key string) int64 {
	icmd := client.LLen(key)
	return icmd.Val()
}

func LRem(client *redis.Client, key string, count int64, value interface{}) int64 {
	icmd := client.LRem(key, count, value)
	return icmd.Val()
}

func LSet(client *redis.Client, key string, index int64, value interface{}) error {
	statusCmd := client.LSet(key, index, value)
	return statusCmd.Err()
}

func LTrim(client *redis.Client, key string, start, stop int64) error {
	statusCmd := client.LTrim(key, start, stop)
	return statusCmd.Err()
}

func BLPop(client *redis.Client, timeout time.Duration, keys ...string) ([]string, error) {
	return client.BLPop(timeout, keys...).Result()
}
