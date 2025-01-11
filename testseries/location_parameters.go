package testseries

import (
	"math"
	"slices"

	"github.com/tel23a-inf/data-analysis/intlists"
)

// Average erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Durchschnittswert.
// Ist die Liste leer, wird 0.0 zurückgegeben.
func Average(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	//Summe aller Werte durch Anzahl der Werte
	return float64(intlists.Sum(values)) / float64(len(values))
}

// Median erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Median.
// Ist die Liste leer, wird 0 zurückgegeben.
func Median(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sorted := make([]int, len(values))
	copy(sorted, values)

	//Slice wird Sortiert
	slices.Sort(sorted)

	//Wenn die Anzahl der Werte gerade ist:
	if len(sorted)%2 == 0 {

		//Beiden Mittleren Werte
		almostMedians := []int{sorted[len(sorted)/2], sorted[len(sorted)/2-1]}

		return int(Average(almostMedians))
	}

	//Bei ungerader Anzahl wird der Wert in der Mitte zurückgegeben
	return sorted[(len(sorted)-1)/2]
}

// Mode erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Wert, der am häufigsten in der Liste vorkommt.
// Falls mehrere Werte am häufigsten vorkommen, wird der kleinste dieser Werte
// zurückgegeben. Ist die Liste leer, wird 0 zurückgegeben.
func Mode(values []int) int {
	if len(values) == 0 {
		return 0
	}

	//Absolute Häufigkeit ermitteln
	absValues := AbsoluteFrequencies(values)
	//Gecheckte Werte ermitteln
	checked := intlists.ValueRange(values)

	//Maximalen Wert aus absValues herauslesen
	maxAbsValue := intlists.Max(absValues)

	//Korrespondierende Position(-en) aus absValues lesen und in positions schreiben
	mostFrequent := []int{}

	for i := range absValues {
		if absValues[i] == maxAbsValue {
			mostFrequent = append(mostFrequent, checked[i])
		}
	}

	return intlists.Min(mostFrequent)
}

// GeometricMean erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das geometrische Mittel.
// Ist die Liste leer, wird 0.0 zurückgegeben.
//
// Anmerkung: Das geometrische Mittel ist definiert als
// die n-te Wurzel aus dem Produkt der n Werte.
func GeometricMean(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	return math.Pow(float64(intlists.Product(values)), 1/float64(len(values)))
}

// HarmonicMean erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das harmonische Mittel.
// Ist die Liste leer, wird 0.0 zurückgegeben.
//
// Anmerkung: Das harmonische Mittel ist definiert als
// die Kehrwert des Durchschnitts der Kehrwerte der n Werte.
func HarmonicMean(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	//Slice initialisieren
	kehrwerte := make([]float64, len(values))

	//Kehrwerte bilden und in Slice schreiben
	for i := range values {
		kehrwerte[i] = 1 / float64(values[i])
	}

	//Summe der Kehrwerte berechnen
	sum := 0.0
	for i := range kehrwerte {
		sum += kehrwerte[i]
	}

	//Summe durch länge teilen
	floatAverage := sum / float64(len(kehrwerte))

	//Durchschnitt der Summe der Kehrwerte
	return 1 / floatAverage
}
