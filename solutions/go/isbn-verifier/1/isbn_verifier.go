package isbnverifier
import "strings"
func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")

    if len(isbn) != 10{
        return false
    }
    	
    sas := 10
    azaz := 0
    for i := 0; i < 10; i++ {
        ch := isbn[i]
        var val int
        
        if ch >= '0' && ch <= '9'{
            val = int(ch - '0')
        } else if ch == 'X' && i == 9 {
            val = 10
        } else {
            return false
        }
        azaz += val * sas
        sas--
    }

    return azaz % 11 == 0
}
