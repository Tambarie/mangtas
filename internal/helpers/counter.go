package helpers

import (
	"Mangtas_/internal/models"
	"regexp"
	"sort"
)

func Counter(txt string) []models.Sorter {
	var sorter []models.Sorter
	texts := regexp.MustCompile("[^0-9a-zA-Z]+").Split(txt,-1)

	m := map[string]int{}
	for _, word := range texts{
		m[word]++
	}

	for key, value := range m{
		sorter = append(sorter, models.Sorter{Value: value, Key: key})
	}

	sort.Slice(sorter, func(i, j int) bool {
		return sorter[i].Value > sorter[j].Value
	})

	total := len(sorter)
	if total > 10 {
		total = 10
	}

	return sorter[:total]
}