package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
}

func NewRedis(network, addr, password string, dbNum, idle, active int64, dialTimeout, MaxConnLifetime time.Duration) (*Redis, func()) {
	db := &Redis{}
	db.pool = &redis.Pool{
		MaxIdle:         int(idle),
		MaxActive:       int(active),
		IdleTimeout:     dialTimeout,
		MaxConnLifetime: MaxConnLifetime,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", dbNum); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	cf := func() { db.pool.Close() }
	return db, cf
}
func (r *Redis) Set(key string, data string, seconds int64) (err error) {
	conn := r.pool.Get()
	defer conn.Close()
	if _, err = conn.Do("SET", key, data); err != nil {
		return
	}
	if _, err = conn.Do("EXPIRE", key, seconds); err != nil {
		return
	}
	return
}
func (r *Redis) Get(key string) (data string, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	if data, err = redis.String(conn.Do("GET", key)); err == redis.ErrNil {
		return "", nil
	}
	return
}
