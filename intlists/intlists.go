package intlists

// Min berechnet das Minimum einer Liste von Integer-Werten.
// Funktioniert nur für nicht-leere Listen.
func Min(values []int) int {
	min := values[0]

	for i := range values {

		if i != 0 && values[i] < min {
			min = values[i]
		}
	}
	return min
}

// Max berechnet das Maximum einer Liste von Integer-Werten.
// Funktioniert nur für nicht-leere Listen.
func Max(values []int) int {
	max := values[0]
	for i := range values {

		if i != 0 && values[i] > max {
			max = values[i]
		}
	}
	return max
}

// ValueRange erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert eine lückenlose Liste mit allen Zahlen zwischen
// dem Minimum und dem Maximum der Messreihe.
func ValueRange(values []int) []int {
	result := []int{}

	i := Min(values)

	for i <= Max(values) {
		result = append(result, i)

		i++
	}

	return result

}

// Sum erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert die Summe aller Werte.
func Sum(values []int) int {
	sum := 0

	for i := range values {
		sum += values[i]
	}

	return sum
}

// Product erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das Produkt aller Werte.
func Product(values []int) int {
	product := 1

	for i := range values {
		product *= values[i]
	}

	return product
}
