package config

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func RedisConn() redis.Conn {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
