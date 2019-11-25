//+build linux

package internal

func (l Level) String() string {
	switch l {
	case LevelLow:
		return "dialog-information"
	case LevelMiddle:
		return "dialog-warning"
	case LevelHigh:
		return "dialog-error"
	}
	return ""
}
