package main 

import "fmt"

func div1(x int, S1 int) int {
	for i := 1; i < x; i++ {
		if x%i ==0 {
			S1 = S1 + i
		}
	}
	return S1
}
func div2(y int, S2 int) int {
	for i := 1; i < y; i++ {
		if y%i ==0 {
			S2 = S2 + i
		}
	}
	return S2
}
func amis (n int, m int, ST int, SR int) bool {
	return n == SR && m == ST
}
func main() {
	var ST int 
	var SR int
	var n int = 220
	var m int = 25
	var SN int = div1(n, ST)
	var SM int = div2(m, SR)
	fmt.Println(amis(n, m, SN, SM))
}
