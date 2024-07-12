package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
	"os"
	"strconv"
	"strings"
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

func findMagicHash(algorithm func() hash.Hash, prefix string, offset int, threads int) {
	postfix := addSequence("0", offset)
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

		postfix = addSequence(postfix, threads)
	}
}

func main() {
	var args []string = os.Args[1:]
	threads, _ := strconv.Atoi(args[2])
	var prefix string = args[1]
	var hasher func() hash.Hash

	switch args[0] {
	case "md5":
		hasher = md5.New
		break
	case "sha1":
		hasher = sha1.New
		break
	}

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		//go findMagicHash(md5.New, prefix, i+1, threads)
		go findMagicHash(hasher, prefix, i+1, threads)
	}
	wg.Wait() // Called so the program does not close immediately
}
