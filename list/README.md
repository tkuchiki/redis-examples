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