package tab_trie

func dichotomie(tab []int, x int) (b bool, n int) {
	if tab == nil{
		return false, -1
	}
	if len(tab) == 0{
		return false, -1
	}

	var debut int = 0
	var fin int = len(tab)
	for debut < fin {
		var milieu int = (debut + fin) / 2
		if tab[milieu] == x {
			return true, milieu
		} else if tab[milieu] > x {
			fin = milieu
		} else if tab[milieu] < x {
			debut = milieu + 1
		}
	}
	return false, -1
}
