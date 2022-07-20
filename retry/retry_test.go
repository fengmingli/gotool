package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/**
 * @Author: LFM
 * @Date: 2022/7/20 23:17
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestUntilMaxRetry(t *testing.T) {
	err := UntilMaxRetry(func() error {
		fmt.Println("===============")
		return errors.New("test retry")
	}, 5, 5*time.Second)

	if err != nil {
		_ = fmt.Errorf("test:%v", err.Error())
	}
}
