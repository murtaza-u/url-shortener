package db

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type DB struct {
	conn redis.Conn
}

func (db *DB) Set(key, value string) error {
	_, err := db.conn.Do("SET", key, value)
	return err
}

func (db *DB) Get(key string) error {
	_, err := db.conn.Do("GET", key)
	return err
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
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				log.Panic(err)
			}

			return c, nil
		},
	}
}
