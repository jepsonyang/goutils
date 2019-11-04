package testRedisUtils

import (
	"github.com/gomodule/redigo/redis"
	"redisUtils"
)

const kRedisAddr = "9.134.7.64:6379"
const kRedisPassword = ""
const kRedisIndex = 15

const kRedisMaxOpenConns = 0	//一般设置为0,表示无限制
const kRedisMaxIdleConns = 20

var connect redisUtils.Connection

func init() {
	connect.New(kRedisAddr, kRedisPassword, kRedisIndex, kRedisMaxOpenConns, kRedisMaxIdleConns)
}

func GetRedisConnection() redis.Conn {
	return connect.GetConn()
}