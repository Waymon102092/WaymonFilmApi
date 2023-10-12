package cache

import (
	"Waymon_api/internal"
	"fmt"
	"time"
)

type RedisStore struct {
}

func (r RedisStore) CodeKey(id string) string {
	return fmt.Sprintf("Code:%s", id)
}

func (r RedisStore) Set(id, value string) error {
	key := r.CodeKey(id)
	_, err := internal.RedisClient.Set(key, value, time.Minute*5).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisStore) Get(id string, clear bool) string {
	key := r.CodeKey(id)
	result, err := internal.RedisClient.Get(key).Result()
	if err != nil {
		return ""
	}
	if clear {
		_, err = internal.RedisClient.Del(key).Result()
		if err != nil {
			return ""
		}
	}
	return result
}

func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
