package testseries

import (
	"fmt"

	"github.com/tel23a-inf/data-analysis/intlists"
)

// PrintDistribution erwartet eine Liste mit Ergebnissen einer Integer-Messreihe.
// Die Funktion gibt die absolute und relative Häufigkeit sowie
// den Wert der empirischen Verteilungsfunktion für jede Zahl aus.
func PrintDistribution(values []int) {

	fmt.Println("Wert   Abs.   Rel.     Vert.")

	valueRange := intlists.ValueRange(values)
	absFreq := AbsoluteFrequencies(values)
	relFreq := RelativeFrequencies(absFreq)
	empVals := EmpiricalDistribution(relFreq)

	for i := range valueRange {
		fmt.Printf("%d       %d     %.2f     %.2f\n", valueRange[i], absFreq[i], relFreq[i], empVals[i])
	}
}

// PrintHistogram erwartet eine Liste mit Ergebnissen einer Integer-Messreihe.
// Die Funktion gibt ein Histogramm aus.
func PrintHistogram(values []int) {

	absFreq := AbsoluteFrequencies(values)
	valueRange := intlists.ValueRange(values)

	for i := range absFreq {

		fmt.Printf("%d: |", valueRange[i])

		for j := 0; j < absFreq[i]; j++ {

			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}
