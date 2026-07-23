package lasagna


const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {

    ost := OvenTime - actualMinutesInOven
    return ost
	panic("RemainingOvenTime not implemented")
}


func PreparationTime(numberOfLayers int) int {
	slice := numberOfLayers * 2
    return slice
	panic("PreparationTime not implemented")
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    fin := PreparationTime(numberOfLayers)+ actualMinutesInOven
    return fin
	panic("ElapsedTime not implemented")
}
