package main 

import "fmt"

func main(){
	var tab []int = []int {12, 32, 7, 23, 9}
	var tri []int 
	for i := 0; i < len(tab); i++{
		tri = append(tri, tab[i])
		var j int = 0
		var v int = tab[i]
		for j < len(tri) && tri[j] <= v{
			j++
		}
		for j < len(tri){
			var tmp int = tri[j]
			tri[j] = v
			v = tmp
			j++
		}
	}	
	fmt.Println(tri)
}
