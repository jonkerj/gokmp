package physical

func DeStuff(escaped []byte) []byte {
	var result []byte

	for i := 0; i < len(escaped); i++ {
		b := escaped[i]
		if b == EscapeByte {
			i++
			result = append(result, escaped[i]^0xff)
		} else {
			result = append(result, escaped[i])
		}
	}

	return result
}

func Stuff(raw []byte) []byte {
	var result []byte

	for _, b := range raw {
		if _, ok := escapes[b]; ok {
			result = append(result, EscapeByte)
			result = append(result, b^0xff)
		} else {
			result = append(result, b)
		}
	}

	return result
}
