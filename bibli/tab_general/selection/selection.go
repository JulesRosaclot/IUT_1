package tab_general

func selection(tab []int, x int) (rep []int) {

	if tab == nil {
		return nil
	}

	if len(tab) == 0 {
		return nil
	}

	for i := 0; i < len(tab); i++ {
		if tab[i] <= x {
			rep = append(rep, tab[i])
		}
	}
	if len(rep) == 0 {
		return nil
	}
	return rep
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
