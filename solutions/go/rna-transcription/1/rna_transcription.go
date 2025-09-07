package strand

var rnaTranscrib = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	dnaRunes := []rune(dna)
	for i := 0; i < len(dnaRunes); i++ {
		dnaRunes[i] = rnaTranscrib[dnaRunes[i]]
	}
	return string(dnaRunes)
}
