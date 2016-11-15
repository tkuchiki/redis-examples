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
			fmt.Println("# client1> BLPOP", key1, 0, "(blocking...)")
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
			fmt.Println("# client2> RPUSH", key1, v)
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
	list, _ = util.BLPop(client, 5*time.Second, key1, key2)

	pp.Println(list)
	fmt.Println()

	// RPUSH one
	fmt.Println("# RPUSH", key1, val1)
	_ = util.RPush(client, key1, val1)

	fmt.Println()

	fmt.Println("# LRANGE", key1, 0, -1)
	list, _ = util.LRangeAll(client, key1)

	pp.Println(list)
	fmt.Println()

	fmt.Println("# BRPOP", key1, 0)
	list, _ = util.BRPop(client, time.Duration(0), key1)

	pp.Println(list)
	fmt.Println()

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("# client1> BRPOP", key1, 0, "(blocking...)")
			list, _ = util.BRPop(client, time.Duration(0), key1)
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
			fmt.Println("# client2> RPUSH", key1, v)
			_ = util.RPush(client, key1, v)
			break
		}
	}()

	wg.Wait()

	// RPUSH one
	fmt.Println("# RPUSH", key1, val1)
	_ = util.RPush(client, key1, val1)
	// RPUSH two
	fmt.Println("# RPUSH", key2, val2)
	_ = util.RPush(client, key2, val2)

	fmt.Println()

	// BLPOP listkey listkey2 0
	fmt.Println("# BRPOP", key1, key2, 0)
	list, _ = util.BRPop(client, time.Duration(0), key1, key2)

	pp.Println(list)
	fmt.Println()

	// BLPOP listkey listkey2 0
	fmt.Println("# BRPOP", key1, key2, 0)
	list, _ = util.BRPop(client, time.Duration(0), key1, key2)

	pp.Println(list)
	fmt.Println()

	// BLPOP listkey listkey2 0
	fmt.Println("# BRPOP", key1, key2, 5, "(timeout)")
	list, _ = util.BRPop(client, 5*time.Second, key1, key2)

	pp.Println(list)
	fmt.Println()

	srckey := "src"
	destkey := "dest"
	val3 := "three"

	// RPUSH src one
	fmt.Println("# RPUSH", srckey, val1)
	_ = util.RPush(client, srckey, val1)
	// RPUSH src two
	fmt.Println("# RPUSH", srckey, val2)
	_ = util.RPush(client, srckey, val2)
	// RPUSH dest three
	fmt.Println("# RPUSH", destkey, val3)
	_ = util.RPush(client, destkey, val3)

	fmt.Println()

	// BLPOP listkey listkey2 0
	fmt.Println("# BRPOPLPUSH", srckey, destkey, 0)
	rpoped, _ := util.BRPopLPush(client, srckey, destkey, time.Duration(0))

	pp.Println(rpoped)
	fmt.Println()

	// BLPOP listkey listkey2 0
	fmt.Println("# BRPOPLPUSH", srckey, destkey, 0)
	rpoped, _ = util.BRPopLPush(client, srckey, destkey, time.Duration(0))

	pp.Println(rpoped)
	fmt.Println()

	fmt.Println("# LRANGE", srckey, 0, -1)
	list, _ = util.LRangeAll(client, srckey)

	pp.Println(list)
	fmt.Println()

	fmt.Println("# LRANGE", destkey, 0, -1)
	list, _ = util.LRangeAll(client, destkey)

	pp.Println(list)
	fmt.Println()

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("# client1> BRPOPLPUSH", srckey, destkey, 0, "(blocking...)")
			rpoped, _ = util.BRPopLPush(client, srckey, destkey, time.Duration(0))
			fmt.Println("// RPUSHed")
			pp.Println(rpoped)
			fmt.Println()
			break
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(5 * time.Second)
			v := "blocking"
			fmt.Println("# client2> RPUSH", srckey, v)
			_ = util.RPush(client, srckey, v)
			break
		}
	}()

	wg.Wait()

	// delete listkey
	fmt.Println("# DEL", key1)
	_ = util.Del(client, key1)

	// delete listkey2
	fmt.Println("# DEL", key2)
	_ = util.Del(client, key2)
}
