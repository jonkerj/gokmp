package client

import (
	"errors"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/jonkerj/gokmp/pkg/application"
	"go.bug.st/serial"
)

type (
	FakeHeatMeter struct {
		response []byte
	}
)

func (f *FakeHeatMeter) SetMode(*serial.Mode) error {
	return nil
}

func (f *FakeHeatMeter) Read(p []byte) (int, error) {
	n := copy(p, f.response)
	f.response = []byte{}
	return n, nil
}

func (f *FakeHeatMeter) Write(p []byte) (int, error) {
	switch cid := p[2]; cid {
	case application.CommandGetSerialNo:
		f.response = p
		f.response = append(f.response, []byte{0x40, 0x3f, 0x02, 0x01, 0x23, 0x45, 0x67, 0xe9, 0x56, 0x0d}...)
		return len(p), nil
	case application.CommandGetRegister:
		f.response = p
		f.response = append(f.response, []byte{0x40, 0x3f, 0x10, 0x00, 0x1b, 0x7f, 0x16, 0x04, 0x11, 0x01, 0x2a, 0xf0, 0x24, 0x63, 0x03, 0x0d}...)
		return len(p), nil
	default:
		return 0, errors.New("not implemented")
	}
}

func (f *FakeHeatMeter) Drain() error {
	return nil
}

func (f *FakeHeatMeter) ResetInputBuffer() error {
	return nil
}

func (f *FakeHeatMeter) ResetOutputBuffer() error {
	return nil
}

func (f *FakeHeatMeter) SetDTR(bool) error {
	return nil
}

func (f *FakeHeatMeter) SetRTS(bool) error {
	return nil
}

func (f *FakeHeatMeter) GetModemStatusBits() (*serial.ModemStatusBits, error) {
	return nil, nil
}

func (f *FakeHeatMeter) SetReadTimeout(time.Duration) error {
	return nil
}

func (f *FakeHeatMeter) Close() error {
	return nil
}

func (f *FakeHeatMeter) Break(time.Duration) error {
	return nil
}

func TestSerialClient(t *testing.T) {
	f := &FakeHeatMeter{}
	c := NewSerialClient(f)
	t.Run("SerialClient.GetSerialNo()", func(t *testing.T) {
		serial, err := c.GetSerialNo()
		if err != nil {
			t.Errorf("SerialClient.GetSerialNo() returned error: %v", err)
		}
		if serial != 19088743 {
			t.Errorf("SerialClient.GetSerialNo() returned invalid serial: %d", serial)
		}
	})
	t.Run("SerialClient.GetRegister()", func(t *testing.T) {
		registers, err := c.GetRegister([]uint16{0x0080})
		if err != nil {
			t.Errorf("SerialClient.GetRegister() returned error: %v", err)
		}
		if len(registers) != 1 {
			t.Errorf("SerialClient.GetRegister() expected 1 register, got %d", len(registers))
		}
		want := application.Register{
			Id:    0x0080,
			Unit:  application.Unit(0x16),
			Value: 1.9591204e+24,
		}
		if diff := deep.Equal(registers[0], want); diff != nil {
			t.Errorf("SerialClient.GetRegister() expected %v, got %v, diff %v", want, registers[0], diff)
		}
	})
}
