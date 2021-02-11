package myreg

import (
	"testing"
)

func Test_match(t *testing.T) {
	type args struct {
		pattern string
		text    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "matchOne",
			args: args{
				pattern: "a",
				text:    "a",
			},
			want: true,
		},
		{
			name: "empty",
			args: args{
				pattern: "",
				text:    "",
			},
			want: true,
		},
		{
			name: "not match one",
			args: args{
				pattern: "a",
				text:    "b",
			},
			want: false,
		},
		{
			name: "match one any caracter",
			args: args{
				pattern: ".",
				text:    "b",
			},
			want: true,
		},
		{
			name: "match empty",
			args: args{
				pattern: "",
				text:    "abc",
			},
			want: true,
		},
		{
			name: "match empty pattern $",
			args: args{
				pattern: "$",
				text:    "",
			},
			want: true,
		},
		{
			name: "match pattern $",
			args: args{
				pattern: "$",
				text:    "abc",
			},
			want: false,
		},
		{
			name: "match string",
			args: args{
				pattern: "abc",
				text:    "abc",
			},
			want: true,
		},
		{
			name: "not match string",
			args: args{
				pattern: "abc",
				text:    "bca",
			},
			want: false,
		},
		{
			name: "match any string",
			args: args{
				pattern: "a.c",
				text:    "abc",
			},
			want: true,
		},
		{
			name: "match any string fail",
			args: args{
				pattern: "a.b",
				text:    "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.args.pattern, tt.args.text); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search(t *testing.T) {
	type args struct {
		pattern string
		text    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "search ^",
			args: args{
				pattern: "^abc",
				text:    "abc",
			},
			want: true,
		},
		{
			name: "search ^ fail",
			args: args{
				pattern: "^abc",
				text:    "bca",
			},
			want: false,
		},
		{
			name: "search ^ $",
			args: args{
				pattern: "^abc$",
				text:    "abc",
			},
			want: true,
		},
		{
			name: "search ^ $ fail",
			args: args{
				pattern: "^abc$",
				text:    "bca",
			},
			want: false,
		},
		{
			name: "search ^ more ",
			args: args{
				pattern: "^a",
				text:    "abc",
			},
			want: true,
		},
		{
			name: "search ^ more fail",
			args: args{
				pattern: "^a",
				text:    "bca",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.pattern, tt.args.text); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
