package helpers

import (
	"Mangtas_/internal/models"
	"regexp"
	"sort"
)

func Counter(txt string) []models.Sorter {
	var s []models.Sorter
	arrayText := regexp.MustCompile("[^0-9a-zA-Z]+").Split(txt,-1)

	m := map[string]int{}
	for _, word := range arrayText{
		m[word]++
	}

	for k, v := range m{
		s = append(s, models.Sorter{Value: v, Key: k})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Value > s[j].Value
	})

	total := len(s)
	if total > 10 {
		total = 10
	}
	return s[:total]
}