package main

import (
    "crypto/md5"
	"strings"
    "fmt"
	"os"
	"sync"
)

func addSequence(s string, n int) string {
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	maxIndex := len(alphabet) - 1
	runes := []rune(s)

	for n > 0 {
		for i := len(runes) - 1; i >= 0; i-- {
			idx := strings.IndexRune(alphabet, runes[i])
			if idx == maxIndex {
				runes[i] = rune(alphabet[0])
				if i == 0 {
					runes = append([]rune{rune(alphabet[0])}, runes...)
				}
			} else {
				runes[i] = rune(alphabet[idx+1])
				break
			}
		}
		n--
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

func findMagicHash(offset int) {
	postfix := addSequence("0", offset)

    for {
        str := fmt.Sprintf("dubniczky-%s", postfix)
        h := md5.New()
        h.Write([]byte(str))
        md5_hash := fmt.Sprintf("%x", h.Sum(nil))
        if md5_hash[:2] == "0e" && isNumeric(md5_hash[2:]) {
            fmt.Printf("%s -> %s\n", str, md5_hash)
			os.Exit(0)
        }

        postfix = addSequence(postfix, offset)
    }
}

func main() {
	var wg sync.WaitGroup
    for i := 0; i < 8; i++ {
		wg.Add(1)
		go findMagicHash(i)
	}

	wg.Wait()
}
