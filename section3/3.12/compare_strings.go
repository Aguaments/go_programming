package main

import(
		"fmt"
)

func compare_string(s1, s2 string) bool{
		flag := 0
		if len(s1) != len(s2){
				return false
		}else{
				for _, c1 := range s1 {	
						for _, c2 := range s2 {
								if c1 == c2{
										flag = 0
										break
								}else{
										flag = 1
								}
						}
						if flag == 1{
								return false
						}
				}
		return true
		}
}

func main(){
		s1 := "abce"
		s2 := "ebca"
		if compare_string(s1, s2){
				fmt.Println("True")
		}else{
				fmt.Println("False")
		}
}
