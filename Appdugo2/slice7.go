package main

import "fmt"

func main() {
	var length int = 5
	var capacity int = 7
	var tab []int = make([]int, length, capacity)
	var t []int = tab
	var max int = 10
	for i := length; i <= max; i++ {
		t = append(t, i)
		t[0] = i
		fmt.Println("Tour de boucle", i)
		fmt.Println("\ttab =", tab, "len(tab) =", len(tab), "cap(tab) =", cap(tab))
		fmt.Println("\tt =", t, "len(t) =", len(t), "cap(t) =", cap(t))
	}
}
