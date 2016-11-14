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
redis> LRANGE 0 1
1) "one"
2) "two"

# retrieve all the items of a list
redis> LRANGE 0 -1
1) "one"
2) "two"
3) "three
```

## LPUSH key value [value ...]

- O(1)

```
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
redis> RPOP listkey
"three"
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
```

## RPOPLPUSH source destination

- O(1)

```
redis> RPUSH listkey three
(integer) 3
redis> RPOPLPUSH listkey destlistkey
"three"
redis> LRANGE listkey 0 -1
1) "one"
2) "two"
redis> LRANGE destlistkey 0 -1
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
redis> RPUSH listkey2 two
(integer) 1
redis> BLPOP listkey listkey2 5 // non-blocking
1) "listkey"
2) "one"
redis> BLPOP listkey listkey2 5 // non-blocking
1) "listkey2"
2) "two"
redis> BLPOP listkey listkey2 5 // blocking
(nil)
(5.04s)
```