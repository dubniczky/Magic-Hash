package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"os"
	"strconv"
	"strings"
	"sync"
)

const PasswordSequenceAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func incrementStringSequence(s string, n int) string {
	alphabet := PasswordSequenceAlphabet
	characters := len(alphabet)
	var runes []rune
	var remainder int

	// We consider an empty string the first character in the alphabet
	if s == "" {
		runes = []rune{rune(alphabet[0])}
	} else {
		runes = []rune(s)
	}

	remainder = n
	for i := len(runes) - 1; i >= 0; i-- {
		idx := strings.IndexRune(alphabet, runes[i])
		if idx == -1 {
			return "" // Invalid character in string
		}
		runes[i] = rune(alphabet[(idx+remainder)%characters])
		remainder = (idx + remainder) / characters
	}

	// If there are any remainders after the last character we start adding more at the beginning
	// We use remainder-1, because in this case we also allow the first character to count as different
	// than when it's not present, line "01" != "1". Otherwise we'd miss all sequences that start with a 0.
	for remainder > 0 {
		runes = append([]rune{rune(alphabet[(remainder-1)%characters])}, runes...)
		remainder /= characters
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

func findMagicHash(algorithm func() hash.Hash, prefix string, offset int, step int) {
	postfix := incrementStringSequence("0", offset)
	for {
		// Create hash
		str := fmt.Sprintf("%s%s", prefix, postfix)
		h := algorithm()
		h.Write([]byte(str))
		hash := fmt.Sprintf("%x", h.Sum(nil))

		// Test hash
		if hash[:2] == "0e" && isNumeric(hash[2:]) {
			fmt.Printf("%s -> %s\n", str, hash)
			os.Exit(0)
		}

		postfix = incrementStringSequence(postfix, step)
	}
}

func main() {
	var args []string = os.Args[1:]
	threads, _ := strconv.Atoi(args[2])
	var prefix string = args[1]
	var hasher func() hash.Hash

	switch args[0] {
	case "crc32":
		hasher = crc32New
		break
	case "md5":
		hasher = md5.New
		break
	case "sha1":
		hasher = sha1.New
		break
	case "sha224":
		hasher = sha256.New224
		break
	case "sha256":
		hasher = sha256.New
		break
	}

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go findMagicHash(hasher, prefix, i+1, threads)
	}
	wg.Wait()
}
