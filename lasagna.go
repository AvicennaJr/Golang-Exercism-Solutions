package main

// TODO: define the 'OvenTime' constant

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	return 40 - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {

	return (numberOfLayers * 2) + actualMinutesInOven

}
