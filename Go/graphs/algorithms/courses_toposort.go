package algorithms

import "sort"

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discret math":          {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// Depth-first, post-order traversal
func topoSort(courses map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true // pre-order: to prevent cycles from causing stack overflow
				visitAll(courses[item])
				order = append(order, item) // post-order
			}
		}
	}
	var keys []string
	for key := range courses {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)

	return order
}
