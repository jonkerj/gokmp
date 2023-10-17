package application

import (
	"fmt"
)

type (
	Unit byte

	Register struct {
		Id    uint16
		Unit  Unit
		Value float64
	}
)

func (u Unit) String() string {
	switch byte(u) {
	case 0x01:
		return "Wh"
	case 0x02:
		return "kWh"
	case 0x03:
		return "MWh"
	case 0x08:
		return "Gj"
	case 0x0c:
		return "Gcal"
	case 0x16:
		return "kW"
	case 0x17:
		return "MW"
	case 0x25:
		return "°C"
	case 0x26:
		return "K"
	case 0x27:
		return "l"
	case 0x28:
		return "m³"
	case 0x29:
		return "l/h"
	case 0x2a:
		return "m³/h"
	case 0x2b:
		return "m³xC"
	case 0x2c:
		return "ton"
	case 0x2d:
		return "ton/h"
	case 0x2e:
		return "h"
	case 0x2f:
		return "clock"
	case 0x30:
		return "date1"
	case 0x32:
		return "date2"
	case 0x33:
		return "number"
	case 0x34:
		return "bar"
	default:
		return "unknown"
	}
}

func RegisterFromBytes(data []byte) (*Register, byte, error) {
	if len(data) < 5 {
		return nil, 0, ErrShortRegister
	}
	len := data[3]
	value, err := BinaryToFloat(data[3 : len+5])
	if err != nil {
		return nil, 0, err
	}

	r := &Register{
		Id:    uint16(data[0])<<8 + uint16(data[1]),
		Unit:  Unit(data[2]),
		Value: value,
	}

	return r, len + 5, nil
}

func (r *Register) String() string {
	return fmt.Sprintf("%04x: %v %s", r.Id, r.Value, r.Unit)
}
