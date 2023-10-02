package application

import (
	"github.com/jonkerj/gokmp/pkg/datalink"
)

type (
	GetType struct {
		Type    uint16
		Version uint16
	}
)

func NewGetType() GetType {
	return GetType{}
}

func (g GetType) ToFrame() datalink.Frame {
	f := basicCommandFrame(CommandGetType)
	f.Checksum = f.CalculateCRC()
	f.FrameType = datalink.FrameTypeCommand

	return f
}

func (g GetType) FromFrame(f datalink.Frame) (Command, error) {
	c := NewGetType()
	if len(f.Data) < 5 {
		return nil, ErrShortGetType
	}
	if len(f.Data) > 5 {
		return nil, ErrLongGetType
	}

	c.Type = uint16(f.Data[1])<<8 + uint16(f.Data[2])
	c.Version = uint16(f.Data[3])<<8 + uint16(f.Data[4])

	return c, nil
}
