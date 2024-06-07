package slices

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args[S interface{ ~[]V }, V any, E any] struct {
		s       S
		extract func(V V) E
	}
	type testCase[S interface{ ~[]V }, V any, E any] struct {
		name string
		args args[S, V, E]
		want []E
	}
	tests := []testCase[[]int, int, bool]{
		{
			name: "empty",
			args: args[[]int, int, bool]{
				s: []int{1, -1, 2, -2, 3, -3},
				extract: func(v int) bool {
					return v < 0
				},
			},
			want: []bool{false, true, false, true, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.s, tt.args.extract); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
