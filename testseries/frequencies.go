package testseries

import "github.com/tel23a-inf/data-analysis/intlists"

// AbsoluteFrequencies erwartet eine Liste mit den Werten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den absoluten Häufigkeiten für jeden Wert
// zwischen dem Minimum und dem Maximum der Messreihe.
func AbsoluteFrequencies(values []int) []int {

	//chechValues beinhaltet alle Werte zwischen Minimum und Maximum der Messreihe
	checkValues := intlists.ValueRange(values)
	//absValues soll die gleiche Kapazität wie checkValues haben
	absValues := make([]int, len(checkValues))

	for i := range checkValues {

		//counter ist anfangs 0
		counter := 0

		//Es wird gezählt, wie oft checkValues[i] in values[j] vorkommt
		for j := range values {

			if checkValues[i] == values[j] {
				counter++
			}
		}

		//counter wird in absValues geschrieben
		absValues[i] = counter
	}

	return absValues
}

// RelativeFrequencies erwartet eine Liste mit absoluten Häufigkeiten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den relativen Häufigkeiten.
func RelativeFrequencies(values []int) []float64 {
	freq := make([]float64, len(values))

	//Anzahl der Messungen berechnen
	sum := 0

	for i := range values {
		sum += values[i]
	}

	//Jede absolute Häufigkeit durch die Anzahl der Messungen Teilen und in freq speichern

	for i := range values {
		freq[i] = float64(values[i]) / float64(sum)
	}

	return freq
}
