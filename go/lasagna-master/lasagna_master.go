package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepPerLayer int) int {
	if avgPrepPerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgPrepPerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauces := 0.0
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauces += 0.2
		}
	}
	return noodles, sauces

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, ownRecipe []string) {
	secret := friendsList[len(friendsList)-1]
	ownRecipe[len(ownRecipe)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, numberOfPortions int) []float64 {
	scaled := make([]float64, len(amounts))
	scale := float64(numberOfPortions) / 2.0
	for i := range amounts {
		scaled[i] = amounts[i] * scale
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
