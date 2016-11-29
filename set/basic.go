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

	// SADD setkey one
	val1 := "one"
	fmt.Println("# SADD", key1, val1)
	_ = util.SAdd(client, key1, val1)

	// SADD setkey two
	val2 := "two"
	fmt.Println("# SADD", key1, val2)
	_ = util.SAdd(client, key1, val2)

	// SADD setkey two
	fmt.Println("# SADD", key1, val2)
	_ = util.SAdd(client, key1, val2)

	fmt.Println()

	// retrieve all the items of a set
	var set []string
	fmt.Println("# SMEMBERS", key1)
	set, _ = util.SMembers(client, key1)

	pp.Println(set)
	fmt.Println()

	// SCARD setkey
	fmt.Println("# SCARD", key1)
	pp.Println(util.SCard(client, key1))

	// SADD setkey three
	val3 := "three"
	fmt.Println("# SADD", key1, val3)
	_ = util.SAdd(client, key1, val3)

	// SADD setkey three
	key2 := "setkey2"
	fmt.Println("# SADD", key2, val3)
	_ = util.SAdd(client, key2, val3)

	// SADD setkey four
	val4 := "four"
	fmt.Println("# SADD", key2, val4)
	_ = util.SAdd(client, key2, val4)

	// SADD setkey five
	val5 := "five"
	fmt.Println("# SADD", key2, val5)
	_ = util.SAdd(client, key2, val5)

	// SADD setkey3 two
	key3 := "setkey3"
	fmt.Println("# SADD", key3, val2)
	_ = util.SAdd(client, key3, val2)

	fmt.Println("# SMEMBERS", key1)
	set, _ = util.SMembers(client, key1)

	pp.Println(set)
	fmt.Println()

	fmt.Println("# SMEMBERS", key2)
	set, _ = util.SMembers(client, key2)

	pp.Println(set)
	fmt.Println()

	fmt.Println("# SMEMBERS", key3)
	set, _ = util.SMembers(client, key3)

	pp.Println(set)
	fmt.Println()

	var diff []string
	fmt.Println("# SDIFF", key1, key2)
	diff, _ = util.SDiff(client, key1, key2)

	pp.Println(diff)
	fmt.Println()

	fmt.Println("# SDIFF", key2, key1)
	diff, _ = util.SDiff(client, key2, key1)

	pp.Println(diff)
	fmt.Println()

	fmt.Println("# SDIFF", key1, key2, key3)
	diff, _ = util.SDiff(client, key1, key2, key3)

	pp.Println(diff)
	fmt.Println()

	// delete setkey
	fmt.Println("# DEL", key1)
	_ = util.Del(client, key1)

	// delete setkey2
	fmt.Println("# DEL", key2)
	_ = util.Del(client, key2)

	// delete setkey3
	fmt.Println("# DEL", key3)
	_ = util.Del(client, key3)
}
