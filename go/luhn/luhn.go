package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func reverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func Valid(id string) bool {
	cleanId := strings.ReplaceAll(id, " ", "")
	if len(cleanId) <= 1 {
		return false
	}

	cleanId = reverseString(cleanId)

	// process length of `id` starting from the right
	var total int
	for idx := len(cleanId) - 1; idx >= 0; idx-- {
		char := fmt.Sprintf("%c", cleanId[idx])
		num, err := strconv.Atoi(char)
		if err != nil {
			return false
		}

		// process every 2nd digit
		if idx%2 != 0 {
			// double the number
			num *= 2

			// if doubling the number results in a number greater
			// than 9 then subtract 9 from the product.
			if num > 9 {
				num -= 9
			}
		}

		total += num
	}

	return total%10 == 0
}
