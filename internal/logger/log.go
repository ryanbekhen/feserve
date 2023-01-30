package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ryanbekhen/feserve/internal/timeutils"
)

type Config struct {
	Timezone string
}

type Log struct {
	location *time.Location
}

var (
	log     *Log
	logOnce sync.Once
)

func New(config ...Config) *Log {
	logOnce.Do(func() {
		if len(config) == 0 {
			loc := timeutils.Location("UTC")
			log = &Log{loc}
			return
		}
		loc := timeutils.Location(config[0].Timezone)
		log = &Log{loc}
	})

	return log
}

func (l *Log) buildMessage(a ...any) string {
	logtime := fmt.Sprintf("[%s]", time.Now().In(l.location).Format(time.RFC3339))
	message := fmt.Sprint(a...)
	return fmt.Sprintf("%s - %s", logtime, message)
}

func (l *Log) Info(a ...any) {
	fmt.Println(l.buildMessage(a...))
}

func (l *Log) Error(a ...any) {
	fmt.Println(l.buildMessage(a...))
}

func (l *Log) Fatal(a ...any) {
	fmt.Println(l.buildMessage(a...))
	os.Exit(1)
}
