package stringcompression

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantChars []byte
	}{
		{
			args: args{
				chars: []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
			},
			want:      6,
			wantChars: []byte{'a', '3', 'b', '2', 'a', '2'},
		},
		{
			args: args{
				chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			},
			want:      6,
			wantChars: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			args: args{
				chars: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
			},
			want:      3,
			wantChars: []byte{'a', '1', '0'},
		},
		{
			args: args{
				chars: []byte{'a', 'a'},
			},
			want:      2,
			wantChars: []byte{'a', '2'},
		},
		{
			args: args{
				chars: []byte{'a'},
			},
			want:      1,
			wantChars: []byte{'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.args.chars[0:tt.want], tt.wantChars) {
				t.Errorf("compress() = %v, want %v", string(tt.args.chars), string(tt.wantChars))
			}
		})
	}
}
