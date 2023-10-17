package datalink

import (
	"testing"

	"github.com/go-test/deep"
)

func TestDecodeFrame(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *Frame
		wantErr bool
	}{
		{"heat meter command", []byte{0x80, 0x3f, 0x12, 0x13, 0xae, 0xd7, 0x0d}, &Frame{DestinationHeatMeter, FrameTypeCommand, []byte{0x12, 0x13}, 0xaed7}, false},
		{"heat meter bad crc", []byte{0x3f, 0x12, 0x13, 0x11, 0x22}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeFrame(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeFrame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("DecodeFrame() = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}

func TestFrame_CalculateCRC(t *testing.T) {
	type fields struct {
		DestinationAddress Destination
		FrameType          FrameType
		Data               []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		{"get serial cmd", fields{DestinationHeatMeter, FrameTypeCommand, []byte{0x02}}, 0x35e9},
		{"get serial res", fields{DestinationHeatMeter, FrameTypeResponse, []byte{0x02, 0x01, 0x23, 0x45, 0x67}}, 0xe956},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Frame{
				DestinationAddress: tt.fields.DestinationAddress,
				FrameType:          tt.fields.FrameType,
				Data:               tt.fields.Data,
			}
			if got := f.CalculateCRC(); got != tt.want {
				t.Errorf("Frame.CalculateCRC() = %04x, want %04x", got, tt.want)
			}
		})
	}
}

func TestFrame_EncodeFrame(t *testing.T) {
	type fields struct {
		DestinationAddress Destination
		FrameType          FrameType
		Data               []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"get serial cmd", fields{DestinationHeatMeter, FrameTypeCommand, []byte{0x02}}, []byte{0x80, 0x3f, 0x02, 0x35, 0xe9, 0x0d}},
		{"get serial res", fields{DestinationHeatMeter, FrameTypeResponse, []byte{0x02, 0x01, 0x23, 0x45, 0x67}}, []byte{0x40, 0x3f, 0x02, 0x01, 0x23, 0x45, 0x67, 0xe9, 0x56, 0x0d}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Frame{
				DestinationAddress: tt.fields.DestinationAddress,
				FrameType:          tt.fields.FrameType,
				Data:               tt.fields.Data,
			}
			got := f.EncodeFrame()
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("Frame.EncodeFrame() = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}
