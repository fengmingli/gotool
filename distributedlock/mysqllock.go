package distributedlock

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * @Author: LFM
 * @Date: 2024/6/22 10:50
 * @Since: 1.0.0
 * @Desc: 分布式锁，mysql实现
 */

type MySQLLock struct {
	db      *sql.DB
	key     string
	value   string
	timeout time.Duration
	mu      sync.Mutex
}

// NewMySQLLock creates a new MySQLLock
func NewMySQLLock(db *sql.DB, key string, value string, timeout time.Duration) *MySQLLock {
	return &MySQLLock{
		db:      db,
		key:     key,
		value:   value,
		timeout: timeout,
	}
}

// Lock tries to acquire the lock
func (lock *MySQLLock) Lock() (bool, error) {
	lock.mu.Lock()
	defer lock.mu.Unlock()

	query := fmt.Sprintf("INSERT INTO distributed_locks (lock_key, lock_value, expires_at) VALUES ('%s', '%s', NOW() + INTERVAL %d SECOND) ON DUPLICATE KEY UPDATE lock_value=VALUES(lock_value), expires_at=VALUES(expires_at)",
		lock.key, lock.value, int(lock.timeout.Seconds()))

	_, err := lock.db.Exec(query)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Unlock releases the lock
func (lock *MySQLLock) Unlock() (bool, error) {
	lock.mu.Lock()
	defer lock.mu.Unlock()

	query := fmt.Sprintf("DELETE FROM distributed_locks WHERE lock_key='%s' AND lock_value='%s'", lock.key, lock.value)

	result, err := lock.db.Exec(query)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
