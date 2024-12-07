package limiter

import "time"

type RateLimiter struct {
	storage        RateLimiterStorage
	rateLimitIP    int
	rateLimitToken int
	blockDuration  time.Duration
}

func NewRateLimiter(storage RateLimiterStorage, rateLimitIP, rateLimitToken int, blockDuration time.Duration) *RateLimiter {
	return &RateLimiter{
		storage:        storage,
		rateLimitIP:    rateLimitIP,
		rateLimitToken: rateLimitToken,
		blockDuration:  blockDuration,
	}
}

func (r *RateLimiter) Allow(key string, limit int) (bool, error) {
	count, err := r.storage.Increment(key)
	if err != nil {
		return false, err
	}

	if count == 1 {
		_ = r.storage.SetExpiration(key, r.blockDuration)
	}

	return count <= limit, nil
}
