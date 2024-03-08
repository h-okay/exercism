package strand

var nucleotidePairs = map[byte]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var output []byte
	for i := 0; i < len(dna); i++ {
		output = append(output, nucleotidePairs[dna[i]])
	}
	return string(output)
}
