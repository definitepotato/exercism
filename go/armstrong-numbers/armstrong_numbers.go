package armstrong

import (
	"fmt"
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	if n == 0 {
		return true
	}

	var sum float64
	digits, numStr := len(strconv.Itoa(n)), strconv.Itoa(n)

	fmt.Printf("%d %s\n", digits, numStr)

	for _, v := range numStr {
		i := v - '0'
		sum += math.Pow(float64(i), float64(digits))
	}

	return int(sum) == n
}
