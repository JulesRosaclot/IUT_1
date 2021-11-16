package tab_trie

func recherche(tab []int, x int)(res bool, rep int){
	for i := 0; i < len(tab); i++{
		if tab[i] > x {
			return true, i-1
		}
	}
	return false, 0
}
