package tab_trie

import (
	"testing"
)

func Test1(t *testing.T) {
	var tab []int = []int{1, 2, 4, 6, 8, 9}
	var x int = 6
	var a bool
	var b int
	a, b = dichotomie(tab, x)
	if a != true || b != 3 {
		t.Fail()
	}
}
func Test2(t *testing.T) {
	var tab []int
	var x int = 6
	var a bool
	var b int
	a, b = dichotomie(tab, x)
	if a != false || b != -1 {
		t.Fail()
	}
}
func Test3(t *testing.T) {
	var tab []int = []int{}
	var x int = 6
	var a bool
	var b int
	a, b = dichotomie(tab, x)
	if a != false || b != -1 {
		t.Fail()
	}
}
func Test4(t *testing.T) {
	var tab []int = []int{280, 320, 400, 470, 560, 900}
	var x int = 420
	var a bool
	var b int
	a, b = dichotomie(tab, x)
	if a != false || b != -1 {
		t.Fail()
	}
}
