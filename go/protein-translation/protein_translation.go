package protein

import (
	"errors"
)

var ErrStop error = errors.New("ErrStop")
var ErrInvalidBase error = errors.New("ErrInvalidBase")

var CODONS = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGC": "Cysteine",
	"UGU": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	proteins := make([]string, 0, len(rna)/3)
	codon := make([]rune, 0, 3)
	for i, nucleotid := range rna {
		codon = append(codon, nucleotid)
		if (i+1)%3 == 0 {
			protein, err := FromCodon(string(codon))
			if err == ErrStop {
				break
			}

			proteins = append(proteins, protein)
			codon = make([]rune, 0, 3)
		}
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := CODONS[codon]
	if !ok {
		return "", ErrInvalidBase
	}

	if protein == "STOP" {
		return "", ErrStop
	}

	return protein, nil
}
