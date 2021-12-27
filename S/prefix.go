package S

import "strings"

//Except will be priority.
func HasPrefix(str string, prefix []string, except []string) bool {
	for _, v := range except {
		if strings.HasPrefix(str, v) {
			return false
		}
	}
	for _, v := range prefix {
		if strings.HasPrefix(str, v) {
			return true
		}
	}
	return false
}
