package zadatci

import "fmt"

func Zadatak1(proizvodi map[string]int) {

	najskupljiProizvod := ""
	najskupljiProizvodCijena := 0

	// Prvi dio
	for proizvod, cijena := range proizvodi {
		if cijena > najskupljiProizvodCijena {
			najskupljiProizvodCijena = cijena
			najskupljiProizvod = proizvod
		}
	}

	fmt.Printf("Najskuplji proizvod je %s sa cijenom %d \n", najskupljiProizvod, najskupljiProizvodCijena)

	// Drugi dio
	const TRIGGER_CIJENA int = 150

	fmt.Print("\n------------------[Proizvodi sa cijenom vecom od 150]----------------- \n")
	for proizvod, cijena := range proizvodi {
		if cijena > TRIGGER_CIJENA {
			fmt.Printf("Proizvod: %s koji košta %d ima cijenu vecu od %d \n", proizvod, cijena, TRIGGER_CIJENA)
		}
	}
}
