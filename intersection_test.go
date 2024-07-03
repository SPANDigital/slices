package slices

import (
	"cmp"
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	type args[T cmp.Ordered] struct {
		pS [][]T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "empty",
			args: args[string]{
				pS: make([][]string, 0),
			},
			want: make([]string, 0),
		},
		{
			name: "none-in-commons",
			args: args[string]{
				pS: [][]string{{"one", "two"}, {"three", "four"}},
			},
			want: make([]string, 0),
		},
		{
			name: "one-in-commons",
			args: args[string]{
				pS: [][]string{{"one", "two", "common"}, {"common", "three", "four"}},
			},
			want: []string{"common"},
		},
		{
			name: "two-in-commons",
			args: args[string]{
				pS: [][]string{{"one", "two", "common", "second-common"}, {"common", "second-common", "three", "four"}},
			},
			want: []string{"common", "second-common"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.pS...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
