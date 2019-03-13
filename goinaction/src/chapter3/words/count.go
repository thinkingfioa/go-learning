package words

import "strings"

func CountWords(text string) (count int) {
	return len(strings.Fields(text))
}
