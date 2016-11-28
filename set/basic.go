package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/tkuchiki/redis-examples/util"
	//"gopkg.in/redis.v5"
)

func main() {
	client := util.NewClient()

	key1 := "setkey"

	// SADD one
	val1 := "one"
	fmt.Println("# SADD", key1, val1)
	_ = util.SAdd(client, key1, val1)

	// SADD two
	val2 := "two"
	fmt.Println("# SADD", key1, val2)
	_ = util.SAdd(client, key1, val2)

	// SADD two
	fmt.Println("# SADD", key1, val2)
	_ = util.SAdd(client, key1, val2)

	fmt.Println()

	// retrieve all the items of a set
	var set []string
	fmt.Println("# SMEMBERS", key1)
	set, _ = util.SMembers(client, key1)

	pp.Println(set)
	fmt.Println()

	// delete setkey
	fmt.Println("# DEL", key1)
	_ = util.Del(client, key1)
}
