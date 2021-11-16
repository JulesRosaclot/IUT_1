package main
import "fmt"

func addAA(aP *int, bP int) {
	*aP = *aP + bP
}
func main(){
	var a int
	var ptr_a *int
	
	ptr_a = &a
	
	fmt.Println(&a)
	fmt.Println(ptr_a)
	
	*ptr_a =10
	fmt.Println(a)
	fmt.Println(ptr_a)
	
	var b int = 20
	fmt.Println(a, b)
	addAA(&a, b)
	fmt.Println(a, b)
}
