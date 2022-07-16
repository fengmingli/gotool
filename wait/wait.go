package wait

import (
	"math/rand"
	"time"
)

/**
 * @Author: LFM
 * @Date: 2022/7/16 22:38
 * @Since: 1.0.0
 * @Desc: TODO
 */

func Until(f func(), period time.Duration, stopCh <-chan struct{}) {
	JitterUntil(f, period, 0.0, true, stopCh)
}

func JitterUntil(f func(), period time.Duration, jitterFactor float64, sliding bool, stopCh <-chan struct{}) {
	var t *time.Timer
	var sawTimeout bool

	for {
		select {
		case <-stopCh:
			return
		default:
		}

		jitteredPeriod := period
		if jitterFactor > 0.0 {
			jitteredPeriod = Jitter(period, jitterFactor)
		}

		if !sliding {
			t = resetOrReuseTimer(t, jitteredPeriod, sawTimeout)
		}

		func() {
			f()
		}()

		if sliding {
			t = resetOrReuseTimer(t, jitteredPeriod, sawTimeout)
		}
		select {
		case <-stopCh:
			return
		case <-t.C:
			sawTimeout = true
		}
	}
}

func Jitter(duration time.Duration, maxFactor float64) time.Duration {
	if maxFactor <= 0.0 {
		maxFactor = 1.0
	}
	wait := duration + time.Duration(rand.Float64()*maxFactor*float64(duration))
	return wait
}

func resetOrReuseTimer(t *time.Timer, d time.Duration, sawTimeout bool) *time.Timer {
	if t == nil {
		return time.NewTimer(d)
	}
	if !t.Stop() && !sawTimeout {
		<-t.C
	}
	t.Reset(d)
	return t
}
