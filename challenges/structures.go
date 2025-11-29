package challenges
import "strings"
// WordFrequency counts the frequency of each word in a given text.
// The function should be case-insensitive (treat "Go" and "go" as the same word).
// It should ignore punctuation (e.g., "Go!" should be treated as "Go").
//
// Examples:
// Input: "Hello world! Hello Go."
// Output: map[string]int{"hello": 2, "world": 1, "go": 1}
//
// Input: "Test, test, TEST"
// Output: map[string]int{"test": 3}
func WordFrequency(text string) map[string]int {
	// TODO: Implement this function
	text = strings.ToLower(text)

	punctuations := []string {",",".",";",":","!","?"}

	for _, p := range punctuations {
		text = strings.ReplaceAll(text, p, "")
	}

	words := strings.Fields(text)

	result := map[string]int {}

	for _, word := range words {
		result[word]++
	}
	return result
	
}
