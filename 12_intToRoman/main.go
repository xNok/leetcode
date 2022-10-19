package inttoroman

import "strconv"

func intToRoman(num int) string {
	result := ""
	num_string := strconv.Itoa(num)

	l := len(num_string)

	if len(num_string) > 3 {
		result += convert(num_string[l-4], "M", "M", "M")
	}
	if len(num_string) > 2 {
		result += convert(num_string[l-3], "C", "D", "M")
	}
	if len(num_string) > 1 {
		result += convert(num_string[l-2], "X", "L", "C")
	}
	result += convert(num_string[l-1], "I", "V", "X")

	return result

}

func convert(b byte, one, five, ten string) string {
	switch b {
	case byte(49):
		return one
	case byte(50):
		return one + one
	case byte(51):
		return one + one + one
	case byte(52):
		return one + five
	case byte(53):
		return five
	case byte(54):
		return five + one
	case byte(55):
		return five + one + one
	case byte(56):
		return five + one + one + one
	case byte(57):
		return one + ten
	}

	return ""

}
