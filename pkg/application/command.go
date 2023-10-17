package application

import "github.com/jonkerj/gokmp/pkg/datalink"

type (
	Command struct {
		CID                byte
		DestinationAddress datalink.Destination
	}
)

func (c *Command) ToBasicFrame() *datalink.Frame {
	data := []byte{c.CID}

	return &datalink.Frame{
		DestinationAddress: c.DestinationAddress,
		Data:               data,
		Checksum:           0x00,
	}
}
