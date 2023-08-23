package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// LoggerConfig defines the configuration for the logger
type LoggerConfig struct {
	FilenamePrefix string        // Prefix for the filename
	MaxLines       int           // Maximum number of lines before rotating
	RotationTime   time.Duration // Duration before rotating
}

// Logger is the main struct for logging
type Logger struct {
	config      LoggerConfig
	currentFile *os.File
	lineCount   int
	mu          sync.Mutex
}

// NewLogger creates a new logger with the given configuration
func NewLogger(config LoggerConfig) (*Logger, error) {
	logger := &Logger{
		config: config,
	}
	err := logger.rotateFile()
	if err != nil {
		return nil, err
	}
	go logger.timeBasedRotation() // Start time-based rotation
	return logger, nil
}

// Log logs the object as a JSON line
func (l *Logger) Log(obj interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	line, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	// Write the line
	_, err = fmt.Fprintln(l.currentFile, string(line))
	if err != nil {
		return err
	}

	l.lineCount++
	if l.lineCount >= l.config.MaxLines {
		return l.rotateFile() // Rotate file based on line count
	}

	return nil
}

// rotateFile rotates the log file
func (l *Logger) rotateFile() error {
	if l.currentFile != nil {
		l.currentFile.Close()
	}

	filename := fmt.Sprintf("%s_%s.log", l.config.FilenamePrefix, time.Now().Format(time.RFC3339))
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	l.currentFile = file
	l.lineCount = 0
	return nil
}

// timeBasedRotation periodically rotates the log file based on time
func (l *Logger) timeBasedRotation() {
	for range time.Tick(l.config.RotationTime) {
		l.mu.Lock()
		l.rotateFile() // Rotate file based on time
		l.mu.Unlock()
	}
}
