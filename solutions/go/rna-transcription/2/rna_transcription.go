// Package strand provides functions for DNA to RNA transcription.
package strand

// rnaTranscrib maps DNA nucleotides to their RNA complements.
var rnaTranscrib = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns the RNA complement of the given DNA strand.
func ToRNA(dna string) string {
	dnaRunes := []rune(dna)
	for i := 0; i < len(dnaRunes); i++ {
		dnaRunes[i] = rnaTranscrib[dnaRunes[i]]
	}
	return string(dnaRunes)
}
