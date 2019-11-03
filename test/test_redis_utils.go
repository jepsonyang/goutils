package test

import (
	"fmt"
	redisutils "github.com/jepsonyang/goutils/redis_utils"
)

func TestRedisUtils() {
	conn :=	GetRedisConnection()
	defer conn.Close()

	//redis_db
	if true {
		res, err := redisutils.RedisExistKey(conn, "NameList")
		fmt.Println(res, err)
	}

}