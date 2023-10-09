package reader

import (
	"testing"

	"github.com/ch1s7ya/otus_go_basic_hw/hw06_testing/hw02_fix_app/types"
	"github.com/stretchr/testify/require"
)

func TestReadJSON(t *testing.T) {
	tests := []struct {
		desc     string
		path     string
		expected []types.Employee
	}{
		{
			desc: "Example from hw02",
			path: "../../test/testdata/data1.json",
			expected: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				},
				{
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
		},
		{
			desc:     "Empty json",
			path:     "../../test/testdata/data2.json",
			expected: []types.Employee{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			result, err := ReadJSON(tt.path)
			if err != nil {
				t.Fatalf("Unexpeted error: %v", err)
			}
			require.Equal(t, tt.expected, result)
		})
	}
}
