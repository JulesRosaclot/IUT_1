package main

import "fmt"
func decoupe (x string) (string, string){
	var espace int
	for i := 0; i < len(x); i++ {
		if  x[i] == ' '{
			break
		} else {
			espace = espace + 1
		}
	}
	return x[:espace], x[espace+1:]
}
func main() {
	var c string = "HELLO WORLD"
	var c1, c2 string
	c1, c2 = decoupe(c)
	fmt.Println(c1)
	fmt.Println(c2)
	
}
