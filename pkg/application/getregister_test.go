package application

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/jonkerj/gokmp/pkg/datalink"
)

func TestGetRegisterCommand_ToFrame(t *testing.T) {
	tests := []struct {
		name        string
		destination datalink.Destination
		registers   []uint16
		want        *datalink.Frame
	}{
		{"test1", 0x3f, []uint16{0x0080}, &datalink.Frame{DestinationAddress: 0x3f, FrameType: datalink.FrameTypeCommand, Data: []byte{0x10, 0x01, 0x00, 0x80}, Checksum: 0xd408}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGetRegisterCommand(tt.destination, tt.registers)
			if got := g.ToFrame(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRegisterCommand.ToFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRegisterCommandFromFrame(t *testing.T) {
	tests := []struct {
		name    string
		frame   datalink.Frame
		want    *GetRegisterCommand
		wantErr bool
	}{
		{
			name: "register0x80",
			frame: datalink.Frame{
				DestinationAddress: datalink.DestinationHeatMeter,
				FrameType:          datalink.FrameTypeResponse,
				Data:               []byte{0x10, 0x00, 0x80, 0x16, 0x01, 0x03, 0xff},
				Checksum:           0xca18,
			},
			want: &GetRegisterCommand{
				Command: Command{CID: CommandGetRegister, DestinationAddress: datalink.DestinationHeatMeter},
				Registers: []*Register{{
					Id:    0x0080,
					Unit:  Unit(0x16),
					Value: 255e+03,
				}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRegisterCommandFromFrame(&tt.frame)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegisterCommandFromFrame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("GetRegisterCommandFromFrame() = %v, want %v, diff = %v", got, tt.want, diff)
			}
		})
	}
}
