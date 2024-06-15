package goalparserinterpretation

import "testing"

func Test_interpret(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				command: "G",
			},
			want: "G",
		},
		{
			args: args{
				command: "()",
			},
			want: "o",
		},
		{
			args: args{
				command: "(al)",
			},
			want: "al",
		},
		{
			args: args{
				command: "G()(al)",
			},
			want: "Goal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interpret(tt.args.command); got != tt.want {
				t.Errorf("interpret() = %v, want %v", got, tt.want)
			}
		})
	}
}
