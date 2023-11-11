package strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		str  string
		want string
		desc string
	}{
		{
			str:  "",
			want: "",
		},
		{
			str:  "12",
			want: "21",
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%d %v", i+1, tC.desc), func(t *testing.T) {
			var temp = []byte(tC.str)
			reverseString(temp)

			if !reflect.DeepEqual([]byte(tC.want), temp) {
				t.Errorf("expected %v but %v", tC.want, temp)
			}
		})
	}
}
