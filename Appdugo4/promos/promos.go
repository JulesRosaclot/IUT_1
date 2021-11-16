package promos

type étudiant struct {
	prenom  string
	nom     string
	notes   []float64
}

func (e étudiant) moy() float64 {
	var somme float64
	for i := 0; i < len(e.notes); i++ {
		somme = somme + float64(e.notes[i])
	}
	var moyenne float64 = somme / float64(len(e.notes))
	return moyenne
}
func calc(prenom string, nom string, notes []float64)(prenom1 string, nom1 string, moyenne float64){
	if notes == nil {
		
	}
	var e étudiant = étudiant{prenom, nom, notes}
	return e.prenom, e.nom, e.moy()
}
