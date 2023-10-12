package internal

import (
	"Waymon_api/pkg/log"
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"strings"
)

var RabbitMQ *amqp.Connection

func InitRabbitMQ() {
	name := viper.GetString("amqp.name")
	host := viper.GetString("amqp.host")
	port := viper.GetString("amqp.port")
	username := viper.GetString("amqp.username")
	password := viper.GetString("amqp.password")
	addr := strings.Join([]string{name, "://", username, ":", password, "@", host, ":", port, "/"}, "")
	conn, err := amqp.Dial(addr)
	if err != nil {
		log.WaymonLogger.Error("amqp Dial err:" + err.Error())
		zap.S().Error("amqp Dial err:" + err.Error())
		panic(err)
	}
	RabbitMQ = conn
	fmt.Println("RabbitMQ 初始化成功")
}
