package tests

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/jonkerj/gokmp/pkg/application"
	"github.com/jonkerj/gokmp/pkg/datalink"
)

func TestEncodeGetRegister(t *testing.T) {
	tests := []struct {
		name    string
		command application.GetRegister
		want    []byte
	}{
		{
			"getregister command",
			application.GetRegister{Registers: []application.Register{{Id: 0x0080}}},
			[]byte{0x80, 0x3f, 0x10, 0x01, 0x00, 0x1b, 0x7f, 0xd4, 0x08, 0x0d},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.command.ToFrame()
			got := f.EncodeFrame()

			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("EncodeFrame() = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}

func TestDecodeGetRegisterReponse(t *testing.T) {
	tests := []struct {
		name            string
		wire            []byte
		want            application.GetRegister
		wantFrameErr    bool
		wantRegisterErr bool
	}{
		{
			"getregister response",
			[]byte{0x40, 0x3f, 0x10, 0x00, 0x1b, 0x7f, 0x16, 0x04, 0x11, 0x01, 0x2a, 0xf0, 0x24, 0x63, 0x03, 0x0d},
			application.GetRegister{
				Registers: []application.Register{{
					Id:    0x0080,
					Unit:  application.Unit(0x16),
					Value: 1.9591204e+24,
				}},
			},
			false,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := datalink.DecodeFrame(tt.wire)
			if (err != nil) != tt.wantFrameErr {
				t.Errorf("DecodeFrame() error = %v, wantErr %v", err, tt.wantFrameErr)
				return
			}

			got, err := (&application.GetRegister{}).FromFrame(*f)
			if (err != nil) != tt.wantRegisterErr {
				t.Errorf("DecodeFrame() error = %v, wantErr %v", err, tt.wantRegisterErr)
				return
			}

			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("DecodeFrame() = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}
