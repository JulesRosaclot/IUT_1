package main
import "fmt"
func f(x int, y int) {
	x = x + y
	fmt.Println(x)
}

func main() {
	var n int = 1
	var m int = 2
	f(n, m)

}
