package main

import "fmt"

func main() {
	var a int = 9
	var b int = 10
	var c int = 50
	if a > b && a > c {
		fmt.Println("La variable la plu grande est a :", a)
	}else if b > a && b > c{
		fmt.Println("La variable la plu grande est b :", b)
	}else if c > a && c > b{
		fmt.Println("La variable la plu grande est c :", c)
	}
}
