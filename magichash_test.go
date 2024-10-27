package main

import "testing"

type stringSequenceTestPair struct {
	s      string
	n      int
	result string
}

var passwordSequenceLastItem = len(PasswordSequenceAlphabet)
var expectedStringSequences = []stringSequenceTestPair{
	{"", 0, "0"},
	{"0", 0, "0"},
	{"0", 1, "1"},
	{"0", 15, "f"},
	{"Z", 1, "00"},
	{"Z", 3, "02"},
	{"X3", 128, "Z7"},
	{"ZZ", 62, "00Z"},
	{"A5j", 0, "A5j"},
	{"A5j", 1, "A5k"},
	{"A5j", 2, "A5l"},
	{"A5j", 3, "A5m"},
}

func TestPasswordSequencing(t *testing.T) {
	for i, s := range expectedStringSequences {
		res := incrementStringSequence(s.s, s.n)
		if res != s.result {
			t.Fatalf("incrementStringSequence(\"%s\", %d) = \"%s\", expected: \"%s\", test: %d",
				s.s, s.n, res, s.result, i+1)
		}
	}
}
