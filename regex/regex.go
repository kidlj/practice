package regex

import (
	"regexp"
)

// ValidNginxSize validates whether a string is a valid size in Nginx conf.
// See: http://nginx.org/en/docs/syntax.html
func ValidNginxSize(size string) bool {
	validSize := regexp.MustCompile(`^[0-9]+[kKmM]?$`)
	valid := validSize.MatchString(size)
	return valid
}

// ValidNginxPrefix validates whether a prefix string starts with `/`
func ValidNginxPrefix(prefix string) bool {
	validPrefix := regexp.MustCompile(`^/[^\s]*$`)
	valid := validPrefix.MatchString(prefix)
	return valid
}
