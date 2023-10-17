package application

import "github.com/jonkerj/gokmp/pkg/datalink"

type (
	Command interface {
		ToFrame() datalink.Frame
		FromFrame(datalink.Frame) (Command, error)
	}
)

func basicCommandFrame(CID byte) datalink.Frame {
	return datalink.Frame{
		DestinationAddress: datalink.DestinationHeatMeter,
		Data:               []byte{CID},
		FrameType:          datalink.FrameTypeCommand,
		Checksum:           0x00,
	}
}
