package mathUtils

import (
	"math/rand"
	"time"
)

/*
* 生成指定范围内的随机数
* @return 生成的随机数范围[min, max)
**/
func RandomScope(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return min + random.Intn(max-min)
}
