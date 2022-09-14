package main

import(
		"fmt"
		"strings"
)

func comma_f(s string) string{
		n := len(s)
		if n <= 3{
			return s
		}
		return comma_F(s[: n - 3]) + "," + s[n - 3 :]
}

func comma_l(s string) string{
		n := len(s)
		if n <= 3{
				return s
		}
		return s[0 : 3] + ", " + comma_l(s[3 :])
}

func main(){
	s := "121314252.42342442"
	c := strings.LastIndex(s, '.')
	
	fmt.Println(comma_f(s[: c]) + "." + comma_l(s[c + 1:]))
}
