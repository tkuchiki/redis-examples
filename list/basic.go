package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/tkuchiki/redis-examples/util"
	//"gopkg.in/redis.v5"
)

func main() {
	client := util.NewClient()

	key1 := "listkey"

	// RPUSH one
	val1 := "one"
	fmt.Println("# RPUSH", key1, val1)
	_ = util.RPush(client, key1, val1)

	// RPUSH two
	val2 := "two"
	fmt.Println("# RPUSH", key1, val2)
	_ = util.RPush(client, key1, val2)

	// RPUSH three
	val3 := "three"
	fmt.Println("# RPUSH", key1, val3)
	_ = util.RPush(client, key1, val3)

	fmt.Println()

	// retrieve all the items of a list
	var list []string
	fmt.Println("# LRange", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)

	// retrieve two three
	fmt.Println("# LRange", key1, 1, 2)
	list, _ = util.LRange(client, key1, int64(1), int64(2))

	pp.Println(list)

	// LPUSH zero
	val0 := "zero"
	fmt.Println("# LPUSH", key1, val0)
	_ = util.LPush(client, key1, val0)

	// LPUSHed LIST
	fmt.Println("# LRange", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)

	// delete listkey1
	fmt.Println("# DEL", key1)
	_ = util.Del(client, key1)
}
