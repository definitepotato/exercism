package house

import (
	"fmt"
	"strings"
)

var lyrics = [12]string{
	"the house that Jack built.",
	"the malt\nthat lay in ",
	"the rat\nthat ate ",
	"the cat\nthat killed ",
	"the dog\nthat worried ",
	"the cow with the crumpled horn\nthat tossed ",
	"the maiden all forlorn\nthat milked ",
	"the man all tattered and torn\nthat kissed ",
	"the priest all shaven and shorn\nthat married ",
	"the rooster that crowed in the morn\nthat woke ",
	"the farmer sowing his corn\nthat kept ",
	"the horse and the hound and the horn\nthat belonged to ",
}

var verses = []string{}

func makeVerse(i int) string {
	if i == 0 {
		return lyrics[i]
	}

	return fmt.Sprintf("%s%s", lyrics[i], makeVerse(i-1))
}

func Song() string {
	for i := 1; i < len(lyrics)+1; i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n\n")
}

func Verse(i int) string {
	return fmt.Sprintf("This is %s", makeVerse(i-1))
}
