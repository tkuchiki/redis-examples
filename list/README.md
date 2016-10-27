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

## RPOPLPUSH

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

## LINDEX

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

## LINSERT

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

## LLEN

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
