package limiter

import "time"

type RateLimiterStorage interface {
	Increment(key string) (int, error)
	SetExpiration(key string, ttl time.Duration) error
	GetCount(key string) (int, error)
	Reset(key string) error
}
