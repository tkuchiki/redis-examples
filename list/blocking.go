package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/tkuchiki/redis-examples/util"
	//"gopkg.in/redis.v5"
	"sync"
	"time"
)

func main() {
	client := util.NewClient()

	key1 := "listkey"

	// RPUSH one
	val1 := "one"
	fmt.Println("# RPUSH", key1, val1)
	_ = util.RPush(client, key1, val1)

	fmt.Println()

	var list []string
	fmt.Println("# LRANGE", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)
	fmt.Println()

	fmt.Println("# BLPOP", key1, 0)
	list, _ = util.BLPop(client, time.Duration(0), key1)

	pp.Println(list)
	fmt.Println()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("# BLPOP listkey 0 (blocking...)")
			list, _ = util.BLPop(client, time.Duration(0), key1)
			fmt.Println("// RPUSHed")
			pp.Println(list)
			fmt.Println()
			break
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(5 * time.Second)
			v := "blocking"
			fmt.Println("# RPUSH", key1, v)
			_ = util.RPush(client, key1, v)
			break
		}
	}()

	wg.Wait()

	// RPUSH one
	fmt.Println("# RPUSH", key1, val1)
	_ = util.RPush(client, key1, val1)

	fmt.Println()

	key2 := "listkey2"

	// RPUSH two
	val2 := "two"
	fmt.Println("# RPUSH", key2, val2)
	_ = util.RPush(client, key2, val2)

	fmt.Println()

	// BLPOP listkey listkey2 0
	fmt.Println("# BLPOP", key1, key2, 0)
	list, _ = util.BLPop(client, time.Duration(0), key1, key2)

	pp.Println(list)
	fmt.Println()

	// BLPOP listkey listkey2 0
	fmt.Println("# BLPOP", key1, key2, 0)
	list, _ = util.BLPop(client, time.Duration(0), key1, key2)

	pp.Println(list)
	fmt.Println()

	// BLPOP listkey listkey2 0
	fmt.Println("# BLPOP", key1, key2, 5, "(timeout)")
	list, _ = util.BLPop(client, time.Duration(5), key1, key2)

	pp.Println(list)
	fmt.Println()

	// delete listkey
	fmt.Println("# DEL", key1)
	_ = util.Del(client, key1)
}
