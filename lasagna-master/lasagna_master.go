package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2 // default preparation time per layer
	}
	return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	const noodlesPerLayer = 50 // grams of noodles per layer
	const saucePerLayer = 0.2  // liters of sauce per layer

	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += noodlesPerLayer
		}
		if layer == "sauce" {
			sauce += saucePerLayer
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsIngredientsList []string, myIngredientsList []string) {
	lastFriendIngredient := friendsIngredientsList[len(friendsIngredientsList)-1]
	myIngredientsList[len(myIngredientsList)-1] = lastFriendIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	const defaultPortions = 2
	myQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		myQuantities[i] = float64(portions) * quantities[i] / float64(defaultPortions)
	}
	return myQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
