package raindrops

import "fmt"

func Convert(number int) string {
	pling := number % 3
	plang := number % 5
	plong := number % 7

	if pling == 0 && plang == 0 && plong == 0 {
		return "PlingPlangPlong"
	}

	if pling == 0 && plang == 0 {
		return "PlingPlang"
	}

	if pling == 0 && plong == 0 {
		return "PlingPlong"
	}

	if plang == 0 && plong == 0 {
		return "PlangPlong"
	}

	if pling == 0 {
		return "Pling"
	}

	if plang == 0 {
		return "Plang"
	}

	if plong == 0 {
		return "Plong"
	}

	return fmt.Sprint(number)
}
