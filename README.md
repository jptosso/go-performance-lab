# Conclusions

- If a string is ASCII, for string to lower and upper it's faster to create a helper
- If possible, always preallocate slices
- Arrays are created faster than slices, use Arrays
- Avoid CGO
- Avoid sync.Mutex if possible
- Sync.Pool can make your program slower, use it wisely
- strings to byte and backwards consumes a lot of resources
- Reflection is slower than type assertion
- Type assertion without switch is faster. Type assertion is 50%~ slower than using literal
- It is faster to create a big array and shorten it into a slice than using append, but it elevates memory consumption (`make([]type, length, capacity)` is better)

# How to run

```sh
$ git clone https://github.com/jptosso/go-performance-lab
$ cd go-performance-lab
$ go test -bench=. .
```

You may also add `-memprofile filename` to write pprof compatible profiles.
