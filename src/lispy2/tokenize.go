package lispy2
import (
	"strings"
)
func Tokenize(str string) ([]string) {
	x1 := strings.Replace(str, "(", " ( ", -1)
	x2 := strings.Replace(x1, ")", " ) ", -1)
	var ary []string
	for _, token := range strings.Split(x2, " ") {
		if token != "" {
			ary = append(ary, token)
		}
	}
	return ary
}
