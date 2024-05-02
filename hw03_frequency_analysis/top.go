package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(input string) []string {
	counter := stringsCount(input)
	countToString := countToStrings(counter)
	topCounts := sortTopCounts(countToString)
	topWord := topWords(topCounts, countToString)

	resultLen := 10

	if len(topWord) < resultLen {
		resultLen = len(topWord)
	}

	result := make([]string, resultLen)

	copy(result, topWord)

	return result
}

func stringsCount(input string) map[string]int {
	result := make(map[string]int)
	s := strings.Fields(input)

	for _, v := range s {
		if _, ok := result[v]; !ok {
			result[v] = 0
		}

		result[v]++
	}

	return result
}

func countToStrings(in map[string]int) map[int][]string {
	result := make(map[int][]string)

	for k, v := range in {
		result[v] = append(result[v], k)
	}

	return result
}

func sortTopCounts(countToString map[int][]string) []int {
	result := make([]int, 0)

	for k := range countToString {
		result = append(result, k)
	}

	sort.Slice(result, func(i, j int) bool { return result[i] > result[j] })

	return result
}

func topWords(topCounts []int, countToString map[int][]string) []string {
	result := make([]string, 0, 10)

	for i := 0; i < len(topCounts); i++ {
		key := topCounts[i]
		words := countToString[key]
		sort.Strings(words)
		result = append(result, words...)
	}

	return result
}
