module github.com/jepsonyang/goutils/test

go 1.12

require (
	github.com/gomodule/redigo v2.0.0+incompatible
	redisUtils v0.0.0
)

replace redisUtils => ../redisUtils
