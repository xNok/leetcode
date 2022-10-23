package lettercombinations

func letterCombinations(digits string) []string {
	result := []string{}

	if len(digits) == 0 {
		return result
	}

	sDigits := []byte(string(digits))

	for _, d := range sDigits {
		combinations := dummyDigitConvert(d)

		if len(result) == 0 {
			result = append(result, combinations...)
			continue
		}

		// r = {a,b,c} + c = {j, k, l} => {aj, bj, cj, ...}
		for i, r := range result {
			// update in place
			for j, c := range combinations {
				if j == 0 {
					result[i] = r + c
				} else {
					result = append(result, r+c)
				}
			}
		}
	}

	return result
}

func digitConvert(b byte) []string {
	// 2 -> 50 => a = 97, b = 98, c = 99
	// 3 -> 51 => d = 100
	// 5 -> 53 => j = 106, k = 107, l = 108

	// how many time 3
	mod := (b - 50) * 3

	// 3 -> 1* 3

	// 3 -> 97 + 3 = 100
	return []string{string(97 + mod), string(97 + mod + 1), string(97 + mod + 2)}

}

func dummyDigitConvert(b byte) []string {
	switch b {
	case 50:
		return []string{"a", "b", "c"}
	case 51:
		return []string{"d", "e", "f"}
	case 52:
		return []string{"g", "h", "i"}
	case 53:
		return []string{"j", "k", "l"}
	case 54:
		return []string{"m", "n", "o"}
	case 55:
		return []string{"p", "q", "r", "s"}
	case 56:
		return []string{"t", "u", "v"}
	case 57:
		return []string{"w", "x", "y", "z"}
	}

	return []string{}
}

func KeyboardLetters() map[string][]string {
	keyboard := map[string][]string{"1": {},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	return keyboard
}
