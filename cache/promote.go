package cache

import (
	"Waymon_api/internal"
	"fmt"
	"time"
)

// 访问量

func PromoteAccessKey(promoteId int64) string {
	return fmt.Sprintf("PromoteAccess_%d", promoteId)
}

func IncrByPromoteAccess(promoteId int64) bool {
	key := PromoteAccessKey(promoteId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, 1, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetPromoteAccess(promoteId, startTime, endTime int64) int64 {
	key := PromoteAccessKey(promoteId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}

// 新用户量

func PromoteMemberKey(promoteId int64) string {
	return fmt.Sprintf("PromoteMember_%d", promoteId)
}

func IncrByPromoteMember(promoteId int64) bool {
	key := PromoteMemberKey(promoteId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, 1, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetPromoteMember(promoteId, startTime, endTime int64) int64 {
	key := PromoteMemberKey(promoteId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}

// 订单数量

func PromoteOrderKey(promoteId int64) string {
	return fmt.Sprintf("PromoteOrder_%d", promoteId)
}

func IncrByPromoteOrder(promoteId int64) bool {
	key := PromoteOrderKey(promoteId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, 1, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetPromoteOrder(promoteId, startTime, endTime int64) int64 {
	key := PromoteOrderKey(promoteId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}

// 销售额

func PromoteMoneyKey(promoteId int64) string {
	return fmt.Sprintf("PromoteMoney_%d", promoteId)
}

func IncrByPromoteMoney(promoteId int64, money float64) bool {
	key := PromoteMoneyKey(promoteId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, money, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetPromoteMoney(promoteId, startTime, endTime int64) int64 {
	key := PromoteMoneyKey(promoteId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}

// 收入

func PromoteAmountKey(promoteId int64) string {
	return fmt.Sprintf("PromoteAmount_%d", promoteId)
}

func IncrByPromoteAmount(promoteId int64, money float64) bool {
	key := PromoteAmountKey(promoteId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, money, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetPromoteAmount(promoteId, startTime, endTime int64) int64 {
	key := PromoteAmountKey(promoteId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}
