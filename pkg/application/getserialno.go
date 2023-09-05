package application

import (
	"github.com/jonkerj/gokmp/pkg/datalink"
)

type (
	GetSerialNo struct {
		Serial uint32
	}
)

func NewGetSerialNo() GetSerialNo {
	return GetSerialNo{}
}

func (g GetSerialNo) ToFrame() datalink.Frame {
	f := basicCommandFrame(CommandGetSerialNo)
	f.Checksum = f.CalculateCRC()
	f.FrameType = datalink.FrameTypeCommand

	return f
}

func (g GetSerialNo) FromFrame(f datalink.Frame) (Command, error) {
	c := NewGetSerialNo()
	if len(f.Data) < 5 {
		return nil, ErrShortSerial
	}
	if len(f.Data) > 5 {
		return nil, ErrLondSerial
	}

	c.Serial = uint32(f.Data[1])<<24 + uint32(f.Data[2])<<16 + uint32(f.Data[3])<<8 + uint32(f.Data[4])

	return c, nil
}
