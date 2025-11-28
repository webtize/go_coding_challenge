package challenges

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "simple sentence",
			input: "Hello world",
			want:  map[string]int{"hello": 1, "world": 1},
		},
		{
			name:  "repeated words case insensitive",
			input: "Go go GO",
			want:  map[string]int{"go": 3},
		},
		{
			name:  "punctuation handling",
			input: "Hello, world! This is Go.",
			want:  map[string]int{"hello": 1, "world": 1, "this": 1, "is": 1, "go": 1},
		},
		{
			name:  "empty string",
			input: "",
			want:  map[string]int{},
		},
		{
			name:  "mixed whitespace and punctuation",
			input: "  one,  two;   three. ",
			want:  map[string]int{"one": 1, "two": 1, "three": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordFrequency(tt.input)
			if got == nil && len(tt.want) == 0 {
				// Accept nil map for empty result
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
