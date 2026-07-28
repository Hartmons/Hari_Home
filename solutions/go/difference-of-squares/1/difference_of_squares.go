package differenceofsquares

func SquareOfSum(n int) int {
	var sosum int
    count := 0
    for i := 1; i <= n; i++{
        count += i
    }
    sosum = count * count
    return sosum
}

func SumOfSquares(n int) int {
    count := 0
    for i := 1; i <= n; i++{
        count += i * i
    }
    return count
}

func Difference(n int) int {
	a := SquareOfSum(n)
    b := SumOfSquares(n)
    return a - b
}
