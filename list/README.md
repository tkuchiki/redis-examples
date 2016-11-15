# LIST Commands

## RPUSH key value [value ...]

- O(1)

```
redis> RPUSH listkey "one"
(integer) 1
redis> RPUSH listkey "two"
(integer) 2
redis> RPUSH listkey "three"
(integer) 3
```

## RPUSHX key value

- O(1)

## LRANGE key start stop

- O(S+N)

```
redis> RPUSH listkey "one"
(integer) 1
redis> RPUSH listkey "two"
(integer) 2
redis> LRANGE 0 1
1) "one"
2) "two"
redis> RPUSH listkey "three"
(integer) 3
# retrieve all the items of a list
redis> LRANGE 0 -1
1) "one"
2) "two"
3) "three
```

## LPUSH key value [value ...]

- O(1)

```
redis> RPUSH listkey "one"
(integer) 1
redis> RPUSH listkey "two"
(integer) 2
redis> RPUSH listkey "three"
(integer) 3
redis> LRANGE 0 -1
1) "one"
2) "two"
3) "three
redis> LPUSH listkey "zero"
(integer) 4
redis> LRANGE 0 -1
1) "zero"
2) "one"
3) "two"
4) "three
```

## LPUSHX key value

- O(1)

## LPOP key

- O(1)

```
redis> RPUSH listkey "zero"
(integer) 1
redis> RPUSH listkey "one"
(integer) 2
redis> RPUSH listkey "two"
(integer) 3
redis> RPUSH listkey "three"
(integer) 4
redis> LRANGE 0 -1
1) "zero"
2) "one"
3) "two"
4) "three
redis> LPOP listkey
"zero"
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
3) "three"
```

## RPOP key

- O(1)

```
redis> RPUSH listkey "one"
(integer) 1
redis> RPUSH listkey "two"
(integer) 2
redis> RPUSH listkey "three"
(integer) 3
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
3) "three"
redis> RPOP listkey
"three"
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
```

## RPOPLPUSH source destination

- O(1)

```
redis> RPUSH src "one"
(integer) 1
redis> RPUSH src "two"
(integer) 2
redis> RPUSH src three
(integer) 3
redis> LRANGE src 0 -1
1) "one"
2) "two"
3) "three"
redis> RPOPLPUSH src dest
"three"
redis> LRANGE src 0 -1
1) "one"
2) "two"
redis> LRANGE dest 0 -1
1) "three"
```

## LINDEX key index

- O(N)

```
redis> RPUSH listkey one
(integer) 1
redis> RPUSH listkey two
(integer) 2
redis> RPUSH listkey three
(integer) 3
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
3) "three"
redis> LINDEX listkey 0
"one"
redis> LINDEX listkey 1
"two"
redis> LINDEX listkey 2
"three"
redis> LINDEX listkey -1
"three"
```

## LINSERT key BEFORE|AFTER pivot value

- O(N)

```
redis> RPUSH listkey one
(integer) 1
redis> RPUSH listkey two
(integer) 2
redis> RPUSH listkey three
(integer) 3
redis> LINSERT listkey BEFORE "three" "two and a half"
(integer) 4
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
3) "two and a half"
4) "three"
127.0.0.1:6379> LINSERT listkey AFTER "one" "one and a half"
(integer) 5
127.0.0.1:6379> LRANGE listkey 0 -1
1) "one"
2) "one and a half"
3) "two"
4) "two and a half"
5) "three"
```

## LLEN key

- O(1)

```
redis> RPUSH listkey one
(integer) 1
redis> RPUSH listkey two
(integer) 2
redis> RPUSH listkey three
(integer) 3
redis> LLEN listkey
(integer) 3
```

## LREM key count value

- O(N)

```
redis> RPUSH listkey one
(integer) 1
redis> RPUSH listkey two
(integer) 2
redis> RPUSH listkey one
(integer) 3
redis> RPUSH listkey three
(integer) 4
redis> RPUSH listkey one
(integer) 5
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
3) "one"
4) "three"
5) "one"
redis> LREM listkey 2 one
(integer) 2
redis> LRANGE listkey 0 -1
1) "two"
2) "three"
3) "one"
redis> RPUSH listkey two
(integer) 4
redis> RPUSH listkey two
(integer) 5
redis> LRANGE listkey 0 -1
1) "two"
2) "three"
3) "one"
4) "two"
5) "two"
redis> LREM listkey 0 two
(integer) 3
redis> LRANGE listkey 0 -1
1) "three"
2) "one"
```

## LSET key index value

- O(N)
    - first or last = O(1)

```
redis> RPUSH listkey one
(integer) 1
redis> RPUSH listkey two
(integer) 2
redis> RPUSH listkey three
(integer) 3
redis> LSET listkey 0 1
OK
redis> LRANGE listkey 0 -1
1) "1"
2) "two"
3) "three"
redis> LSET listkey 3 four
(error) ERR index out of range
```

## LTRIM key start stop

- O(N)

```
redis> RPUSH listkey one
(integer) 1
redis> RPUSH listkey two
(integer) 2
redis> RPUSH listkey three
(integer) 3
redis> RPUSH listkey four
(integer) 4
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
3) "three"
4) "four"
redis> LTRIM listkey 2 -1
OK
redis> LRANGE listkey 0 -1
1) "three"
2) "four"
```

## BLPOP key [key ...] timeout

- O(1)

```
redis> RPUSH listkey one
(integer) 1
redis> RPUSH listkey two
(integer) 2
redis> RPUSH listkey2 three
(integer) 1
redis> BLPOP listkey listkey2 0 // non-blocking
1) "listkey"
2) "one"
redis> BLPOP listkey listkey2 0 // non-blocking
1) "listkey"
2) "two"
redis> BLPOP listkey listkey2 0 // non-blocking
1) "listkey2"
2) "three"
redis> BLPOP listkey listkey2 5 // blocking
(nil)
(5.02s)
```

```
redis client1> BLPOP listkey 0 // blocking | redis client2> RPUSH listkey one
                                           | (integer) 1
1) "listkey"
2) "one"
(3.30s)
```

## BRPOP key [key ...] timeout

- O(1)

```
redis> RPUSH listkey one
(integer) 1
redis> RPUSH listkey two
(integer) 2
redis> RPUSH listkey2 three
(integer) 1
redis> BRPOP listkey listkey2 5 // non-blocking
1) "listkey"
2) "two"
redis> BRPOP listkey listkey2 5 // non-blocking
1) "listkey"
2) "one"
redis> BRPOP listkey listkey2 5 // non-blocking
1) "listkey2"
2) "three"
redis> BRPOP listkey listkey2 5 // blocking
(nil)
(5.02s)
```

```
redis client1> BRPOP listkey 0 // blocking | redis client2> RPUSH listkey one
                                           | (integer) 1
1) "listkey"
2) "one"
(3.30s)
```

## BRPOPLPUSH source destination timeout

- O(1)

```
redis> RPUSH src one
(integer) 1
redis> RPUSH src two
(integer) 2
redis> RPUSH dest three
(integer) 1
redis> BRPOPLPUSH src dest 5 // non-blocking
"two"
redis> BRPOPLPUSH src dest 5 // non-blocking
"one"
redis> LRANGE src 0 -1
(empty list or set)
redis> LRANGE dest 0 -1
1) "one"
2) "two"
3) "three"
redis> BRPOPLPUSH src dest 5 // blocking
(nil)
(5.05s)
```

```
redis client1> RPUSH src one
(integer) 1
redis client1> RPUSH dest two
(integer) 1
redis client1> BRPOPLPUSH src dest 0
"one"
redis client1> BRPOPLPUSH src dest 0 // blocking | redis client2> RPUSH src zero
                                                 | (integer) 1
"zero"
(15.42s)
redis client1> LRANGE dest 0 -1
1) "zero"
2) "one"
3) "two"
```
