
package leap


func IsLeapYear(year int) bool {
	var yb bool
    if year % 4 == 0 && year % 100 != 0 || year % 400 == 0{
        yb = true
        return yb
    }else {
        return yb
    }
    
}
