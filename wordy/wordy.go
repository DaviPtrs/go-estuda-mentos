package wordy

import (
	"strconv"
	"strings"
)

func calcByOperator(n1 int, n2 int, operator string) (int, bool) {
	err := false
	var result int
	switch operator {
	case "plus":
		result = n1 + n2
	case "minus":
		result = n1 - n2
	case "multiplied":
		result = n1 * n2
	case "divided":
		result = n1 / n2
	default:
		err = true
	}
	return result, err
}

func cleanQuestion(question string) (string, bool) {
	question = strings.ToLower(question)
	question, found := strings.CutPrefix(question, "what is ")
	if !found {
		return "", false
	}
	question, found = strings.CutSuffix(question, "?")
	if !found {
		return "", false
	}
	question = strings.ReplaceAll(question, "by", "")
	return question, true
}

func calculate(elems []string) (int, bool) {
	ok := false
	numbers := []int{}
	operators := []string{}
	waitOp := false // true when I'm waiting for a operator // otherwise I'm waiting for a number

	for _, elem := range elems {
		num, invalid := strconv.Atoi(elem)
		if invalid != nil {
			if !waitOp {
				return 0, ok
			}
			operators = append(operators, elem)
			waitOp = false
		} else {
			if waitOp {
				return 0, ok
			}
			numbers = append(numbers, num)
			waitOp = true
		}
		if len(numbers) == 2 && len(operators) == 1 {
			result, invalidCalc := calcByOperator(numbers[0], numbers[1], operators[0])
			if invalidCalc {
				return 0, ok
			}
			numbers = []int{result}
			operators = []string{}
		}
	}

	if len(numbers) == 1 && len(operators) == 0 {
		ok = true
		return numbers[0], ok
	}

	return 0, ok
}

func Answer(question string) (int, bool) {
	question, ok := cleanQuestion(question)
	if !ok || len(question) == 0 {
		return 0, ok
	}
	splited := strings.Fields(question)
	return calculate(splited)

}
