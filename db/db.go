package db

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

type DB struct {
	conn redis.Conn
}

func (db *DB) Set(key, val string) error {
	_, err := db.conn.Do("SET", key, val)
	return err
}

func (db *DB) Get(key string) (interface{}, error) {
	return db.conn.Do("GET", key)
}

func InitDB() *DB {
	pool := newPool()

	return &DB{
		conn: pool.Get(),
	}
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", os.Getenv("REDIS_PORT"))
			if err != nil {
				log.Panic(err)
			}

			return c, nil
		},
	}
}
