package config

import "Waymon_api/internal"

const (
	Direct  = "direct"
	Topic   = "topic"
	Fanout  = "fanout"
	Headers = "headers"

	OrderStatusExchange   = "exchange.order_status"
	OrderStatusQueue      = "queue.order_status"
	OrderStatusRoutingKey = "key.order_status"
	OrderStatusConsumer   = "consumer.order_status"

	OrderSuccessExchange   = "exchange.order_success"
	OrderSuccessQueue      = "queue.order_success"
	OrderSuccessRoutingKey = "key.order_success"
	OrderSuccessConsumer   = "consumer.order_success"

	OrderRefundExchange   = "exchange.order_refund"
	OrderRefundQueue      = "queue.order_refund"
	OrderRefundRoutingKey = "key.order_refund"
	OrderRefundConsumer   = "consumer.order_refund"

	OrderCompeteExchange   = "exchange.order_compete"
	OrderCompeteQueue      = "queue.order_compete"
	OrderCompeteRoutingKey = "key.order_compete"
	OrderCompeteConsumer   = "consumer.order_compete"

	RefundSuccessExchange   = "exchange.refund_success"
	RefundSuccessQueue      = "queue.refund_success"
	RefundSuccessRoutingKey = "key.refund_success"
	RefundSuccessConsumer   = "consumer.refund_success"
)

func init() {
	internal.InitViper()
}

func InitConfig() {
	internal.InitDB()
	internal.InitRedis()
	//internal.InitRabbitMQ()
	//internal.InitRedsync()
	//internal.InitES()
}
