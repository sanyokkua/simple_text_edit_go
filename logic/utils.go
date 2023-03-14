package logic

import (
	"sort"
	"strings"
)

func sortAsc(text []string) {
	sort.Strings(text)
}

func sortAscIgnoreCase(text []string) {
	sort.Slice(text, func(i, j int) bool {
		return strings.ToLower(text[i]) < strings.ToLower(text[j])
	})
}

func sortDesc(text []string) {
	sort.Slice(text, func(i, j int) bool {
		return text[i] > text[j]
	})
}

func sortDescIgnoreCase(text []string) {
	sort.Slice(text, func(i, j int) bool {
		return strings.ToLower(text[i]) > strings.ToLower(text[j])
	})
}
