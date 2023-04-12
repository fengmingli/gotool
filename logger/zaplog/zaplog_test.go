package zaplog

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"testing"
	"time"
)

/**
 * @Author: LFM
 * @Date: 2023/4/12 23:19
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestZap(t *testing.T) {
	// 默认配置
	conf := NewConfig()
	if err := Init(conf); err != nil {
		log.Println(err)
	}
	testLogger()
}

func testLogger() {
	startTime := time.Now()

	for i := 0; i < 1000; i++ {
		// 标准记录
		Logger.Info("Logger info")
		Logger.Warn("Logger warn")
		Logger.Error("Logger err")

		// 添加其他的key-value
		Logger.Info("Logger info", zap.Any("user", "alnk"))
		Logger.Warn("Logger Warn", zap.Any("user", "alnk"))
		Logger.Error("Logger Error", zap.Any("user", "alnk"))
	}

	fmt.Println("Logger执行时间: ", time.Since(startTime))
}
