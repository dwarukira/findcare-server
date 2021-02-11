package txt

import (
	"regexp"
	"strconv"
)

var UnknownCountryCode = "zz"
var CountryWordsRegexp = regexp.MustCompile("[\\p{L}]{2,}")

// Int returns a string as int or 0 if it can not be converted.
func Int(s string) int {
	if s == "" {
		return 0
	}

	result, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0
	}

	return int(result)
}

// IsUInt returns true if a string only contains an unsigned integer.
func IsUInt(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r < 48 || r > 57 {
			return false
		}
	}

	return true
}
