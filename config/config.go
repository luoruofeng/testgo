package config

import (
	"fmt"
)

var s string

func init(){
	s = "sss"
	fmt.Println("init")
}

func TestFunc(){
	fmt.Println(s)
	fmt.Println("test func\n")
}
