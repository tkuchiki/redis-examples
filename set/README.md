# Set Commands

## SMEMBERS key

- O(N)

```
redis> SADD setkey one
(integer) 1
redis> SADD setkey two
(integer) 1
redis> SMEMBERS setkey
1) "two"
2) "one"
```

## SADD key member [member ...]

- O(1)
- O(N)
    - multiple arguments(member)

```
redis> SADD setkey one
(integer) 1
redis> SADD setkey two
(integer) 1
redis> SADD setkey two
(integer) 0
redis> SMEMBERS setkey
1) "two"
2) "one"
```

## SCARD key

- O(1)

```
redis> SADD setkey one
(integer) 1
redis> SADD setkey two
(integer) 1
redis> SCARD setkey
(integer) 2
```

## SDIFF key [key ...]

- O(N)
    - N = total number of elements in all given sets

```
redis> SADD setkey one
(integer) 1
redis> SADD setkey two
(integer) 1
redis> SADD setkey three
(integer) 1
redis> SADD setkey2 three
(integer) 1
redis> SADD setkey2 four
(integer) 1
redis> SADD setkey2 five
(integer) 1
redis> SDIFF setkey setkey2
1) "one"
2) "two"
redis> SDIFF setkey2 setkey
1) "five"
2) "four"
redis> SADD setkey3 two
(integer) 1
redis> SDIFF setkey setkey2 setkey3
1) "one"
```
