package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0.00 {
		return 3.213
	}

	if balance >= 0.00 && balance < 1000.00 {
		return 0.5
	}

	if balance >= 1000.00 && balance < 5000.00 {
		return 1.621
	}

	if balance >= 5000.00 {
		return 2.475
	}

	return 0.00
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	updatedBalance := balance

	for updatedBalance < targetBalance {
		updatedBalance += Interest(updatedBalance)
		years += 1
	}

	return years
}
