package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc"}

func TestDay2_Part1(t *testing.T) {
	s := &PasswordSolver{Input: input}
	require.Equal(t, 2, s.Solve(true))
}

func TestDay2_Part2(t *testing.T) {
	s := &PasswordSolver{Input: input}
	require.Equal(t, 1, s.Solve(false))
}

func Test_ParseLine_Part1(t *testing.T) {
	type args struct {
		line string
		old  bool
	}
	tests := []struct {
		name         string
		args         args
		wantPolicy   *oldPolicy
		wantPassword string
	}{
		{
			name: "part1: first line",
			args: args{line: "1-3 a: abcde", old: true},
			wantPolicy: &oldPolicy{
				frequencyMin: 1,
				frequencyMax: 3,
				character:    "a",
			},
			wantPassword: "abcde",
		},
		{
			name: "part1: second line",
			args: args{line: "1-3 b: cdefg", old: true},
			wantPolicy: &oldPolicy{
				frequencyMin: 1,
				frequencyMax: 3,
				character:    "b",
			},
			wantPassword: "cdefg",
		},
		{
			name: "part1: third line",
			args: args{line: "2-9 c: ccccccccc", old: true},
			wantPolicy: &oldPolicy{
				frequencyMin: 2,
				frequencyMax: 9,
				character:    "c",
			},
			wantPassword: "ccccccccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPolicy, gotPassword := parseLine(tt.args.line, tt.args.old)
			require.EqualValues(t, tt.wantPolicy, gotPolicy)
			require.Equal(t, tt.wantPassword, gotPassword)
		})
	}
}

func Test_ParseLine_Part2(t *testing.T) {
	type args struct {
		line string
		old  bool
	}
	tests := []struct {
		name         string
		args         args
		wantPolicy   *newPolicy
		wantPassword string
	}{
		{
			name: "part1: first line",
			args: args{line: "1-3 a: abcde", old: false},
			wantPolicy: &newPolicy{
				pos1:      1,
				pos2:      3,
				character: "a",
			},
			wantPassword: "abcde",
		},
		{
			name: "part1: second line",
			args: args{line: "1-3 b: cdefg", old: false},
			wantPolicy: &newPolicy{
				pos1:      1,
				pos2:      3,
				character: "b",
			},
			wantPassword: "cdefg",
		},
		{
			name: "part1: third line",
			args: args{line: "2-9 c: ccccccccc", old: false},
			wantPolicy: &newPolicy{
				pos1:      2,
				pos2:      9,
				character: "c",
			},
			wantPassword: "ccccccccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPolicy, gotPassword := parseLine(tt.args.line, tt.args.old)
			require.EqualValues(t, tt.wantPolicy, gotPolicy)
			require.Equal(t, tt.wantPassword, gotPassword)
		})
	}
}

func Test_oldPolicy_apply(t *testing.T) {
	type fields struct {
		frequencyMin int
		frequencyMax int
		character    string
	}
	type args struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "first line",
			fields: fields{
				frequencyMin: 1,
				frequencyMax: 3,
				character:    "a",
			},
			args: args{password: "abcde"},
			want: true,
		},
		{
			name: "second line",
			fields: fields{
				frequencyMin: 1,
				frequencyMax: 3,
				character:    "b",
			},
			args: args{password: "cdefg"},
			want: false,
		},
		{
			name: "third line",
			fields: fields{
				frequencyMin: 2,
				frequencyMax: 9,
				character:    "c",
			},
			args: args{"ccccccccc"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &oldPolicy{
				frequencyMin: tt.fields.frequencyMin,
				frequencyMax: tt.fields.frequencyMax,
				character:    tt.fields.character,
			}
			require.Equal(t, tt.want, p.apply(tt.args.password))
		})
	}
}

func Test_newPolicy_apply(t *testing.T) {
	type fields struct {
		pos1      int
		pos2      int
		character string
	}
	type args struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "first line",
			fields: fields{
				pos1:      1,
				pos2:      3,
				character: "a",
			},
			args: args{password: "abcde"},
			want: true,
		},
		{
			name: "second line",
			fields: fields{
				pos1:      1,
				pos2:      3,
				character: "b",
			},
			args: args{password: "cdefg"},
			want: false,
		},
		{
			name: "third line",
			fields: fields{
				pos1:      2,
				pos2:      9,
				character: "c",
			},
			args: args{"ccccccccc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &newPolicy{
				pos1:      tt.fields.pos1,
				pos2:      tt.fields.pos2,
				character: tt.fields.character,
			}
			require.Equal(t, tt.want, p.apply(tt.args.password))
		})
	}
}
