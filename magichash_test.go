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
