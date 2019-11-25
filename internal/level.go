package internal

import (
	"fmt"
)

type Level int

const (
	LevelLow = 1 + iota
	LevelMiddle
	LevelHigh
)

func (l *Level) Set(str string) error {
	next, err := ParseLevel(str)
	if err != nil {
		return err
	}
	*l = next
	return nil
}

func ParseLevel(str string) (Level, error) {
	switch str {
	case "low":
		return LevelLow, nil
	case "middle":
		return LevelMiddle, nil
	case "high":
		return LevelHigh, nil
	}
	return Level(0), fmt.Errorf("unsupported value %s; use low, middle or high alternatively", str)
}
