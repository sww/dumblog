package dumblog

import (
	"fmt"
	"sync"
	"time"
)

type DumbLog struct {
	mu sync.Mutex
	Debug bool
}

func New(debug bool) *DumbLog {
	return &DumbLog{Debug: debug}
}

func now() string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	microseconds := now.Nanosecond() / 1e3

	return fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d.%04d", year, int(month), day, hour, min, sec, microseconds)
}

func (s *DumbLog) Print(v ...interface{}) {
	if !s.Debug {
		return
	}

	now := now()

	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println(now, fmt.Sprint(v...))
}

func (s *DumbLog) Printf(format string, v ...interface{}) {
	if !s.Debug {
		return
	}

	now := now()

	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println(now, fmt.Sprintf(format, v...))
}
