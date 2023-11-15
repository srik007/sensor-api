package aggregators

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/srik007/sensor-api/infrastructure/cache"
)

type AggregatorSchedulerJob struct {
	CacheStore *cache.RedisCache
}

func NewScheduler(c *cache.RedisCache) *AggregatorSchedulerJob {
	return &AggregatorSchedulerJob{
		CacheStore: c}
}

func (a *AggregatorSchedulerJob) Run(c *gin.Context, queryHandler AggregatorQueryHandler) {
	go func() {
		for {
			time.Sleep(time.Duration(10) * time.Second)
			groupAggregatorResults := queryHandler.CollectGroupAggregators()
			for _, groupAggregator := range groupAggregatorResults {
				result, _ := json.Marshal(groupAggregator)
				a.CacheStore.Set(c, groupAggregator.GroupName, string(result))
			}
		}
	}()
}
