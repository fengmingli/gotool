package distributedlock

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
 * @Author: LFM
 * @Date: 2024/6/22 10:50
 * @Since: 1.0.0
 * @Desc: 分布式锁，redis实现
 */

var ctx = context.Background()

type RedisLock struct {
	client  *redis.Client
	key     string
	value   string
	timeout time.Duration
	mu      sync.Mutex
	cancel  context.CancelFunc
}

// NewRedisLock creates a new RedisLock
func NewRedisLock(client *redis.Client, key string, value string, timeout time.Duration) *RedisLock {
	return &RedisLock{
		client:  client,
		key:     key,
		value:   value,
		timeout: timeout,
	}
}

// Lock tries to acquire the lock and starts the auto-renewal process
func (lock *RedisLock) Lock() (bool, error) {
	lock.mu.Lock()
	defer lock.mu.Unlock()

	ok, err := lock.client.SetNX(ctx, lock.key, lock.value, lock.timeout).Result()
	if !ok || err != nil {
		return ok, err
	}

	// Start the renewal process
	renewalCtx, cancel := context.WithCancel(ctx)
	lock.cancel = cancel
	go lock.renew(renewalCtx)

	return true, nil
}

// renew periodically extends the lock's expiration time
func (lock *RedisLock) renew(ctx context.Context) {
	ticker := time.NewTicker(lock.timeout / 2)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			lock.mu.Lock()
			lock.client.Expire(ctx, lock.key, lock.timeout)
			fmt.Println("Lock renewed!")
			lock.mu.Unlock()
		case <-ctx.Done():
			return
		}
	}
}

// Unlock releases the lock and stops the auto-renewal process
func (lock *RedisLock) Unlock() (bool, error) {
	lock.mu.Lock()
	defer lock.mu.Unlock()

	if lock.cancel != nil {
		lock.cancel()
	}

	script := redis.NewScript(`
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
    `)
	result, err := script.Run(ctx, lock.client, []string{lock.key}, lock.value).Result()
	return result.(int64) == 1, err
}
