package main
import "fmt"
func estParfait(n uint64) (b bool) {
	var n uint64 = 28
	var t []uint64
	var i uint64
	var res uint64
	for i = 1; i < n; i++ {
		if n%i == 0 {
			t = append(t, i)
		} else {
			return false
		}
	}
	for i = 0; i < uint64(len(t)); i++ {
		res = res + t[i]
	}
	if res == n {
		b = true
	} else {
		b = false
	}
	return b
}
fmt.Println(estParfait(res, b))