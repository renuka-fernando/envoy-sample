package sample

import (
	"log"
)

type Logger struct {
	Debug bool
}

func (l Logger) Debugf(format string, args ...interface{}) {
	if l.Debug {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

func (l Logger) Infof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format+"\n", args...)
}

func (l Logger) Warnf(format string, args ...interface{}) {
	log.Printf("[WARN] "+format+"\n", args...)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format+"\n", args...)
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	log.Printf("[FATAL] "+format+"\n", args...)
}
