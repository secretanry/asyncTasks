package main

func GroupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, t := range temps {
		key := int(t/10) * 10
		groups[key] = append(groups[key], t)
	}
	return groups
}
