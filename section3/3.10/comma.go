package main

import(
	"fmt"
	"bytes"
)

func comma(s []byte) string{
	var buf bytes.Buffer
	n := len(s)
	m := n % 3
	var k int

	if m == 0{
		k = n / 3 - 1
	}else{
		k = n / 3
	}

	for  _, c := range s{
		m --
		buf.WriteByte(c)	
		if m == 0{
			if k == 0 {
				break
			}else{
				buf.WriteByte(',')
				k --
				m = 3
			}
		}
	}
	return buf.String()
}

func main(){
	s := "12124342342"
	b := []byte(s)

	fmt.Println(comma(b))
}
