package logger_custome

import (
	"testing"
)

func Test_Log_Info(t *testing.T) {
	logger := New(Info)

	logger.Info("Не розовый текст")
}

func Test_Log_Warn(t *testing.T) {
	logger := New(Info)

	logger.Warn("Не розовый текст")
}

func Test_Log_Error(t *testing.T) {
	logger := New(Info)

	logger.Error("Не розовый текст")
}
