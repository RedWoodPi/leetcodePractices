package __String_to_Integer

import (
	"fmt"
	"strings"
	"math"
	"strconv"
)

func main()  {
	fmt.Println("answer:"+strconv.Itoa(myAtoi("2147483648")))
}

func myAtoi(str string) int {
	tag := false
	str = strings.TrimLeft(str, " ")
	number := int64(0)
	if len(str) < 1 {
		return 0
	}
	s := str
	if str[0] == '-' || str[0] == '+' {
		s = str[1:]
		if len(s) < 1 {
			return 0
		}
	}

	for _,v := range []byte(s) {
		v = v - '0'
		if v > 9 {
			break
		}
		number = number*10 + int64(v)
		if number < math.MinInt32 || number > math.MaxInt32 {
			tag = true
			break
		}
	}
	if tag {
		number = math.MaxInt32
		if str[0] == '-'{
			number = math.MinInt32
		}
	} else {
		if str[0] == '-'{
			number = -number
		}
	}

	return int(number)
}