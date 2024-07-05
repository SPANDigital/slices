package slices

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

func sort[T cmp.Ordered](s []T) []T {
	ret := s
	slices.Sort(ret)
	return ret
}

func TestUnique(t *testing.T) {
	type args[T comparable] struct {
		s []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "empty",
			args: args[string]{
				s: make([]string, 0),
			},
			want: make([]string, 0),
		},
		{
			name: "one-dup",
			args: args[string]{
				s: []string{"a", "b", "b"},
			},
			want: []string{"a", "b"},
		},
		{
			name: "three-dup",
			args: args[string]{
				s: []string{"a", "b", "b", "c", "c", "d", "d"},
			},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(Unique(tt.args.s)); !(len(got) == 0 && len(tt.want) == 0) && !reflect.DeepEqual(got, sort(tt.want)) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
