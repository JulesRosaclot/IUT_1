package tab_trie

import (
	"testing"
)
func Test1(t *testing.T) {
	var tab []int = []int{1, 2, 4, 6, 8, 9}
	var a bool
	var b int
	a, b = recherche(tab, 6)
	if a != true || b != 3 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var tab []int
	var a bool
	var b int
	a, b = recherche(tab, 1)
	if a != false || b != 0 {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	var tab []int = []int{}
	var a bool
	var b int
	a, b = recherche(tab, 1)
	if a != false || b != 0 {
		t.Fail()
	}
}
func Test4(t *testing.T) {
	var tab []int = []int{2, 8, 4, 9, 6, 5, 3}
	var a bool
	var b int
	a, b = recherche(tab, 10)
	if a != false || b != 0 {
		t.Fail()
	}
}
