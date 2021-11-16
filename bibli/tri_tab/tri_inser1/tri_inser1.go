package tri_tab

func tri_inser1(tab []int)(tri []int){
	for i := 0; i < len(tab); i++{
		tri = append(tri, tab[i])
		var j int = 0
		var v int = tab[i]
		for j < len(tri) && tri[j] <= v{
			j++
		}
		for j < len(tri){
			var tmp int = tri[j]
			tri[j] = v
			v = tmp
			j++
		}
	}
	return tri
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