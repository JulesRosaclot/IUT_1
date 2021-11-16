package tri_tab

func tri_inser2(tab []int)(tri []int){
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