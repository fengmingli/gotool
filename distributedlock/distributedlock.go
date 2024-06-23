package distributedlock

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
 * @Author: LFM
 * @Date: 2024/6/22 10:50
 * @Since: 1.0.0
 * @Desc: TODO
 */

// DistributedLock defines the interface for a distributed lock
type DistributedLock interface {
	Lock() (bool, error)
	Unlock() (bool, error)
}

// LockType represents the type of distributed lock
type LockType int

const (
	// RedisLockType represents a Redis lock
	RedisLockType LockType = iota
	// MySQLLockType represents a MySQL lock
	MySQLLockType
)

// NewDistributedLock creates a new distributed lock based on the provided type
func NewDistributedLock(lockType LockType, client interface{}, key string, value string, timeout time.Duration) (DistributedLock, error) {
	switch lockType {
	case RedisLockType:
		redisClient, ok := client.(*redis.Client)
		if !ok {
			return nil, fmt.Errorf("invalid Redis client")
		}
		return NewRedisLock(redisClient, key, value, timeout), nil
	case MySQLLockType:
		db, ok := client.(*sql.DB)
		if !ok {
			return nil, fmt.Errorf("invalid MySQL client")
		}
		return NewMySQLLock(db, key, value, timeout), nil
	default:
		return nil, fmt.Errorf("unsupported lock type")
	}
}
