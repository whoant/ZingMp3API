package redis_wrapper

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisWrapper struct {
	Client *redis.Client
	ctx    context.Context
}

func NewRedisWrapper(redisClient *redis.Client, ctx context.Context) *RedisWrapper {
	return &RedisWrapper{
		Client: redisClient,
		ctx:    ctx,
	}
}

func (r *RedisWrapper) Get(key string) (interface{}, error) {
	return r.Client.Get(r.ctx, key).Result()
}

func (r *RedisWrapper) Set(key string, value interface{}, duration time.Duration) error {
	return r.Client.Set(r.ctx, key, value, duration).Err()
}

func (r *RedisWrapper) SetWithoutDuration(key string, value interface{}) error {
	return r.Client.Set(r.ctx, key, value, 0).Err()
}

// IsExist check if key exists then return true else false
func (r *RedisWrapper) IsExist(key string) (bool, error) {
	ttl, err := r.Client.TTL(r.ctx, key).Result()
	if err != nil {
		return false, err
	}

	// -2 if the key does not exist
	// -1 if the key exists but has no associated expire
	if ttl <= 0 && ttl != -1 {
		return false, nil
	}

	return true, nil
}
