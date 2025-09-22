package flogger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetColor(t *testing.T) {
	red := getColorLog(ErrorLvl)
	assert.Equal(t, red, Red)

	yellow := getColorLog(WarnLvl)
	assert.Equal(t, yellow, Yellow)

	white := getColorLog(InfoLvl)
	assert.Equal(t, white, White)

	gray := getColorLog("")
	assert.Equal(t, gray, Gray)
}
