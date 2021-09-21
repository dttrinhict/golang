package pattern

import "strings"

func TrimSlash(str string) string  {
	return strings.Trim(strings.Trim(str, "/"), "\\")

}
