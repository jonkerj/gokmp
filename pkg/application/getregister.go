package application

import (
	"github.com/jonkerj/gokmp/pkg/datalink"
)

type (
	GetRegisterCommand struct {
		Command
		Registers []*Register
	}
)

func NewGetRegisterCommand(destination datalink.Destination, registerIds []uint16) *GetRegisterCommand {
	registers := []*Register{}
	for _, id := range registerIds {
		registers = append(registers, &Register{Id: id})
	}

	return &GetRegisterCommand{
		Command: Command{
			CID:                CommandGetRegister,
			DestinationAddress: destination,
		},
		Registers: registers,
	}
}

func (g *GetRegisterCommand) ToFrame() *datalink.Frame {
	f := g.ToBasicFrame()
	f.FrameType = datalink.FrameTypeCommand

	f.Data = append(f.Data, byte(len(g.Registers)))

	for _, r := range g.Registers {
		f.Data = append(f.Data, []byte{byte(r.Id >> 8), byte(r.Id & 0xff)}...)
	}

	f.Checksum = f.CalculateCRC()

	return f
}

func GetRegisterCommandFromFrame(f *datalink.Frame) (*GetRegisterCommand, error) {
	c := NewGetRegisterCommand(f.DestinationAddress, []uint16{})
	remaining := f.Data[1:] // strip CID
	for len(remaining) > 0 {
		register, len, err := RegisterFromBytes(remaining)
		if err != nil {
			return nil, err
		}

		c.Registers = append(c.Registers, register)
		remaining = remaining[len:]
	}

	return c, nil
}
