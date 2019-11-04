package testRedisUtils

import (
	"fmt"
	"redisUtils"
)

func Test() {
	conn := GetRedisConnection()
	defer conn.Close()

	//redis_db
	if false {
		res, err := redisUtils.RedisExistByKey(conn, "NameList")
		fmt.Println(res, err)
	}

	//redis_string
	if false {
		err := redisUtils.RedisStringSetByKey(conn, "jepson", "This is test.", 60)
		fmt.Println(err)

		res, err := redisUtils.RedisStringGetByKey(conn, "jepson")
		fmt.Println(res, err)
	}

	//redis_set
	if true {
		var nAdd, nRem int
		var err error

		nAdd, err = redisUtils.RedisSetAddByKey(conn, "jepson_set", []string{"jepson1"})
		fmt.Println(nAdd, err)

		nAdd, err = redisUtils.RedisSetAddByKey(conn, "jepson_set", []string{"jepson2"})
		fmt.Println(nAdd, err)

		nRem, err = redisUtils.RedisSetRemoveByKey(conn, "jepson_set", []string{"jepson2"})
		fmt.Println(nRem, err)
	}
}