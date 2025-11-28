package challenges

import (
	"testing"
)

func TestProcessNumbers(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		wantSum int
		wantAvg float64
		wantMax int
		wantMin int
	}{
		{
			name:    "positive numbers",
			input:   []int{1, 2, 3, 4, 5},
			wantSum: 15,
			wantAvg: 3.0,
			wantMax: 5,
			wantMin: 1,
		},
		{
			name:    "mixed numbers",
			input:   []int{10, -2, 0},
			wantSum: 8,
			wantAvg: 8.0 / 3.0,
			wantMax: 10,
			wantMin: -2,
		},
		{
			name:    "single number",
			input:   []int{42},
			wantSum: 42,
			wantAvg: 42.0,
			wantMax: 42,
			wantMin: 42,
		},
		{
			name:    "empty slice",
			input:   []int{},
			wantSum: 0,
			wantAvg: 0.0,
			wantMax: 0,
			wantMin: 0,
		},
		{
			name:    "negative numbers",
			input:   []int{-5, -1, -10},
			wantSum: -16,
			wantAvg: -16.0 / 3.0,
			wantMax: -1,
			wantMin: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gotAvg, gotMax, gotMin := ProcessNumbers(tt.input)

			if gotSum != tt.wantSum {
				t.Errorf("ProcessNumbers() sum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotAvg != tt.wantAvg {
				t.Errorf("ProcessNumbers() avg = %v, want %v", gotAvg, tt.wantAvg)
			}
			if gotMax != tt.wantMax {
				t.Errorf("ProcessNumbers() max = %v, want %v", gotMax, tt.wantMax)
			}
			if gotMin != tt.wantMin {
				t.Errorf("ProcessNumbers() min = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}
