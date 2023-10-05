package main

import (
	"testing"
)

func Test_calculateArea(t *testing.T) {
	tests := []struct {
		name    string
		shape   any
		want    float64
		wantErr bool
	}{
		{
			name:    "Calculate circle area",
			shape:   &circle{5},
			want:    78.53981633974483,
			wantErr: false,
		},
		{
			name:    "Calculate rectangle area",
			shape:   &rectangle{10, 5},
			want:    50.00,
			wantErr: false,
		},
		{
			name:    "Calculate triangle area",
			shape:   &triangle{8, 6},
			want:    24.00,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateArea(tt.shape)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateArea() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
