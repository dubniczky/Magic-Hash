package main

import (
    "crypto/md5"
	"strings"
    "fmt"
)

func nextSequence(s string) string {
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	maxIndex := len(alphabet) - 1

	runes := []rune(s)

	for i := len(runes) - 1; i >= 0; i-- {
		idx := strings.IndexRune(alphabet, runes[i])
		if idx == maxIndex {
			runes[i] = rune(alphabet[0])
		} else {
			runes[i] = rune(alphabet[idx+1])
			break
		}
	}

	// Check if all characters are '0', if so, add another '0' at the beginning
	if strings.Repeat(string(alphabet[0]), len(runes)) == string(runes) {
		runes = append([]rune{rune(alphabet[0])}, runes...)
	}

	return string(runes)
}

func isNumeric(s string) bool {
    for _, r := range s {
        if r < '0' || r > '9' {
            return false
        }
    }
    return true
}

func main() {
    postfix := "0"
    for {
        str := fmt.Sprintf("dubniczky-%s", postfix)
        h := md5.New()
        h.Write([]byte(str))
        md5_hash := fmt.Sprintf("%x", h.Sum(nil))
        if md5_hash[:2] == "0e" && isNumeric(md5_hash[2:]) {
            fmt.Printf("%s -> %s\n", str, md5_hash)
			break
        }

        postfix = nextSequence(postfix)
    }
}

