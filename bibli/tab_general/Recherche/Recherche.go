package tab_general

func Recherche(tab []int, x int) (rep bool, res int) {
	if tab == nil {
		return false, -1
	}
	if len(tab) == 0{
		return false, -1
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] == x {
			return true, i
		}
	}
	return false, -1
}
