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

// Len will return total number of lines
func (l *lineLogger) Len() int {
	l.mu.Lock()
	result := len(l.lines)
	l.mu.Unlock()
	return result
}

// Get will get a specific line
func (l *lineLogger) Get(i int) string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.lines[i]
}

// GetFrom will get the logs from a position to the end.
// It also returns the last index of the log read.
func (l *lineLogger) GetFrom(from int) ([]string, int) {
	l.mu.Lock()
	lastLine := len(l.lines)
	result := make([]string, 0)
	for ; from < lastLine; from++ {
		result = append(result, l.lines[from])
	}
	l.mu.Unlock()
	return result, lastLine
}
