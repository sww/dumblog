package dumblog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type DumbLog struct {
	Debug bool
	file  *os.File
	mu    sync.Mutex
}

func New(debug bool) *DumbLog {
	return &DumbLog{
		Debug: debug,
		file:  os.Stdout,
	}
}

func now() string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	microseconds := now.Nanosecond() / 1e3

	return fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d.%04d", year, int(month), day, hour, min, sec, microseconds)
}

func (s *DumbLog) SetOutput(f *os.File) {
	s.file = f
}

func (s *DumbLog) Print(v ...interface{}) {
	if !s.Debug {
		return
	}

	now := now()

	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Fprintln(s.file, now, fmt.Sprint(v...))
}

func (s *DumbLog) Printf(format string, v ...interface{}) {
	if !s.Debug {
		return
	}

	now := now()

	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Fprintln(s.file, now, fmt.Sprintf(format, v...))
}
