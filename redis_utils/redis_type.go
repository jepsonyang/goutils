package redisutils

const (
	KRedisString = "String"
	KRedisSet    = "Set"
	KRedisZSet   = "ZSet"
	KRedisList   = "List"
	KRedisHash   = "Hash"
)

type RedisType interface {
	GetKey() string
	GetType() string
}
