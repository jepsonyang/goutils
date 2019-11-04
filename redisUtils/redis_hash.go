package redisUtils

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

var RedisHashTypeErr = errors.New("type is not hash")

/*
* 设置hash值
* @value 可以传入map[string]interface{}或者结构体对象
* @note 传入map[string]interface{}时，只会更新(或者覆盖)map包含的字段，其他字段的值不会被改变;
**/
func RedisHashSet(conn redis.Conn, redisType RedisType, value interface{}) error {
	if redisType.GetType() != KRedisHash {
		return RedisHashTypeErr
	}
	return RedisHashSetByKey(conn, redisType.GetKey(), value)
}

func RedisHashSetByKey(conn redis.Conn, key string, value interface{}) error {
	args := redis.Args{}
	args = args.Add(key)
	args = args.AddFlat(value)
	_, err := conn.Do("HMSET", args...)
	return err
}

/*
* 获取hash值
* @dst 必须传入结构体对象指针
**/
func RedisHashGet(conn redis.Conn, redisType RedisType, dst interface{}) error {
	if redisType.GetType() != KRedisHash {
		return RedisHashTypeErr
	}
	return RedisHashGetByKey(conn, redisType.GetKey(), dst)
}

func RedisHashGetByKey(conn redis.Conn, key string, dst interface{}) error {
	arrReply, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return err
	}
	err = redis.ScanStruct(arrReply, dst)
	if err != nil {
		return err
	}
	return nil
}
