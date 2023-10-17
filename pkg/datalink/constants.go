package datalink

import (
	"errors"
)

const (
	AckByte = byte(0x06)
)

var (
	ErrFrameTooShort = errors.New("frame too short for data link layer")
	ErrInvalidCRC    = errors.New("invalid CRC")
)
