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
