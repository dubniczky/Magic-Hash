package main

import (
	"log"
	"strings"
	"time"
)

const PasswordSequenceAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func incrementStringSequence_iteration(s string, n int) string {
	alphabet := PasswordSequenceAlphabet
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

func incrementStringSequence_mapping(s string, n int) string {
	alphabet := PasswordSequenceAlphabet
	base := len(alphabet)

	// Create a map for O(1) character to index lookup
	runeToIndex := make(map[rune]int)
	for i, r := range alphabet {
		runeToIndex[r] = i
	}

	// Convert the string to a slice of indices
	digits := make([]int, len(s))
	for i, r := range s {
		digits[i] = runeToIndex[r]
	}

	// Add n to the number represented by digits
	carry := n
	for i := len(digits) - 1; i >= 0 && carry > 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % base
		carry = sum / base
	}

	// Handle any remaining carry
	for carry > 0 {
		digits = append([]int{carry % base}, digits...)
		carry /= base
	}

	// Convert the digits back to the string
	result := make([]rune, len(digits))
	for i, d := range digits {
		result[i] = rune(alphabet[d])
	}

	return string(result)
}

func incrementStringSequence_remainder(s string, n int) string {
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

func main() {
	runs := 100000
	input := "0"

	start := time.Now()
	for i := 0; i < runs; i++ {
		input = incrementStringSequence_iteration(input, 12)
	}
	log.Printf("Iterator: %s (%s)", time.Since(start), input)

	start = time.Now()
	input = "0"
	for i := 0; i < runs; i++ {
		input = incrementStringSequence_mapping(input, 12)
	}
	log.Printf("Mapping: %s, (%s)", time.Since(start), input)

	start = time.Now()
	input = "0"
	for i := 0; i < runs; i++ {
		input = incrementStringSequence_remainder(input, 12)
	}
	log.Printf("Remainder: %s, (%s)", time.Since(start), input)
}
