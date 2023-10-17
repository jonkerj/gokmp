package application

import (
	"math"
	"testing"
)

func TestBinaryToFloat(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    float64
		wantErr bool
	}{
		{"one", []byte{0x04, 0xc2, 0x00, 0x00, 0x30, 0x39}, -123.45, false},
		{"two", []byte{0x04, 0x03, 0x05, 0x39, 0x4f, 0xb1}, 8.7642033e+10, false}, // documentation shows 8.7654321e+10, close enough
		{"three", []byte{0x01, 0x03, 0xff}, 255e+3, false},
		{"four", []byte{0x01, 0x03}, math.NaN(), true},
		{"five", []byte{0x00, 0x01}, 0, false},
		{"six", []byte{0x00}, math.NaN(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryToFloat(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryToFloat(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !(math.IsNaN(got) && math.IsNaN(tt.want)) && (got != tt.want) {
				t.Errorf("BinaryToFloat(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
