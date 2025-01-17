package designaddandsearchwordsdatastructure

import "testing"

func TestWordDictionary_Search(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		this *WordDictionary
		args args
		want bool
	}{
		{
			this: func() *WordDictionary {
				d := Constructor()
				d.AddWord("word")
				return &d
			}(),
			args: args{
				word: "word",
			},
			want: true,
		},
		{
			this: func() *WordDictionary {
				d := Constructor()
				d.AddWord("want")
				return &d
			}(),
			args: args{
				word: "word",
			},
			want: false,
		},
		{
			this: func() *WordDictionary {
				d := Constructor()
				d.AddWord("waitress")
				return &d
			}(),
			args: args{
				word: "wait",
			},
			want: false,
		},
		{
			this: func() *WordDictionary {
				d := Constructor()
				d.AddWord("card")
				d.AddWord("core")
				return &d
			}(),
			args: args{
				word: "cor.",
			},
			want: true,
		},
		{
			this: func() *WordDictionary {
				d := Constructor()
				d.AddWord("card")
				d.AddWord("coreandr")
				return &d
			}(),
			args: args{
				word: "cor.",
			},
			want: false,
		},
		{
			this: func() *WordDictionary {
				d := Constructor()
				d.AddWord("card")
				d.AddWord("core")
				return &d
			}(),
			args: args{
				word: "c.r.",
			},
			want: true,
		},
		{
			this: func() *WordDictionary {
				d := Constructor()
				d.AddWord("card")
				d.AddWord("core")
				return &d
			}(),
			args: args{
				word: "c..e",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Search(tt.args.word); got != tt.want {
				t.Errorf("WordDictionary.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
