package ransomnote

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	testCases := []struct {
		magazine    string
		ransomeNote string
		want        bool
	}{
		{
			magazine:    "a",
			ransomeNote: "a",
			want:        true,
		},
		{
			magazine:    "b",
			ransomeNote: "a",
			want:        false,
		},
		{
			magazine:    "aab",
			ransomeNote: "baa",
			want:        true,
		},
		{
			magazine:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			ransomeNote: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			want:        true,
		},
	}
	for _, tC := range testCases {
		t.Run(getDesc(tC.ransomeNote, tC.magazine), func(t *testing.T) {
			got := canConstruct(tC.ransomeNote, tC.magazine)

			if got != tC.want {
				t.Errorf("want '%v' but got '%v'", tC.want, got)
			}
		})
	}
}

func getDesc(note, magazine string) string {
	return fmt.Sprintf("ransome note '%s' from magazine '%s'", note, magazine)
}
