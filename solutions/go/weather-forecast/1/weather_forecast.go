//Package weather contains two variables and a function.
package weather

var (
    //CurrentCondition - Weather condition variable.
	CurrentCondition string
    //CurrentLocation - A variable containing the name of a location or city.
	CurrentLocation  string
)
//Forecast The function outputs the weather conditions for a city or a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
