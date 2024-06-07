package slices

import "testing"

func TestContains(t *testing.T) {
	type args[S interface{ ~[]E }, E comparable] struct {
		s S
		v E
	}
	type testCase[S interface{ ~[]E }, E comparable] struct {
		name string
		args args[S, E]
		want bool
	}
	tests := []testCase[[]string, string]{
		{
			name: "present-at-0",
			args: args[[]string, string]{
				s: []string{"a", "b", "c"},
				v: "a",
			},
			want: true,
		},
		{
			name: "present-at-1",
			args: args[[]string, string]{
				s: []string{"a", "b", "c"},
				v: "b",
			},
			want: true,
		},
		{
			name: "present-at-2",
			args: args[[]string, string]{
				s: []string{"a", "b", "c"},
				v: "c",
			},
			want: true,
		},
		{
			name: "not-present",
			args: args[[]string, string]{
				s: []string{"a", "b", "c"},
				v: "z",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
