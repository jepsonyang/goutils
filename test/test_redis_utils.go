package test

func Test() {
	conn :=	GetRedisConnection()
	defer conn.Close()


}