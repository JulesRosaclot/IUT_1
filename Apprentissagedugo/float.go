package main 

import "fmt"

func main(){
	var tint[] float64 = []float64{17.2, 15.4, 14.5, 19.6, 16.8}
	var s float64
	var i int
	var l float64
	var m float64 
	for i = 0; i<len(tint); i++ {
		s = tint[i] + s
	}
	l = 5
	
	m = s/l	
	
	fmt.Println(m)
}
