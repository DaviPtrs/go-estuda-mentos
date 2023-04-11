package protein

import "errors"

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid condon")

func FromRNA(rna string) ([]string, error) {
	var output []string
	length := len(rna)
	if length%3 != 0 {
		return nil, ErrInvalidBase
	}

	for i := 0; i < length; i += 3 {
		condon := rna[i : i+3]
		condonProtein, err := FromCodon(condon)
		if err == ErrInvalidBase {
			return nil, err
		}
		if err == ErrStop {
			break
		}
		output = append(output, condonProtein)
	}

	return output, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
