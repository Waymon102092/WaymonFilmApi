package internal

import (
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/spf13/viper"
)

var RedSync *redsync.Mutex

func InitRedsync() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	addr := fmt.Sprintf("%s:%s", host, port)

	host1 := viper.GetString("redis1.host")
	port1 := viper.GetString("redis1.port")
	addr1 := fmt.Sprintf("%s:%s", host1, port1)

	host2 := viper.GetString("redis2.host")
	port2 := viper.GetString("redis2.port")
	addr2 := fmt.Sprintf("%s:%s", host2, port2)

	client := goredislib.NewClusterClient(&goredislib.ClusterOptions{
		Addrs: []string{addr, addr1, addr2},
	})
	pool := goredis.NewPool(client)
	//创建实例
	rs := redsync.New(pool)
	//互斥锁
	mutexName := "waymon_mutex"
	//创建基于key的互斥锁
	RedSync = rs.NewMutex(mutexName)
}
