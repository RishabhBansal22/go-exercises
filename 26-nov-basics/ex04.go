package main

// * Given the slice: `votes := []string{"Go", "Python", "Go", "Java", "Go", "Python"}`.
// * Create an empty map to store the results.
// * Iterate through the slice and count the frequency of each language.
// * Print the final tally (e.g., "Go: 3, Python: 2...").

import "fmt"

func Counter(votes []string) map[string]int {
	countMap := make(map[string]int)
	for _, val := range votes {
		countMap[val]++
	}
	return countMap
}

func main() {
	votes := []string{"Go", "Python", "Go", "Java", "Go", "Python"}
	res := Counter(votes)

	// Print formatted output
	for lang, count := range res {
		fmt.Printf("%s: %d\n", lang, count)
	}
}
