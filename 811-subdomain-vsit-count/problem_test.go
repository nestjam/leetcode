package subdomainvsitcount

import (
	"reflect"
	"slices"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	type args struct {
		cpdomains []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				cpdomains: []string{"1 com"},
			},
			want: []string{"1 com"},
		},
		{
			args: args{
				cpdomains: []string{"1 mail.com"},
			},
			want: []string{"1 com", "1 mail.com"},
		},
		{
			args: args{
				cpdomains: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			},
			want: []string{"1 intel.mail.com", "5 org", "5 wiki.org", "50 yahoo.com", "900 google.mail.com", "901 mail.com", "951 com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subdomainVisits(tt.args.cpdomains)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}
