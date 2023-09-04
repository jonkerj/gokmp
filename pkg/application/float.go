package application

import (
	"math"
)

func BinaryToFloat(b []byte) (float64, error) {
	if len(b) < 2 {
		return math.NaN(), ErrCouldNotDecodeFloat
	}
	if len(b) != int(b[0]+2) {
		return math.NaN(), ErrCouldNotDecodeFloat
	}

	mantissa := 0
	for i := 0; i < int(b[0]); i++ {
		mantissa <<= 8
		mantissa += int(b[i+2])
	}
	if b[1]&0x80 > 0 {
		mantissa *= -1
	}

	exponent := float64(b[1] & 0x3f)
	if b[1]&0x40 > 0 {
		exponent *= -1
	}

	return float64(mantissa) * math.Pow(10, exponent), nil
}

func inEpsilon(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}
