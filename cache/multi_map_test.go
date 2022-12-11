package cache

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @Author: LFM
 * @Date: 2022/12/11 16:41
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestMultiMap_Get(t *testing.T) {

	mMap := &MultiMap{data: make(map[string]map[string]string)}

	go func() {
		for i := 0; i < 10000; i++ {
			mMap.Put(fmt.Sprintf("key-%d", i), fmt.Sprintf("kkey-%d", i), fmt.Sprintf("value-%d", i))
			time.Sleep(1 * time.Millisecond)

		}
	}()
	go func() {
	A:
		for i := 0; i < 20000; i++ {
			get, ok := mMap.Get(fmt.Sprintf("key-%d", i), fmt.Sprintf("kkey-%d", i))
			if !ok {
				goto A
			}
			t.Log(fmt.Sprintf("key-%d", i), fmt.Sprintf("kkey-%d", i), "value:", get)
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			mMap.Delete(fmt.Sprintf("key-%d", i), fmt.Sprintf("kkey-%d", i))
			time.Sleep(2 * time.Millisecond)
		}
	}()
	time.Sleep(10000 * time.Second)
}
