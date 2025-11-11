package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	n := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", n)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		res, _ := strconv.Atoi(fnb.Value())
		return res
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fnb.(type) {
	case FancyNumber:
		n, _ := strconv.Atoi(fnb.Value())
		res := float64(n)
		return fmt.Sprintf("This is a fancy box containing the number %.1f", res)
	default:
		return "This is a fancy box containing the number 0.0"
	}
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		return DescribeNumber(float64(i.(int)))
	case float64:
		return DescribeNumber(i.(float64))
	case NumberBox:
		return DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		return DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		return "Return to sender"
	}
}
