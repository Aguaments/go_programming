package main

import(
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
	"os"
)

func main(){
	if len(os.Args) == 2{
		fmt.Printf("SHA256: %s\n", sha256.Sum256([]byte(os.Args[1])))
	}else if len(os.Args) == 3{
		if os.Args[2] == "384"{
			fmt.Printf("SHA384: %s\n", sha512.Sum384([]byte(os.Args[1])))
		}else if os.Args[2] == "512"{
			fmt.Printf("SHA512: %s\n", sha512.Sum512([]byte(os.Args[1])))
		}else{
			fmt.Println("256(default) 384 512")
		}
	}else{
		fmt.Println("Format: command [data] [nbits]")
	}
}
