package main 
import "fmt"
func main(){
	var tint []int = []int {1, 2, 5, 8, 7, 12, 9, 6, 3, 4}
	for i := 0; i < len(tint); i++ {
		tint[i] = tint[i] * 5

	}
	fmt.Println(tint)
}
