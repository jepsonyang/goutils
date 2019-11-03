package test

import (
	"fmt"
	redisutils "github.com/jepsonyang/goutils/redis_utils"
)

func TestRedisUtils() {
	conn := GetRedisConnection()
	defer conn.Close()

	//redis_db
	if false {
		res, err := redisutils.RedisExistKey(conn, "NameList")
		fmt.Println(res, err)
	}

	//redis_string
	if true {
		err := redisutils.RedisStringSetKey(conn, "jepson", "This is test.", 60)
		fmt.Println(err)

		res, err := redisutils.RedisStringGetKey(conn, "jepson")
		fmt.Println(res, err)
	}
}