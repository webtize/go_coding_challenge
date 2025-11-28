package challenges

import (
	"testing"
)

func TestParseLogEntries(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		wantErr bool
		wantLen int
	}{
		{
			name: "valid logs",
			input: []string{
				"2023-10-27T10:00:00Z INFO: Application started",
				"2023-10-27T10:05:00Z ERROR: Database connection failed",
			},
			wantErr: false,
			wantLen: 2,
		},
		{
			name: "invalid timestamp",
			input: []string{
				"invalid-time INFO: Test",
			},
			wantErr: true,
			wantLen: 0,
		},
		{
			name: "missing level separator",
			input: []string{
				"2023-10-27T10:00:00Z INFO Application started",
			},
			wantErr: true,
			wantLen: 0,
		},
		{
			name:    "empty input",
			input:   []string{},
			wantErr: false,
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLogEntries(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLogEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.wantLen {
				t.Errorf("ParseLogEntries() got length = %v, want %v", len(got), tt.wantLen)
			}
			if !tt.wantErr && len(got) > 0 {
				// Verify first element parsing correctness for valid case
				if got[0].Level != "INFO" && got[0].Level != "ERROR" {
					t.Errorf("ParseLogEntries() parsed level = %v, expected INFO or ERROR", got[0].Level)
				}
				if got[0].Timestamp.IsZero() {
					t.Error("ParseLogEntries() parsed zero timestamp")
				}
			}
		})
	}
}
