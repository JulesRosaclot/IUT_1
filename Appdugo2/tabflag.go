package main 

import ("flag"
		"fmt"
)

func table (x int, y int, z int, t []int) []int {
	for i := 0; i <= y; i++ {
		var prod int = x*i
		if prod%z != 0 {
			t = append(t, prod)
		}
	}
	return t
}

func main(){
	var x, y, z int
	var t []int = make([]int, 0, y+1)
	flag.IntVar(&x, "x", 0, "...")
	flag.IntVar(&y, "y", 0, "...")
	flag.IntVar(&z, "z", 1, "...")
	flag.Parse()
	fmt.Println(table(x, y, z, t))
}
	
