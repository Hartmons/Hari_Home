package luhn
import "strings"
func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")

    if len(id) <= 1{
        return false
    }

    for _, ch := range id {
        if ch < '0' || ch > '9' {
            return false
        }
    }
    
	sum := 0
    
    for i := len(id) - 1; i >= 0; i--{
        digit := int(id[i] - '0')
        if (len(id)-1-i) % 2 == 1 {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
    }
    return sum % 10 == 0
}
