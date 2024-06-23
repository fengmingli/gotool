package distributedlock

import (
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

/**
 * @Author: LFM
 * @Date: 2024/6/22 10:57
 * @Since: 1.0.0
 * @Desc: TODO
 */

//func TestRedisLock(t *testing.T) {
//	redisClient := redis.NewClient(&redis.Options{
//		Addr: "localhost:6379", // Redis server address
//	})
//	defer redisClient.Close()
//
//	// 创建一个 Redis 锁
//	lock := NewRedisLock(redisClient, "my_lock", "unique_value", 10*time.Second)
//
//	// 尝试获取锁
//	success, err := lock.Lock()
//	assert.NoError(t, err)
//	assert.True(t, success, "Failed to acquire lock")
//
//	if success {
//		t.Log("Lock acquired")
//
//		// 模拟一些需要锁保护的操作
//		time.Sleep(5 * time.Second)
//
//		// 释放锁
//		unlocked, err := lock.Unlock()
//		assert.NoError(t, err)
//		assert.True(t, unlocked, "Failed to release lock")
//
//		if unlocked {
//			t.Log("Lock released")
//		}
//	}
//}
//
//func TestRedisLockCompetition(t *testing.T) {
//	redisClient := redis.NewClient(&redis.Options{
//		Addr: "localhost:6379", // Redis server address
//	})
//	defer redisClient.Close()
//
//	var wg sync.WaitGroup
//	successCount := 0
//	mu := sync.Mutex{}
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			lock := NewRedisLock(redisClient, "compete_lock", "unique_value", 10*time.Second)
//			success, err := lock.Lock()
//			assert.NoError(t, err)
//
//			if success {
//				mu.Lock()
//				successCount++
//				mu.Unlock()
//
//				// 模拟一些需要锁保护的操作
//				time.Sleep(5 * time.Second)
//
//				unlocked, err := lock.Unlock()
//				assert.NoError(t, err)
//				assert.True(t, unlocked, "Failed to release lock")
//			}
//		}()
//	}
//
//	wg.Wait()
//	assert.Equal(t, 1, successCount, "More than one goroutine acquired the lock")
//}

func TestContinuousLockUnlock(t *testing.T) {
	// Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})
	defer redisClient.Close()

	// 创建一个 Redis 锁
	lock := NewRedisLock(redisClient, "continuous_lock", "unique_value", 10*time.Second)

	// 并发数和持续时间
	numClients := 1
	duration := 300 * time.Second

	var wg sync.WaitGroup
	wg.Add(numClients)

	// 启动多个并发客户端
	for i := 0; i < numClients; i++ {
		go func(id int) {
			defer wg.Done()
			ticker := time.NewTicker(3 * time.Second)
			defer ticker.Stop()

			stopTime := time.Now().Add(duration)
			for time.Now().Before(stopTime) {
				select {
				case <-ticker.C:
					success, err := lock.Lock()
					assert.NoError(t, err)
					if success {
						t.Logf("Client %d: Lock acquired", id)

						// 模拟一些需要锁保护的操作
						time.Sleep(500 * time.Millisecond)

						unlocked, err := lock.Unlock()
						assert.NoError(t, err)
						assert.True(t, unlocked, "Client %d: Failed to release lock", id)

						t.Logf("Client %d: Lock released", id)
					}
				}
			}
		}(i)
	}

	wg.Wait()
}
