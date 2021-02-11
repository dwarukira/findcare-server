package txt

// ContainsASCIILetters reports if the string only contains ascii chars without whitespace, numbers, and punctuation marks.
func ContainsASCIILetters(s string) bool {
	for _, r := range s {
		if (r < 65 || r > 90) && (r < 97 || r > 122) {
			return false
		}
	}

	return true
}
