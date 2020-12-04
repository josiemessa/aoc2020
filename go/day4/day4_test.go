package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestDay4_Part1(t *testing.T) {
	require.Equal(t, 2, SolveP1(strings.NewReader(input)))
}

func TestDay4_Part2_AllInvalid(t *testing.T) {
	var input = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`
	require.Equal(t, 0, SolveP2(strings.NewReader(input)))

}

func TestDay4_Part2_AllValid(t *testing.T) {
	var input = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
	require.Equal(t, 4, SolveP2(strings.NewReader(input)))

}

func TestPassport_Parse(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "byr valid:   2002",
			args:    args{field: "byr:2002"},
			wantErr: false,
		},
		{
			name:    "byr, invalid: 2003",
			args:    args{field: "byr:2003"},
			wantErr: true,
		},
		{
			name:    "hgt, valid: 60in",
			args:    args{field: "hgt:60in"},
			wantErr: false,
		},
		{
			name:    "hgt, valid: 190cm",
			args:    args{field: "hgt:190cm"},
			wantErr: false,
		},
		{
			name:    "hgt, invalid: 190in",
			args:    args{field: "hgt:190in"},
			wantErr: true,
		},
		{
			name:    "hgt, invalid: 190",
			args:    args{field: "hgt:190"},
			wantErr: true},

		{
			name:    "hcl, valid:   #123abc",
			args:    args{field: "hcl:#123abc"},
			wantErr: false,
		},
		{
			name:    "hcl, invalid: #123abz",
			args:    args{field: "hcl:#123abz"},
			wantErr: true,
		},
		{
			name:    "hcl, invalid: 123abc",
			args:    args{field: "hcl:123abc"},
			wantErr: true},

		{
			name:    "ecl, valid: brn",
			args:    args{field: "ecl:brn"},
			wantErr: false,
		},
		{
			name:    "ecl, invalid: wat",
			args:    args{field: "ecl:wat"},
			wantErr: true,
		},
		{
			name:    "pid, valid: 000000001",
			args:    args{field: "pid:000000001"},
			wantErr: false,
		},
		{
			name:    "pid, invalid: 0123456789}",
			args:    args{field: "pid:0123456789}"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{}
			err := p.Parse(tt.args.field)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
