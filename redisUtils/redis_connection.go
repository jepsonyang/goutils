package redisUtils

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

const kCheckLifeTime = 3 * time.Minute //健康检测间隔时间

type AddressFunc func() (string, error)

var redisOnce sync.Once
var redisPool *redis.Pool

type Connection struct {
	address  string //eg: "192.168.1.100:6379"
	password string
	index    int
	maxOpen  int
	maxIdle  int

	addrFunc AddressFunc
}

func (param *Connection) New(address string, password string, index int, maxOpen int, maxIdle int) {
	param.address = address
	param.password = password
	param.index = index
	param.maxOpen = maxOpen
	param.maxIdle = maxIdle
}

/*
* 设置动态获取IP地址的函数，如果设置了此函数，New()指定的地址将无效;
**/
func (param *Connection) SetAddrFunc(addrFunc AddressFunc) {
	param.addrFunc = addrFunc
}

func (param *Connection) GetConn() redis.Conn {
	redisOnce.Do(func() {
		redisPool = &redis.Pool{}
		redisPool.MaxActive = param.maxOpen
		redisPool.MaxIdle = param.maxIdle
		redisPool.TestOnBorrow = func(c redis.Conn, t time.Time) error { //健康检测(保活)
			if time.Since(t) < kCheckLifeTime {
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

			var addr string
			if param.addrFunc != nil {
				addr, err = param.addrFunc()
				if err != nil {
					return nil, err
				}
			} else {
				addr = param.address
			}

			var con redis.Conn
			con, err = redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if len(param.password) != 0 {
				if _, err := con.Do("AUTH", param.password); err != nil {
					con.Close()
					return nil, err
				}
			}
			if _, err = con.Do("SELECT", param.index); err != nil {
				con.Close()
				return nil, err
			}

			return con, nil
		}
	})
	return redisPool.Get()
}
