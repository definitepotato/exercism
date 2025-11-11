package lasagna

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}

	return len(layers) * prepTime
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "sauce" {
			sauce += 0.2
		}

		if layer == "noodles" {
			noodles += 50
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]

	return myList
}

func ScaleRecipe(list []float64, portions int) []float64 {
	if len(list) == 0 {
		return []float64{}
	}

	quantities := []float64{}
	for _, item := range list {
		quantities = append(quantities, item*float64(portions)/2.00)
	}

	return quantities
}
