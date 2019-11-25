//+build darwin

package internal

func (l Level) String() string {
	switch l {
	case LevelLow:
		//  display dialog "with icon note(= with icon 1)" with icon note
		return "note"
	case LevelMiddle:
		//  display dialog "with icon caution(= with icon 2)" with icon caution
		return "caution"
	case LevelHigh:
		//  display dialog "with icon stop(= with icon 0)" with icon stop
		return "stop"
	}
	return ""
}
