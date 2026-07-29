package darts
//За пределы = 0 очков
//Внешний ркуг = 1 очко, радиус = 10 ед
//средний круг = 5 очков, радиус 5 ед
//внутренний круг = 10 очков, радиус 1 ед
func Score(x, y float64) int {
    rad := x * x + y * y
	switch{
        case rad <= 1: 
        return 10
        case rad <= 25:
        return 5
        case rad <= 100:
        return 1
        default:
        return 0
    }
}
