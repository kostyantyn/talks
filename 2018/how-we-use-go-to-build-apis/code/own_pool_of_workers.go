package main

import (
	"github.com/garyburd/redigo/redis"
)

// START OMIT
func main() {
	pool := 10
	ch := make(chan string, pool+10) // how many messages keep in a memory before taking an action // HL
	for i := 0; i < pool; i++ {
		conn, err := redis.Dial("tcp", ":6379") // connection per worker // HL
		_ = err // handle error
		go func() {
			conn.Do("HSET", "hash", <-ch, 1)
		}()
	}
}
// END OMIT
