package testRedisUtils

import (
	"fmt"
	"redisUtils"
)

type Jepson struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

func (param *Jepson) GetKey() string {
	return "jepson_hash"
}

func (param *Jepson) GetType() string {
	return redisUtils.KRedisHash
}

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
	if false {
		var nAdd, nRem int
		var err error

		nAdd, err = redisUtils.RedisSetAddByKey(conn, "jepson_set", []string{"jepson1"})
		fmt.Println(nAdd, err)

		nAdd, err = redisUtils.RedisSetAddByKey(conn, "jepson_set", []string{"jepson2"})
		fmt.Println(nAdd, err)

		nRem, err = redisUtils.RedisSetRemoveByKey(conn, "jepson_set", []string{"jepson2"})
		fmt.Println(nRem, err)
	}

	//redis_hash
	if true {
		var err error

		//RedisHashSetByKey
		mapHash := make(map[string]interface{})
		mapHash["name"] = "jepson"
		mapHash["age"] = 18
		err = redisUtils.RedisHashSetByKey(conn, "jepson_hash", mapHash)
		fmt.Println(err)
		fmt.Println("----------------------------------------------")

		//RedisHashSet
		var jepson1 Jepson
		jepson1.Name = "jepsonyang"
		jepson1.Age = 20
		err = redisUtils.RedisHashSet(conn, &jepson1, jepson1)
		fmt.Println(err)
		fmt.Println("----------------------------------------------")

		//RedisHashGetByKey
		var jepson2 Jepson
		err = redisUtils.RedisHashGetByKey(conn, "jepson_hash", &jepson2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%+v\n", jepson2)
		}
		fmt.Println("----------------------------------------------")
	}
}
