package iterations

import (
	"strings"
)

const repeatedcount = 5

func Repeat(char string) string {
	var repeated strings.Builder
	for range repeatedcount {
		repeated.WriteString(char)
	}
	return repeated.String()
}
