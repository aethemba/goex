package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"linear algebra":        {"calculus"},
	"data structures":       {"discrete math"},
	"discrete math":         {"intro to programming"},
	"databases":             {"data structures"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func CheckCycles(m map[string][]string) {
	for key, value := range m {
		// for every key in the map, get the values
		for _, v := range value {
			// For every value of a map, check the map to see its values
			for _, vv := range m[v] {
				if vv == key {
					fmt.Println("Cycle! %s %s\n", key, v)
				}
			}
		}
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				// for _, i := range order {
				// 	if i == item {
				// 		fmt.Println("Cycle warning: %s", item)
				// 	}
				// }
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)

	CheckCycles(m)
	return order
}
