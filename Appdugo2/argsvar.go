package main 

import (
	"fmt"
	"os"
	"strconv"
)
func main (){
	var v int
	var e error
	var i int = 3
	v, e = strconv.Atoi(os.Args[i])
	if e == nil {
		fmt.Println("L'argument", i, "vaut", v)
	}

}
