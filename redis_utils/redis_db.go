package redisutils

import "github.com/gomodule/redigo/redis"

/*
* 设置过期时间
* @return 设置成功返回1;当key不存在或者不能为key设置生存时间时(比如在低于2.1.3版本的Redis中你尝试更新key的生存时间),返回0;
**/
func RedisExpire(conn redis.Conn, redisType RedisType, expire int) (int, error) {
	return RedisExpireByKey(conn, redisType.GetKey(), expire)
}

func RedisExpireByKey(conn redis.Conn, key string, expire int) (int, error) {
	return redis.Int(conn.Do("EXPIRE", key, expire))
}

/*
* 是否存在
* @return 存在返回1，不存在返回0
**/
func RedisExist(conn redis.Conn, redisType RedisType) (int, error) {
	return RedisExistByKey(conn, redisType.GetKey())
}

func RedisExistByKey(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("EXISTS", key))
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
