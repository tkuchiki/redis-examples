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
	fmt.Println("# LRANGE", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)

	// LPOP
	fmt.Println("# LPOP", key1)
	lpoped, _ := util.LPop(client, key1)
	pp.Println(lpoped)
	fmt.Println()

	// LPOPed LIST
	fmt.Println("# LRANGE", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)

	fmt.Println()

	// RPOP
	fmt.Println("# RPOP", key1)
	rpoped, _ := util.RPop(client, key1)
	pp.Println(rpoped)
	fmt.Println()

	// RPOPed LIST
	fmt.Println("# LRANGE", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)
	fmt.Println()

	// RPUSH three
	fmt.Println("# RPUSH", key1, val3)
	_ = util.RPush(client, key1, val3)

	fmt.Println()

	// LINDEX listkey 0
	fmt.Println("# LINDEX", key1, 0)
	l, _ := util.LIndex(client, key1, int64(0))

	pp.Println(l)

	// LINDEX listkey 1
	fmt.Println("# LINDEX", key1, 1)
	l, _ = util.LIndex(client, key1, int64(1))

	pp.Println(l)

	// LINDEX listkey 1
	fmt.Println("# LINDEX", key1, -1)
	l, _ = util.LIndex(client, key1, int64(-1))

	pp.Println(l)

	fmt.Println()

	key2 := "destlistkey"

	// RPOPLPUSH
	fmt.Println("# RPOPLPUSH", key1, key2)
	rpoped, _ = util.RPopLPush(client, key1, key2)
	pp.Println(rpoped)
	fmt.Println()

	// RPOPLPUSHed LIST
	fmt.Println("# LRANGE", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)

	fmt.Println()

	// RPOPLPUSHed LIST
	fmt.Println("# LRANGE", key2, 0, -1)
	list, _ = util.LRangeAll(client, key2)

	pp.Println(list)

	fmt.Println()

	// RPUSH three
	fmt.Println("# RPUSH", key1, val3)
	_ = util.RPush(client, key1, val3)

	fmt.Println()

	// LINSERT listkey BEFORE three "two and a half"
	fmt.Println("# LINSERT", key1, "BEFORE", "three", "two and a half")
	_ = util.LInsert(client, key1, "BEFORE", "three", "two and a half")

	fmt.Println()

	// LINSERTed LIST
	fmt.Println("# LRANGe", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)
	fmt.Println()

	// LINSERT listkey AFTER one "one and a half"
	fmt.Println("# LINSERT", key1, "AFTER", "one", "one and a half")
	_ = util.LInsert(client, key1, "AFTER", "one", "one and a half")

	fmt.Println()

	// LINSERTed LIST
	fmt.Println("# LRANGE", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)
	fmt.Println()

	// LLEN listkey
	fmt.Println("# LLEN", key1)
	llen := util.LLen(client, key1)

	pp.Println(llen)
	fmt.Println()

	// delete listkey
	fmt.Println("# DEL", key1)
	_ = util.Del(client, key1)

	// delete destlistkey
	fmt.Println("# DEL", key2)
	_ = util.Del(client, key2)
}
