package tri_tab

import (
	"testing"
)

func Test1(t *testing.T){
	var tab []int = []int {12, 32, 7, 23, 9}
	var tri []int = tri_inser2(tab)
	var res []int = []int {7, 9, 12, 23, 32}
	if comp_tab(tri, res) != true {
		t.Fail()
	}
}
func Test2(t *testing.T){
	var tab []int = []int {}
	var tri []int = tri_inser2(tab)
	var res []int = []int {7, 9, 12, 23, 32}
	if comp_tab(tri, res) != false {
		t.Fail()
	}
}
func Test3(t *testing.T){
	var tab []int = []int {12, 32, 7, 23, 9}
	var tri []int = tri_inser2(tab)
	var res []int = []int {9, 7, 12, 23, 32}
	if comp_tab(tri, res) != false {
		t.Fail()
	}
}
func Test4(t *testing.T){
	var tab []int
	var tri []int = tri_inser2(tab)
	var res []int 
	if comp_tab(tri, res) != false {
		t.Fail()
	}
}