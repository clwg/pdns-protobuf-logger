package writer

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type LoggerConfig struct {
	FilenamePrefix string
	MaxLines       int
	RotationTime   time.Duration
}

type Logger struct {
	config      LoggerConfig
	currentFile *os.File
	lineCount   int
	mu          sync.Mutex
}

func NewLogger(config LoggerConfig) (*Logger, error) {
	logger := &Logger{
		config: config,
	}
	err := logger.rotateFile()
	if err != nil {
		return nil, err
	}
	go logger.timeBasedRotation()
	return logger, nil
}

func (l *Logger) Log(obj interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	line, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(l.currentFile, string(line))
	if err != nil {
		return err
	}

	l.lineCount++
	if l.lineCount >= l.config.MaxLines {
		return l.rotateFile()
	}

	return nil
}

// rotateFile rotates and compresses the log file
func (l *Logger) rotateFile() error {
	if l.currentFile != nil {
		// Open a gzip writer to compress the existing file
		gzipFilename := fmt.Sprintf("%s.gz", l.currentFile.Name())
		gzipFile, err := os.Create(gzipFilename)
		if err != nil {
			return err
		}
		writer := gzip.NewWriter(gzipFile)
		l.currentFile.Seek(0, 0) // Move to start of file
		if _, err := io.Copy(writer, l.currentFile); err != nil {
			return err
		}
		writer.Close()
		gzipFile.Close()
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

func (l *Logger) timeBasedRotation() {
	for range time.Tick(l.config.RotationTime) {
		l.mu.Lock()
		l.rotateFile()
		l.mu.Unlock()
	}
}
