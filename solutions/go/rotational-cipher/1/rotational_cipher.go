package rotationalcipher

const alphabetLen = 26

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey == 0 {
		return plain
	}

	input := []byte(plain)
	output := make([]byte, len(input))

	for i, n := range input {
		switch {
		case 'a' <= n && n <= 'z':
			output[i] = 'a' + (n-'a'+byte(shiftKey))%alphabetLen
		case 'A' <= n && n <= 'Z':
			output[i] = 'A' + (n-'A'+byte(shiftKey))%alphabetLen
		default:
			output[i] = n
		}
	}
	return string(output)
}
