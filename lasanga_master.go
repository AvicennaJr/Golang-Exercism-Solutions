package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(slice []string, ave int) int {
    total := len(slice)
    if ave == 0 {
        return 2 * total
    } else {
    	return ave * total
    }
}

// TODO: define the 'Quantities()' function

func Quantities(slice []string) (int, float64) {
    noodles := 0
    sauce := float64(0)

    for _, v := range slice {
        if v == "noodles" {
            noodles += 50
        } else if v == "sauce" {
        	sauce += 0.2
        }
    }
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendsList, myList []string) {
    friendIndex := len(friendsList) - 1
    myIndex := len(myList) - 1
    myList[myIndex] = friendsList[friendIndex]
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(slice []float64, portions int) []float64 {
    scaledRecipe := make([]float64, len(slice))
    for i, v := range slice {
        scaledRecipe[i] = v * float64(portions)/2.0
    }
	return scaledRecipe
}
