package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var myFile *os.File
	var err error
	myFile, err = os.Create("tonfichier")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(myFile, "coucou")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(myFile, "bye bye")
	if err != nil {
		log.Fatal(err)
	}

	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
