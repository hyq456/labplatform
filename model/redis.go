package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"labplatform/utils"
	"log"
	"sync"
)

var (
	Ctx             = context.Background()
	linkRedisMethod sync.Once
	DbRedis         *redis.Client
)

func InitRedis() {
	linkRedisMethod.Do(func() {
		addr := fmt.Sprintf("%s:%s", utils.RedisHost, utils.RedisPort)
		//连接数据库
		DbRedis = redis.NewClient(&redis.Options{
			Addr:     addr,                // 对应的ip以及端口号
			Password: utils.RedisPassWord, // 数据库的密码
			DB:       utils.RedisDB,       // 数据库的编号，默认的话是0
		})
		// 连接测活
		_, err := DbRedis.Ping(Ctx).Result()
		if err != nil {
			log.Println(err)
			panic(err)
		}
		fmt.Println("连接Redis成功")
	})
}
