package tab_general

import (
	"testing"
)

func Test1(t *testing.T) {
	var tab []int = []int{2, 8, 4, 9, 6, 5, 3}
	var a []int
	var b []int = []int{2, 4, 3}
	a = selection(tab, 4)
	if comp_tab(a, b) != true {
		t.Fail()
	}
}
func Test2(t *testing.T) {
	var tab []int
	var a []int
	var b []int = []int{2, 4, 3}
	a = selection(tab, 4)
	if comp_tab(a, b) != false {
		t.Fail()
	}
}
func Test3(t *testing.T) {
	var tab []int
	var a []int
	var b []int
	a = selection(tab, 4)
	if comp_tab(a, b) != false {
		t.Fail()
	}
}
func Test4(t *testing.T) {
	var tab []int = []int{2, 8, 4, 9, 6, 5, 3}
	var a []int
	var b []int = []int{}
	a = selection(tab, 1)
	if comp_tab(a, b) != false {
		t.Fail()
	}
}
func Test5(t *testing.T) {
	var tab []int = []int{}
	var a []int
	var b []int
	a = selection(tab, 4)
	if comp_tab(a, b) != false {
		t.Fail()
	}
}
func Test6(t *testing.T) {
	var tab []int = []int{2, 8, 4, 9, 6, 5, 3}
	var a []int
	var b []int = []int{1, 4, 3}
	a = selection(tab, 4)
	if comp_tab(a, b) != false {
		t.Fail()
	}
}
