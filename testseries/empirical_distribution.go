package testseries

// EmpiricalDistribution erwartet eine Liste mit relativen Häufigkeiten einer Messreihe.
// Die Funktion liefert eine Liste, in der für jede Zahl der
// entsprechende Wert der empirischen Verteilungsfunktion steht.
func EmpiricalDistribution(relativeFreqs []float64) []float64 {
	emp := make([]float64, len(relativeFreqs))

	for i := range relativeFreqs {

		//Die kumulierte relative häufigkeit errechnen
		cumulativeFreq := 0.0

		for j := 0; j <= i; j++ {
			cumulativeFreq += relativeFreqs[j]
		}

		//in die liste emp schreiben
		emp[i] = cumulativeFreq
	}

	return emp
}

// Distribution erwartet eine Liste mit ganzzahligen Messwerten.
// Die Funktion liefert eine Liste mit den Werten der empirischen Verteilungsfunktion.
func Distribution(values []int) []float64 {
	list := []float64{0.0, 0.0}
	return list
}
