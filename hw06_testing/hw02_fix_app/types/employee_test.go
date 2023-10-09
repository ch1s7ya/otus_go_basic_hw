package types

import "testing"

func TestEmployee_String(t *testing.T) {
	type fields struct {
		UserID       int
		Age          int
		Name         string
		DepartmentID int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Valid employee",
			fields: fields{
				UserID:       1,
				Age:          30,
				Name:         "Evgeny",
				DepartmentID: 29,
			},
			want: "User ID: 1; Age: 30; Name: Evgeny; Department ID: 29;\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Employee{
				UserID:       tt.fields.UserID,
				Age:          tt.fields.Age,
				Name:         tt.fields.Name,
				DepartmentID: tt.fields.DepartmentID,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("Employee.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
