package promos

import "testing"

func Test(t *testing.T) {
	var notes []float64 = []float64{15, 14, 18, 19, 17}
	var e étudiant = étudiant{"Jules", "Rosaclot", notes}
	var moyenne float64 = 16.6
	var moy float64 = e.moy()
	if moy != moyenne {
		t.Fail()
	}
}

func Test2 (t *testing.T) {
	var notes []float64 
	var e étudiant = étudiant{"Jules", "Rosaclot", notes}
	var moyenne float64 = 16.6
	var moy float64 = e.moy()
	if moy != moyenne {
		t.Fail()
	}
}