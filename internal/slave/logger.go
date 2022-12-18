package slave

import (
	"strings"
	"sync"
)

// lineLogger will get each line of written buffer
type lineLogger struct {
	lines []string
	mu    sync.Mutex
}

func (l *lineLogger) Write(b []byte) (int, error) {
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

func (l *lineLogger) Len() int {
	l.mu.Lock()
	result := len(l.lines)
	l.mu.Unlock()
	return result
}

func (l *lineLogger) Get(i int) string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.lines[i]
}
