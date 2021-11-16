package main

import (
	"fmt"
	"log"
	"os"
	"io"
)

func main() {
	// Ouverture du fichier
	var filePath string = "tonfichier"
	var myFile *os.File
	var err error
	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Lecture d'une ligne
	var nbLus int
	var uneChaine string
	var i int 
	nbLus, err = fmt.Fscanln(myFile, &uneChaine)
	for err != io.EOF {
		if err != nil{
			log.Fatal(err)
		}else{
			log.Print("J'ai lu ", nbLus, " valeurs.")
			log.Print("uneChaine = ", uneChaine)
			i = i + nbLus			
		}
		nbLus, err = fmt.Fscanln(myFile, &uneChaine)
	}
	log.Print("il y'avait ", i, " valeurs")
	
	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
