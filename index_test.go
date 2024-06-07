package slices

import "testing"

func TestIndex(t *testing.T) {
	type args[S interface{ ~[]E }, E comparable] struct {
		s S
		v E
	}
	type testCase[S interface{ ~[]E }, E comparable] struct {
		name string
		args args[S, E]
		want int
	}
	stringTests := []testCase[[]string, string]{
		{
			name: "present-at-0",
			args: args[[]string, string]{
				s: []string{"a", "b", "c"},
				v: "a",
			},
			want: 0,
		},
		{
			name: "present-at-1",
			args: args[[]string, string]{
				s: []string{"a", "b", "c"},
				v: "b",
			},
			want: 1,
		},
		{
			name: "present-at-2",
			args: args[[]string, string]{
				s: []string{"a", "b", "c"},
				v: "c",
			},
			want: 2,
		},
		{
			name: "not-present",
			args: args[[]string, string]{
				s: []string{"a", "b", "c"},
				v: "z",
			},
			want: -1,
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}
