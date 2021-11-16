package main

import "fmt"

func main() {
	var année int = 2015
	if (année % 400 == 0) || (année % 4 == 0){
		fmt.Println("L'année", année, "est bissextile")
	}else{ 
		fmt.Println("L'année", année, "n'est pas bissextile")
	}
}
