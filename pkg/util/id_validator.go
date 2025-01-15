package util

import "regexp"

var idRegexp = regexp.MustCompile(`^\d+$`)

func ValidateID(id string) bool {
	return idRegexp.MatchString(id)
}
