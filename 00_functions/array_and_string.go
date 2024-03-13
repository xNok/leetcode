package funcions

// Strings, bytes, runes and characters in Go
// https://go.dev/blog/strings

// reverseString
// we need to use len - 1 for right pointer
// nice way to swap varaibles in go in one lign
func reverseString(s []byte) {
	i, j := 0, len(s)-1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
