package designanorderedstream

import (
	"reflect"
	"testing"
)

func TestOrderedStream_Insert(t *testing.T) {
	type args struct {
		idKey int
		value string
	}
	tests := []struct {
		name     string
		this     OrderedStream
		inserted map[int]string
		args     args
		want     []string
	}{
		{
			this: Constructor(1),
			args: args{
				idKey: 1,
				value: "a",
			},
			want: []string{"a"},
		},
		{
			this:     Constructor(2),
			inserted: map[int]string{1: "a"},
			args: args{
				idKey: 2,
				value: "b",
			},
			want: []string{"b"},
		},
		{
			this:     Constructor(2),
			inserted: map[int]string{2: "b"},
			args: args{
				idKey: 1,
				value: "a",
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.inserted {
				tt.this.Insert(k, v)
			}

			if got := tt.this.Insert(tt.args.idKey, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedStream.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
