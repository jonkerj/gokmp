package datalink

import (
	"github.com/jonkerj/gokmp/pkg/physical"
	"github.com/snksoft/crc"
)

type (
	Destination byte

	FrameType uint

	Frame struct {
		DestinationAddress Destination
		FrameType          FrameType
		Data               []byte
		Checksum           uint16
	}
)

const (
	FrameTypeUnknown FrameType = iota
	FrameTypeCommand
	FrameTypeResponse
)

const (
	DestinationHeatMeter        Destination = 0x3f
	DestinationLoggerTopModule  Destination = 0x7f
	DestinationLoggerBaseModule Destination = 0xbf
)

func (f *FrameType) ToStartByte() byte {
	if *f == FrameTypeResponse {
		return physical.StartByteFromMeter
	}

	return physical.StartByteToMeter
}

func (d Destination) ToPhysical() byte {
	return byte(d)
}

func DecodeFrame(data []byte) (*Frame, error) {
	if len(data) < 6 { // start, dest, command, stop, crc(2)
		return nil, ErrFrameTooShort
	}

	f := Frame{}
	switch data[0] {
	case physical.StartByteFromMeter:
		f.FrameType = FrameTypeResponse
	case physical.StartByteToMeter:
		f.FrameType = FrameTypeCommand
	default:
		f.FrameType = FrameTypeUnknown
	}

	f.DestinationAddress = Destination(data[1]) // should we check if it's a known destination?
	f.Data = physical.DeStuff(data[2 : len(data)-3])
	f.Checksum = (uint16(data[len(data)-3])<<8 + uint16(data[len(data)-2]))
	if !f.ValidChecksum() {
		return nil, ErrInvalidCRC
	}

	return &f, nil
}

func (f *Frame) CalculateCRC() uint16 {
	dataToCheckSum := append([]byte{byte(f.DestinationAddress)}, f.Data...)
	return uint16(crc.CalculateCRC(crc.XMODEM, dataToCheckSum))
}

func (f *Frame) ValidChecksum() bool {
	return f.Checksum == f.CalculateCRC()
}

func (f *Frame) EncodeFrame() []byte {
	data := []byte{}

	data = append(data, f.FrameType.ToStartByte())
	data = append(data, f.DestinationAddress.ToPhysical())
	data = append(data, physical.Stuff(f.Data)...)
	f.Checksum = f.CalculateCRC()
	data = append(data, byte(f.Checksum>>8))
	data = append(data, byte(f.Checksum&0xff))
	data = append(data, physical.StopByte)

	return data
}
