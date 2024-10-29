package registry

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func (r *registry) NewStatisticsDbClient() (*redis.Client, error) {
	redisMetricsMockHost, err := r.conf.String("redisMetricsMockHost")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisMetricsMockHost, // Connect to local Redis server
		Password: "",                   // No password set
		DB:       0,                    // Use default DB
	})

	// Test the connection
	_, err = rdb.Ping(ctx).Result()

	// Set LRU eviction policy
	rdb.ConfigSet(ctx, "maxmemory-policy", "allkeys-lru")

	return rdb, err
}
