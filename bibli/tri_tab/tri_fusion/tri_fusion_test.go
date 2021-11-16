package tri_tab

import (
	"testing"
)

func Test1(t *testing.T){
	var tab1 []int = []int {12, 32, 7, 23, 9}
	var tab2 []int = []int {8, 4, 40, 25, 50}
	var tri []int = tri_fusion(tab1, tab2)
	var tri2 []int = recursive(tri)
	var res []int = []int {4, 7, 8, 9, 12, 23, 25, 32, 40, 50}
	if comp_tab(tri2, res)!= true {
		t.Fail()
	}
}