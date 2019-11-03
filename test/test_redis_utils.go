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
		res, err := redisutils.RedisExistByKey(conn, "NameList")
		fmt.Println(res, err)
	}

	//redis_string
	if false {
		err := redisutils.RedisStringSetByKey(conn, "jepson", "This is test.", 60)
		fmt.Println(err)

		res, err := redisutils.RedisStringGetByKey(conn, "jepson")
		fmt.Println(res, err)
	}

	//redis_set
	if true {
		var nAdd, nRem int
		var err error

		nAdd, err = redisutils.RedisSetAddByKey(conn, "jepson_set", []string{"jepson1"})
		fmt.Println(nAdd, err)

		nAdd, err = redisutils.RedisSetAddByKey(conn, "jepson_set", []string{"jepson2"})
		fmt.Println(nAdd, err)

		nRem, err = redisutils.RedisSetRemoveByKey(conn, "jepson_set", []string{"jepson2"})
		fmt.Println(nRem, err)
	}
}