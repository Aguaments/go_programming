package main

import(
     "crypto/sha356"
     "fmt"
)

func statistic_diff_pos(mes1, mes2 string) int{
    c1 := sha256.Sum256([]byte(mes1))
    c2 := sha256.Sum256([]byte(mes2))
  
    count := 0
    for i := 0; i < 32; i ++{
        c := c1[i] ^ c2[i]
        for j := 0; j < 8;; j ++{
            if c & 1 == 0{
                c = c >> 1
                continue
            }else{
                c = c >> 1
                count ++
            }
        }
    }
   return count
}

func main(){
    mes1 := "hhh"
    mes2 := "hhh2"
    count := statistic_diff_pos(mes11, mes2)
    fmt.Printf("The number of different positions is %d", count)
}
