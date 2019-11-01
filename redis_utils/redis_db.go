package redisutils

import "github.com/gomodule/redigo/redis"

/*
* 设置过期时间
* @返回值 设置成功返回1;当key不存在或者不能为key设置生存时间时(比如在低于2.1.3版本的Redis中你尝试更新key的生存时间),返回0;
**/
func RedisExpire(conn redis.Conn, redisType RedisType, expire int) (int, error) {
	return redis.Int(conn.Do("EXPIRE", redisType.GetKey(), expire))
}

/*
* 是否存在
* @返回值 存在返回1，不存在返回0
**/
func RedisExist(conn redis.Conn, redisType RedisType) (int, error) {
	return redis.Int(conn.Do("EXISTS", redisType.GetKey()))
}

/*
* 删除
* @返回值 被删除的key个数
**/
func RedisDel(conn redis.Conn, redisTypes []RedisType) (int, error) {
	keys := make([]string, 0)
	for _, v := range redisTypes {
		keys = append(keys, v.GetKey())
	}
	return RedisDelKey(conn, keys)
}

/*
* 删除指定keys
* @返回值 被删除的key个数
**/
func RedisDelKey(conn redis.Conn, keys []string) (int, error) {
	if len(keys) <= 0 {
		return 0, nil
	}

	args := redis.Args{}
	for _, v := range keys {
		args = args.Add(v)
	}

	return redis.Int(conn.Do("DEL", args...))
}
