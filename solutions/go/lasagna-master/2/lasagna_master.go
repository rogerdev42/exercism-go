package lasagna

const (
    defaultPrepTime = 2
    noodlesPerLayer = 50
    saucePerLayer = 0.2
    defaultPortionsAmount = 2
)

func PreparationTime(layers []string, time int) int {
    if time > 0 {
        return len(layers) * time
    }
    return len(layers) * defaultPrepTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _, l := range layers {
		switch l {
		case "noodles":
			noodles += noodlesPerLayer
		case "sauce":
			sauce += saucePerLayer
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
    if len(friendsList) == 0 || len(myList) == 0 {
		return
	}
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portionsAmount int) []float64 {
    scaledQuantities := make([]float64, len(quantities))
    scalingFactor := float64(portionsAmount) / float64(defaultPortionsAmount)
    for i := 0; i < len(quantities); i++ {
        scaledQuantities[i] = quantities[i] * scalingFactor
    }
    return scaledQuantities
}
