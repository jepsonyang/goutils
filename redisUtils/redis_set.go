package redisUtils

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

var RedisSetTypeErr = errors.New("type is not set")

/*
* 添加成员到集合
* @return 返回被添加到集合的新元素数量,不包括被忽略的元素;
**/
func RedisSetAdd(conn redis.Conn, redisType RedisType, values []string) (int, error) {
	if redisType.GetType() != KRedisSet {
		return 0, RedisSetTypeErr
	}
	return RedisSetAddByKey(conn, redisType.GetKey(), values)
}

func RedisSetAddByKey(conn redis.Conn, key string, values []string) (int, error) {
	if len(values) <= 0 {
		return 0, nil
	}
	args := redis.Args{}
	args = args.Add(key)
	for _, v := range values {
		args = args.Add(v)
	}
	return redis.Int(conn.Do("SADD", args...))
}

/*
* 删除集合中的成员
* @return 返回被删除的元素数量,不包括被忽略的元素;
**/
func RedisSetRemove(conn redis.Conn, redisType RedisType, values []string) (int, error) {
	if redisType.GetType() != KRedisSet {
		return 0, RedisSetTypeErr
	}
	return RedisSetRemoveByKey(conn, redisType.GetKey(), values)
}

func RedisSetRemoveByKey(conn redis.Conn, key string, values []string) (int, error) {
	args := redis.Args{}
	args = args.Add(key)
	for _, v := range values {
		args = args.Add(v)
	}
	return redis.Int(conn.Do("SREM", args...))
}

/*
* 判断member是否为集合key的成员
* @return 如果member元素是集合的成员,返回true;如果member不是集合的成员,或key不存在,返回false
**/
func RedisSetIsMember(conn redis.Conn, redisType RedisType, member string) (bool, error) {
	if redisType.GetType() != KRedisSet {
		return false, RedisSetTypeErr
	}
	return RedisSetIsMemberByKey(conn, redisType.GetKey(), member)
}

func RedisSetIsMemberByKey(conn redis.Conn, key string, member string) (bool, error) {
	result, err := redis.Int(conn.Do("SISMEMBER", key, member))
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

/*
* 获取集合中的所有成员
**/
func RedisSetMembers(conn redis.Conn, redisType RedisType) ([]string, error) {
	if redisType.GetType() != KRedisSet {
		return nil, RedisSetTypeErr
	}
	return RedisSetMembersByKey(conn, redisType.GetKey())
}

func RedisSetMembersByKey(conn redis.Conn, key string) ([]string, error) {
	return redis.Strings(conn.Do("SMEMBERS", key))
}

/*
* 获取集合的成员总数
**/
func RedisSetMemberCount(conn redis.Conn, redisType RedisType) (int, error) {
	if redisType.GetType() != KRedisSet {
		return 0, RedisSetTypeErr
	}
	return RedisSetMemberCountByKey(conn, redisType.GetKey())
}

func RedisSetMemberCountByKey(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("SCARD", key))
}
