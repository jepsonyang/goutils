package redisutils

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

var RedisStringTypeErr = errors.New("type is not string")

/*
* 设置string值
* expire 过期时间,单位: 秒; 永不过期填负数即可(一般填-1)
**/
func RedisStringSet(conn redis.Conn, redisType RedisType, value string, expire int) error {
	if redisType.GetType() != KRedisString {
		return RedisStringTypeErr
	}
	return RedisStringSetKey(conn, redisType.GetKey(), value, expire)
}

func RedisStringSetKey(conn redis.Conn, key string, value string, expire int) error {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(value)
	if expire >= 0 {
		args = args.Add("EX")
		args = args.Add(expire)
	}
	_, err := conn.Do("SET", args...)
	return err
}

/*
* 获取string值
**/
func RedisStringGet(conn redis.Conn, redisType RedisType) (string, error) {
	if redisType.GetType() != KRedisString {
		return "", RedisStringTypeErr
	}
	return RedisStringGetKey(conn, redisType.GetKey())
}

func RedisStringGetKey(conn redis.Conn, key string) (string, error) {
	return redis.String(conn.Do("GET", key))
}
