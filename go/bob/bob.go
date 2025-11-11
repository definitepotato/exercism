// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

const (
	Whatever = "Whatever."
	Sure     = "Sure."
	Chill    = "Whoa, chill out!"
	CalmDown = "Calm down, I know what I'm doing!"
	Fine     = "Fine. Be that way!"
)

type Remark struct {
	phrase string
}

func (r Remark) isSilence() bool {
	return r.phrase == ""
}

func (r Remark) isShouting() bool {
	hasLetters := strings.IndexFunc(r.phrase, unicode.IsLetter) >= 0
	isUpcased := strings.ToUpper(r.phrase) == r.phrase

	return hasLetters && isUpcased
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.phrase, "?")
}

func (r Remark) isExasperated() bool {
	return r.isShouting() && r.isQuestion()
}

func newRemark(remark string) Remark {
	return Remark{
		strings.TrimSpace(remark),
	}
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := newRemark(remark)
	switch {
	case r.isSilence():
		return Fine
	case r.isExasperated():
		return CalmDown
	case r.isShouting():
		return Chill
	case r.isQuestion():
		return Sure
	default:
		return Whatever
	}
}
