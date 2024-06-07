package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[S interface{ ~[]V }, V any] struct {
		s      S
		filter func(V V) bool
	}
	type testCase[S interface{ ~[]V }, V any] struct {
		name string
		args args[S, V]
		want []V
	}
	tests := []testCase[[]int, int]{
		{
			name: "filter-positive",
			args: args[[]int, int]{
				s: []int{1, -1, 2, -2, 3, -3},
				filter: func(v int) bool {
					return v > 0
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "filter-negative",
			args: args[[]int, int]{
				s: []int{1, -1, 2, -2, 3, -3},
				filter: func(v int) bool {
					return v < 0
				},
			},
			want: []int{-1, -2, -3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.s, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
