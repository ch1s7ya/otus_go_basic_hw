package main

import "testing"

func Test_drawBoardChess(t *testing.T) {
	tests := []struct {
		name      string
		boardSize int
		want      string
	}{
		{
			name:      "board chess without size",
			boardSize: 0,
			want:      "",
		},
		{
			name:      "board chess with size equal to 1",
			boardSize: 1,
			want:      " \n",
		},
		{
			name:      "board chess with size equal to 4",
			boardSize: 4,
			want:      " # #\n# # \n # #\n# # \n",
		},
		{
			name:      "board chess with size equal to 8",
			boardSize: 8,
			want:      " # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := drawBoardChess(tt.boardSize); got != tt.want {
				t.Errorf("drawBoardChess() = %v, want %v", got, tt.want)
			}
		})
	}
}
