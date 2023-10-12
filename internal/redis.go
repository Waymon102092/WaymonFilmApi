package internal

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	addr := fmt.Sprintf("%s:%s", host, port)
	password := viper.GetString("redis.password")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		PoolTimeout:  4 * time.Second,
		OnConnect: func(conn *redis.Conn) error {
			return nil
		},
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		zap.S().Error("RedisCluster.Ping err" + err.Error())
		fmt.Println("RedisCluster.Ping err" + err.Error())
	}
	fmt.Println("Redis初始化完成。。。")
}
