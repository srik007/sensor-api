package cache

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client  *redis.Client
	Context *gin.Context
}

type AggregatedResult struct {
	AverageTemperature  float64
	AverageTransparency float64
	GroupName           string
}

func NewCache(cacheUrl string) *RedisCache {
	return &RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr: cacheUrl,
		})}
}

func (r *RedisCache) VerifyCache(c *gin.Context, groupName string) *AggregatedResult {
	var result AggregatedResult
	val, err := r.Client.Get(c, groupName).Result()
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (r *RedisCache) Set(c *gin.Context, k string, v string) {
	r.Client.Set(c, k, v, 10*time.Second)

}
