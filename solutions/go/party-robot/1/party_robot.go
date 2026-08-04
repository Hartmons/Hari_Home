package partyrobot
import "fmt"

func Welcome(name string) string {
	a := fmt.Sprintf("Welcome to my party, %s!", name)
    return a
}

func HappyBirthday(name string, age int) string {
	b := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return b
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	c := fmt.Sprintf("Welcome to my party, %s!\n", name)
    d := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance)
    i := fmt.Sprintf("You will be sitting next to %s.", neighbor)
    return c + d + i
}
