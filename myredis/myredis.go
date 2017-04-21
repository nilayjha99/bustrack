package myredis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// NewPool is a function which creates and returns a connection pool of redis
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   2,
		MaxActive: 5, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			fmt.Print(err)
			return c, err
		},
	}

}

// RedisPool is where we initialize pool
var RedisPool *redis.Pool

// InitPool is to make a new redis connection pool
func InitPool() bool {
	RedisPool = NewPool()
	if RedisPool == nil {
		return false
	}
	return true
}

// GetConnection returns a connection form newely created redis pool
func GetConnection() redis.Conn {
	return RedisPool.Get()
}

// ClosePool to close the redis pool
func ClosePool() bool {
	err := RedisPool.Close()
	if err != nil {
		return false
	}
	return true
}

//
// c := pool.Get()
// defer c.Close()
