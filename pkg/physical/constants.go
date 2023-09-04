package physical

import "errors"

const (
	StartByteToMeter   = byte(0x80)
	StartByteFromMeter = byte(0x40)
	StopByte           = byte(0x0d)
	EscapeByte         = byte(0x1b)
	ackByte            = byte(0x06) // cannot import from datalink, would lead to import cycle
)

var (
	escapes = map[byte]struct{}{
		StartByteToMeter:   {},
		StartByteFromMeter: {},
		StopByte:           {},
		EscapeByte:         {},
		ackByte:            {},
	}

	ErrInvalidStartStopByte = errors.New("start/stop bytes absent or in weird positions")
)
