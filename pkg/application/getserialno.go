package application

import (
	"github.com/jonkerj/gokmp/pkg/datalink"
)

type (
	GetSerialNoCommand struct {
		Command
		Serial uint32
	}
)

func NewGetSerialNoCommand(destination datalink.Destination) *GetSerialNoCommand {
	return &GetSerialNoCommand{
		Command: Command{
			CID:                CommandGetSerialNo,
			DestinationAddress: destination,
		},
	}
}

func (g *GetSerialNoCommand) ToFrame() *datalink.Frame {
	f := g.ToBasicFrame()
	f.Checksum = f.CalculateCRC()
	f.FrameType = datalink.FrameTypeCommand

	return f
}

func GetSerialNoCommandFromFrame(f *datalink.Frame) (*GetSerialNoCommand, error) {
	c := NewGetSerialNoCommand(f.DestinationAddress)
	if len(f.Data) < 5 {
		return nil, ErrShortSerial
	}
	if len(f.Data) > 5 {
		return nil, ErrLondSerial
	}

	c.Serial = uint32(f.Data[1])<<24 + uint32(f.Data[2])<<16 + uint32(f.Data[3])<<8 + uint32(f.Data[4])

	return c, nil
}
