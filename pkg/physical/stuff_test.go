package physical

import (
	"testing"

	"github.com/go-test/deep"
)

func TestDeStuff(t *testing.T) {
	tests := []struct {
		name      string
		escaped   []byte
		unescaped []byte
	}{
		{"start byte to meter", []byte{0x80, 0x1b, 0x7f, 0x0d}, []byte{0x80, 0x80, 0x0d}},
		{"start byte from meter", []byte{0x80, 0x1b, 0xbf, 0x0d}, []byte{0x80, 0x40, 0x0d}},
		{"stop byte", []byte{0x80, 0x1b, 0xf2, 0x0d}, []byte{0x80, 0x0d, 0x0d}},
		{"escape byte", []byte{0x80, 0x1b, 0xe4, 0x0d}, []byte{0x80, 0x1b, 0x0d}},
		{"ack", []byte{0x80, 0x1b, 0xf9, 0x0d}, []byte{0x80, 0x06, 0x0d}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeStuff(tt.escaped)
			if diff := deep.Equal(got, tt.unescaped); diff != nil {
				t.Errorf("UnEscape(%v) = %v, want %v diff %v", tt.escaped, got, tt.unescaped, diff)
			}
		})
	}
}

func TestStuff(t *testing.T) {
	tests := []struct {
		name      string
		unescaped []byte
		escaped   []byte
	}{
		{"start byte to meter", []byte{0x12, 0x80, 0x34}, []byte{0x12, 0x1b, 0x7f, 0x34}},
		{"start byte from meter", []byte{0x12, 0x40, 0x34}, []byte{0x12, 0x1b, 0xbf, 0x34}},
		{"stop byte", []byte{0x12, 0x0d, 0x34}, []byte{0x12, 0x1b, 0xf2, 0x34}},
		{"escape byte", []byte{0x12, 0x1b, 0x34}, []byte{0x12, 0x1b, 0xe4, 0x34}},
		{"ack", []byte{0x12, 0x06, 0x34}, []byte{0x12, 0x1b, 0xf9, 0x34}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Stuff(tt.unescaped)
			if diff := deep.Equal(got, tt.escaped); diff != nil {
				t.Errorf("Escape(%v) = %v, want %v diff %v", tt.unescaped, got, tt.escaped, diff)
			}
		})
	}
}
