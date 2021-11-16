package tri_tab

func tri_fusion(tab1 []int) (trier []int) {
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
	return trier
}
func recursive(tab []int) (tri []int) {
	if tab == nil {
		tri = nil 
	}
	for i := 1; i < len(tab); i++{
		var v int = tab[i]
		var j int = i - 1
		for j >= 0 && tab[j]> v{
			tab[j + 1] = tab[j]
			j--
		}
		tab[j + 1] = v
	}
	return tab
}

func comp_tab(a []int, b []int) (c bool) {
	if len(a) != len(b) {
		return false
	}
	if a == nil {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}

	}
	return true
}
