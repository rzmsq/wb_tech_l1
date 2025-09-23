package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func findAnagram(strs []string) map[string][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		str := sortString(strings.ToLower(s))
		groups[str] = append(groups[str], s)
	}

	out := make(map[string][]string)
	for _, group := range groups {
		if len(group) < 2 {
			continue
		}
		key := group[0]
		sort.Strings(group)
		out[key] = group
	}

	return out
}

func main() {
	in := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	result := findAnagram(in)

	fmt.Println(result)
}
