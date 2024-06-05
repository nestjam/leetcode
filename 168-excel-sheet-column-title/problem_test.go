package excelsheetcolumntitle

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				columnNumber: 1,
			},
			want: "A",
		},
		{
			args: args{
				columnNumber: 2,
			},
			want: "B",
		},
		{
			args: args{
				columnNumber: 26,
			},
			want: "Z",
		},
		{
			args: args{
				columnNumber: 27,
			},
			want: "AA",
		},
		{
			args: args{
				columnNumber: 701,
			},
			want: "ZY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
