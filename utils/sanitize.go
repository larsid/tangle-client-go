package utils

import "regexp"

// Sanitizes all '\u' in a string.
func SanitizeString(input string) string {
    re := regexp.MustCompile(`\\u[0-9a-fA-F]+`)
    
	sanitized := re.ReplaceAllString(input, "")
    
	return sanitized
}