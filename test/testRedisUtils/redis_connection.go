package testRedisUtils

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

const kRedisAddr = "9.134.7.64:6379"
const kRedisPassword = ""
const kRedisIndex = 0

const kRedisMaxOpenConns = 0	//一般设置为0,表示无限制
const kRedisMaxIdleConns = 20

var redisOnce sync.Once
var redisPool *redis.Pool

func GetRedisConnection() redis.Conn {
	redisOnce.Do(func() {
		redisPool = &redis.Pool{}
		redisPool.MaxActive = kRedisMaxOpenConns
		redisPool.MaxIdle = kRedisMaxIdleConns
		redisPool.TestOnBorrow = func(c redis.Conn, t time.Time) error { //健康检测(保活)
			if time.Since(t) < time.Minute * 3 {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				return nil
			}
			return nil
		}

		redisPool.Dial = func() (redis.Conn, error) {
			var err error

			var con redis.Conn
			con, err = redis.Dial("tcp", kRedisAddr)
			if err != nil {
				return nil, err
			}
			if len(kRedisPassword) > 0 {
				if _, err := con.Do("AUTH", kRedisPassword); err != nil {
					_ = con.Close()
					return nil, err
				}
			}
			if _, err = con.Do("SELECT", kRedisIndex); err != nil {
				_ =con.Close()
				return nil, err
			}

			return con, nil
		}
	})
	return redisPool.Get()
}