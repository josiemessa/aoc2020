package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "byr valid:   2002",
			args: args{field: "byr:2002"},
			want: true,
		},
		{
			name: "byr, invalid: 2003",
			args: args{field: "byr:2003"},
			want: false,
		},
		{
			name: "hgt, valid: 60in",
			args: args{field: "hgt:60in"},
			want: true,
		},
		{
			name: "hgt, valid: 190cm",
			args: args{field: "hgt:190cm"},
			want: true,
		},
		{
			name: "hgt, invalid: 190in",
			args: args{field: "hgt:190in"},
			want: false,
		},
		{
			name: "hgt, invalid: 190",
			args: args{field: "hgt:190"},
			want: false},

		{
			name: "hcl, valid:   #123abc",
			args: args{field: "hcl:#123abc"},
			want: true,
		},
		{
			name: "hcl, invalid: #123abz",
			args: args{field: "hcl:#123abz"},
			want: false,
		},
		{
			name: "hcl, invalid: 123abc",
			args: args{field: "hcl:123abc"},
			want: false},

		{
			name: "ecl, valid: brn",
			args: args{field: "ecl:brn"},
			want: true,
		},
		{
			name: "ecl, invalid: wat",
			args: args{field: "ecl:wat"},
			want: false,
		},
		{
			name: "pid, valid: 000000001",
			args: args{field: "pid:000000001"},
			want: true,
		},
		{
			name: "pid, invalid: 0123456789}",
			args: args{field: "pid:0123456789}"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := strings.Split(tt.args.field, ":")
			require.Equal(t, tt.want, validators[kv[0]](kv[1]))
		})
	}
}
