package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var ErrNo = fmt.Errorf("Your submitted number must be higher than 0")
	var number int = rand.IntN(500) // number := rand.IntN(500)", also possible
	var submitted int

	for {
		fmt.Println("Guess the mystery number:")
		_, err := fmt.Scan(&submitted) // "%d" for int if formatted
		/* vars als Werte übergeben (Kopie), wenn Wert actually geändert werden soll
		dann der Funktion die Adresse der var übergeben -> "&" übergibt Referenz der var*/

		if err != nil {
			fmt.Println("Invalid! Submit number")
			// clears input buffer -> avoid infinite loops without a chance to fix
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if submitted <= 0 {
			fmt.Println(ErrNo)
			continue
		}

		switch {
		case submitted > number:
			fmt.Println("The mystery number is lower.")
			continue

		case submitted < number:
			fmt.Println("The mystery number is higher")
			continue

		case submitted == number:
			fmt.Println("Congrats you guessed the number!")
			return
		}

	}

}
