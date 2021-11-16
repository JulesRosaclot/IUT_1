package main
import "fmt"
func fval(x int) int {
	fmt.Println(2*x + 1)
	return 2*x + 1
}

func fref(a *int) {
	*a = 2*(*a) + 1
	fmt.Println(*a)
}

func main() {
	var n int = 5
	n = fval(n)
	
	var m int = 5
	fref(&m)
}
