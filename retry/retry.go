package retry

import "time"

/**
 * @Author: LFM
 * @Date: 2022/7/20 23:16
 * @Since: 1.0.0
 * @Desc: TODO
 */

func UntilMaxRetry(f func() error, maxRetry int, interval time.Duration) error {
	var err error
	for i := 0; i < maxRetry; i++ {
		err = f()
		if err == nil {
			return nil
		}
		time.Sleep(interval)
	}
	return err
}
