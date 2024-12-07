package limiter

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr, password string) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	return &RedisStorage{client: client}
}

func (r *RedisStorage) Increment(key string) (int, error) {
	val, err := r.client.Incr(context.Background(), key).Result()
	return int(val), err
}

func (r *RedisStorage) SetExpiration(key string, ttl time.Duration) error {
	return r.client.Expire(context.Background(), key, ttl).Err()
}

func (r *RedisStorage) GetCount(key string) (int, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

func (r *RedisStorage) Reset(key string) error {
	return r.client.Del(context.Background(), key).Err()
}
