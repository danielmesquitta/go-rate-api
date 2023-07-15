package util

import "strings"

func FormatEmail(email *string) {
	*email = strings.ToLower(strings.TrimSpace(*email))
}
