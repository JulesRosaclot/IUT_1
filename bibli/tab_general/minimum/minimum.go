package tab_general

func Minimum(tab []int) (rep bool, min int) {
	if tab == nil {
		return false, -1
	}
	if len(tab) == 0 {
		return false, -1
	}

	min = tab[0]
	for i := 1; i < len(tab); i++ {
		if tab[i] < min {
			min = tab[i]
		}
	}
	return true, min
}
