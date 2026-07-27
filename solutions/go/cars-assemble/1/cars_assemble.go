package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    x := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(x) / 60
}


func CalculateCost(carsCount int) uint {
	ost := carsCount / 10 //3
    ost2 := carsCount % 10 //7
    return uint(ost) * 95000 + uint(ost2) * 10000
}
