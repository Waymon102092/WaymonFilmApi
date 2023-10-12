package cache

import (
	"Waymon_api/internal"
	"fmt"
	"time"
)

// 访问量

func MediaAccessKey(mediaId int64) string {
	return fmt.Sprintf("MediaAccess_%d", mediaId)
}

func IncrByMediaAccess(mediaId int64) bool {
	key := MediaAccessKey(mediaId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, 1, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetMediaAccess(mediaId, startTime, endTime int64) int64 {
	key := MediaAccessKey(mediaId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}

// 新用户量

func MediaMemberKey(mediaId int64) string {
	return fmt.Sprintf("MediaMember_%d", mediaId)
}

func IncrByMediaMember(mediaId int64) bool {
	key := MediaMemberKey(mediaId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, 1, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetMediaMember(mediaId, startTime, endTime int64) int64 {
	key := MediaMemberKey(mediaId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}

// 订单数量

func MediaOrderKey(mediaId int64) string {
	return fmt.Sprintf("MediaOrder_%d", mediaId)
}

func IncrByMediaOrder(mediaId int64) bool {
	key := MediaOrderKey(mediaId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, 1, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetMediaOrder(mediaId, startTime, endTime int64) int64 {
	key := MediaOrderKey(mediaId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}

// 销售额

func MediaMoneyKey(mediaId int64) string {
	return fmt.Sprintf("MediaMoney_%d", mediaId)
}

func IncrByMediaMoney(mediaId int64) bool {
	key := MediaMoneyKey(mediaId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, 1, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetMediaMoney(mediaId, startTime, endTime int64) int64 {
	key := MediaMoneyKey(mediaId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}

// 收入

func MediaAmountKey(mediaId int64) string {
	return fmt.Sprintf("MediaAmount_%d", mediaId)
}

func IncrByMediaAmount(mediaId int64) bool {
	key := MediaAmountKey(mediaId)
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateNow := time.Now()
	startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, location).Unix()
	_, err := internal.RedisClient.ZIncrBy(key, 1, fmt.Sprintf("%d", startTime)).Result()
	if err != nil {
		return false
	}
	return true
}

func GetMediaAmount(mediaId, startTime, endTime int64) int64 {
	key := MediaAmountKey(mediaId)
	result, err := internal.RedisClient.ZCount(key, fmt.Sprintf("%d", startTime), fmt.Sprintf("%d", endTime)).Result()
	if err != nil {
		return 0
	}
	return result
}
