package util

import (
	"strings"
	"sync"
)

// LineLogger will get each line of written buffer
type LineLogger struct {
	lines []string
	mu    sync.Mutex
}

// NewLineLogger will create a new LineLogger with given initial value.
// Generally speaking, an uninitialized LineLogger can is safe to use.
func NewLineLogger(initial []string) *LineLogger {
	return &LineLogger{
		lines: initial,
	}
}

func (l *LineLogger) Write(b []byte) (int, error) {
	lines := strings.Split(string(b), "\n")
	l.mu.Lock()
	for _, line := range lines {
		if line == "" {
			continue
		}
		l.lines = append(l.lines, line)
	}
	l.mu.Unlock()
	return len(b), nil
}

// Len will return total number of lines
func (l *LineLogger) Len() int {
	l.mu.Lock()
	lines := len(l.lines)
	l.mu.Unlock()
	return lines
}

// Tail will get the last at last count lines of the log
func (l *LineLogger) Tail(count int) []string {
	l.mu.Lock()
	result := make([]string, 0, count)
	start := len(l.lines) - count
	if start < 0 {
		start = 0
	}
	for ; start < len(l.lines); start++ {
		result = append(result, l.lines[start])
	}
	l.mu.Unlock()
	return result
}

// Head will get the first at last count lines of the log
func (l *LineLogger) Head(count int) []string {
	l.mu.Lock()
	result := make([]string, 0, count)
	for i := 0; i < count && i < len(l.lines); i++ {
		result = append(result, l.lines[i])
	}
	l.mu.Unlock()
	return result
}

// GetFrom will get the logs from a position to the end.
// It also returns the last index of the log read.
func (l *LineLogger) GetFrom(from int) ([]string, int) {
	l.mu.Lock()
	lastLine := len(l.lines)
	result := make([]string, 0)
	for ; from < lastLine; from++ {
		result = append(result, l.lines[from])
	}
	l.mu.Unlock()
	return result, lastLine
}

// Copy will copy the logger output as a string slice
func (l *LineLogger) Copy() []string {
	l.mu.Lock()
	result := make([]string, len(l.lines))
	copy(result, l.lines)
	l.mu.Unlock()
	return result
}
