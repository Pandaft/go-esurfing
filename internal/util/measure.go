package util

import (
	"github.com/charmbracelet/log"
	"time"
)

// MeasureExecTime 测量执行耗时
func MeasureExecTime(start time.Time, log *log.Logger) {
	elapsed := time.Now().Sub(start)
	log.Debugf("耗时：%s", elapsed)
}
