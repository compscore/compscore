package logging

import (
	"bytes"

	"github.com/sirupsen/logrus"
)

type customFormatter struct {
	logrus.Formatter
}

// Format an entry with \n replaced by \r\n
func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	out, err := f.Formatter.Format(entry)
	if err != nil {
		return nil, err
	}
	// Replace \n with \r\n
	return bytes.ReplaceAll(out, []byte("\n"), []byte("\r\n")), nil
}

// Init initializes the logger
func Init() {
	logrus.SetFormatter(&customFormatter{&logrus.TextFormatter{}})
	logrus.SetLevel(logrus.DebugLevel)
}
