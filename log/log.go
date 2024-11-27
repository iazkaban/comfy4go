package log

import (
	stdLog "log"
)

type SimpleLog struct {
}

func (SimpleLog) Info(msg ...any) {
	stdLog.Default().Println(msg)
}

func (SimpleLog) Warn(msg ...any) {
	stdLog.Default().Println(msg)
}

func (SimpleLog) Error(msg ...any) {
	stdLog.Default().Println(msg)
}
