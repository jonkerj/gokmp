package application

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/jonkerj/gokmp/pkg/datalink"
)

func TestGetSerialNo_FromFrame(t *testing.T) {
	tests := []struct {
		name    string
		frame   datalink.Frame
		want    GetSerialNo
		wantErr bool
	}{
		{
			name: "example",
			frame: datalink.Frame{
				DestinationAddress: datalink.DestinationHeatMeter,
				FrameType:          datalink.FrameTypeResponse,
				Data:               []byte{0x02, 0x01, 0x23, 0x45, 0x67},
				Checksum:           0xe956,
			},
			want: GetSerialNo{
				Serial: 19088743,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (&GetSerialNo{}).FromFrame(tt.frame)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSerialNo.FromFrame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("GetSerialNo.FromFrame() = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}

func TestGetSerialNo_ToFrame(t *testing.T) {
	t.Run("getserial command", func(t *testing.T) {
		cmd := NewGetSerialNo()
		got := cmd.ToFrame()
		want := datalink.Frame{
			DestinationAddress: datalink.DestinationHeatMeter,
			FrameType:          datalink.FrameTypeCommand,
			Data:               []byte{CommandGetSerialNo},
			Checksum:           0x35e9,
		}

		if diff := deep.Equal(got, want); diff != nil {
			t.Errorf("GetSerialNo.ToFrame() = %v, want %v, diff %v", got, want, diff)
		}
	})
}
