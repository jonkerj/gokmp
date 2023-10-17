package application

import (
	"errors"
)

const (
	CommandGetType     = byte(0x01)
	CommandGetSerialNo = byte(0x02)
	CommandGetRegister = byte(0x10)
)

var (
	ErrCouldNotDecodeFloat = errors.New("could not decode float")
	ErrShortRegister       = errors.New("not enough data to decode register")
	ErrShortSerial         = errors.New("not enough data to decode serial number")
	ErrLondSerial          = errors.New("too much data to decode serial number")
	ErrShortGetType        = errors.New("not enough data to decode type")
	ErrLongGetType         = errors.New("too much data to decode type")
)
