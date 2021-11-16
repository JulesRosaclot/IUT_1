package main

import "fmt"

func main() {
	var tab1 []int = []int{8}
	var tab2 []int = []int{2}
	var trier []int
	var i int = 0
	var j int = 0
	var length int
	var lengthfin int

	if len(tab1) > len(tab2) {
		length = len(tab2)
		lengthfin = len(tab1)
	} else {
		length = len(tab1)
		lengthfin = len(tab2)
	}

	for i < length {
		if tab1[i] > tab2[j] {
			trier = append(trier, tab2[j])
			j++
		} else if tab1[i] < tab2[j] {
			trier = append(trier, tab1[i])
			i++
		} else {
			trier = append(trier, tab2[j])
			i++
			j++
		}
	}

	if lengthfin == len(tab2) {
		for j < lengthfin {
			trier = append(trier, tab2[j])
			j++
		}
	} else {
		for i < lengthfin {
			trier = append(trier, tab1[i])
			i++
		}
	}
	fmt.Println(trier)
}

