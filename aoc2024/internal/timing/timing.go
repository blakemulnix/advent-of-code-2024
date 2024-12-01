package timing

import (
	"fmt"
	"time"
)

type TimedResult struct {
	Result    interface{}
	Duration  time.Duration
}

func Track(fn func() any) TimedResult {
	start := time.Now()
	result := fn()
	duration := time.Since(start)
	return TimedResult{
		Result:    result,
		Duration:  duration,
	}
}

func FormatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return fmt.Sprintf("%d µs", d.Microseconds())
	}
	return fmt.Sprintf("%.2f ms", float64(d.Microseconds()) / 1000.0)
} 