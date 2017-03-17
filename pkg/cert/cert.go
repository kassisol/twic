package cert

import (
	"strings"

	"github.com/juliengk/go-utils"
)

func GetOU(ou string) string {
	words := []string{
		"Certificate",
		"Authority",
	}

	oldou := strings.Split(ou, " ")

	if len(oldou) > 1 {
		newou := []string{}

		for _, word := range oldou {
			if !utils.StringInSlice(word, words, true) {
				newou = append(newou, word)
			}
		}

		if len(newou) > 0 {
			return strings.Join(newou, " ")
		}
	}

	return ou
}
