package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// Ouverture du fichier
	var filePath string = "tonfichier"
	var myFile *os.File
	var err error
	var i int
	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Préparation de la lecture
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)

	// Lecture des lignes du fichier
	for scanner.Scan() {
		log.Print("Je viens de lire : ", scanner.Text())
		i = i + 1 
		if i == 5 {
			break 
		}else {
			continue
		}
	}

	// Vérification que tout s'est bien passé
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("le fichier contient ", i, " lignes")
}
