package funcaovariatica

func SomarValores(valores ...float64) float64 {
	total := 0.0

	for _, valor := range valores {
		total += valor
	}

	return total
}
