package wait

import "testing"

/**
 * @Author: LFM
 * @Date: 2022/7/16 22:39
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestUntil(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	Until(func() {
		t.Fatal("should not have been invoked")
	}, 0, ch)

	ch = make(chan struct{})
	called := make(chan struct{})
	go func() {
		Until(func() {
			called <- struct{}{}
		}, 0, ch)
		close(called)
	}()
	<-called
	close(ch)
	<-called
}
