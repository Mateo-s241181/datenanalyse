package main

import (
	"fmt"

	"github.com/tel23a-inf/data-analysis/dice"
	"github.com/tel23a-inf/data-analysis/testseries"
)

// readUserInput fragt den Benutzer nach der Anzahl der Würfe und der Anzahl der Würfel.
// Die Funktion liefert beide Werte zurück.
func readUserInput() (int, int) {
	var d, n int

	fmt.Println("Mit wie vielen Würfeln wollen Sie würfeln?")
	fmt.Scanln(&d)
	fmt.Printf("Sie würfeln mit %d Würfeln. \n\n", d)
	fmt.Println("Wie oft wollen Sie würfeln?")
	fmt.Scanln(&n)
	fmt.Printf("Sie würfeln %d mal. \n\n", n)

	return d, n
}

// printDiceStatistics berechnet die Statistik für die Würfelwürfe und gibt sie aus.
func printDiceStatistics(rollResults []int) {
	testseries.PrintDistribution(rollResults)

	fmt.Printf("\n-------------------\n")

	//testseries.PrintHistogram(rollResults)
}

// main kombiniert die anderen Funktionen zu einem Programm.
func main() {

	dices, throws := readUserInput()

	ResultSlice := dice.RollMany(dices, throws)

	printDiceStatistics(ResultSlice)

}
