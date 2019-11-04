package redisUtils

import (
	"github.com/gomodule/redigo/redis"
)

/*
* 设置过期时间
* @return 设置成功返回true;当key不存在或者不能为key设置生存时间时(比如在低于2.1.3版本的Redis中你尝试更新key的生存时间),返回false;
**/
func RedisExpire(conn redis.Conn, redisType RedisType, expire int) (bool, error) {
	return RedisExpireByKey(conn, redisType.GetKey(), expire)
}

func RedisExpireByKey(conn redis.Conn, key string, expire int) (bool, error) {
	result, err := redis.Int(conn.Do("EXPIRE", key, expire))
	return result == 1, err
}

/*
* 是否存在
* @return 存在返回true，不存在返回false
**/
func RedisExist(conn redis.Conn, redisType RedisType) (bool, error) {
	return RedisExistByKey(conn, redisType.GetKey())
}

func RedisExistByKey(conn redis.Conn, key string) (bool, error) {
	result, err := redis.Int(conn.Do("EXISTS", key))
	return result == 1, err
}

/*
* 删除
* @return 被删除的key个数
**/
func RedisDel(conn redis.Conn, redisTypes []RedisType) (int, error) {
	keys := make([]string, 0)
	for _, v := range redisTypes {
		keys = append(keys, v.GetKey())
	}
	return RedisDelByKey(conn, keys)
}

func RedisDelByKey(conn redis.Conn, keys []string) (int, error) {
	if len(keys) <= 0 {
		return 0, nil
	}

	args := redis.Args{}
	for _, v := range keys {
		args = args.Add(v)
	}

	return redis.Int(conn.Do("DEL", args...))
}

/*
* 遍历
* @cursor 遍历使用的游标，开始遍历填0
* @pattern 指定一个glob风格的模式参数,只返回和给定模式相匹配的元素,如果不需要现在，直接填""即可;
* @return nextCursor 下次迭代使用的游标值,0表示遍历完成; keys 本次遍历到的key列表
**/
func RedisScan(conn redis.Conn, cursor int, pattern string) (nextCursor int, keys []string, err error) {
	args := redis.Args{}
	args = args.Add(cursor)
	if len(pattern) > 0 {
		args = args.Add("MATCH")
		args = args.Add(pattern)
	}

	var arrValue []interface{}
	arrValue, err = redis.Values(conn.Do("SCAN", args...))
	if err != nil {
		return
	}

	nextCursor, _ = redis.Int(arrValue[0], nil)
	keys, _ = redis.Strings(arrValue[1], nil)

	return
}
