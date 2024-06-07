package slices

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	type args[S interface{ ~[]V }, V any, K comparable] struct {
		s       S
		extract func(v V) K
	}
	type testCase[S interface{ ~[]V }, V any, K comparable] struct {
		name string
		args args[S, V, K]
		want map[K][]V
	}
	tests := []testCase[[]int, int, int]{
		{
			name: "test",
			args: args[[]int, int, int]{
				s: []int{-3, -2, -1, 0, 1, 2, 3},
				extract: func(v int) int {
					switch {
					case v < 0:
						return -1
					case v == 0:
						return 0
					case v > 0:
						return 1
					default:
						panic("unreachable")
					}
				},
			},
			want: map[int][]int{
				-1: {-3, -2, -1},
				0:  {0},
				1:  {1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.s, tt.args.extract); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapFrom(t *testing.T) {
	type args[S interface{ ~[]I }, I any, K comparable, V any] struct {
		s            S
		extractKey   func(i I) K
		extractValue func(i I) V
	}
	type testCase[S interface{ ~[]I }, I any, K comparable, V any] struct {
		name string
		args args[S, I, K, V]
		want map[K]V
	}
	tests := []testCase[[]int, int, int, int]{
		{
			name: "double-vs-triple",
			args: args[[]int, int, int, int]{
				s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				extractKey: func(i int) int {
					return i * 2
				},
				extractValue: func(i int) int {
					return i * 3
				},
			},
			want: map[int]int{
				2:  3,
				4:  6,
				6:  9,
				8:  12,
				10: 15,
				12: 18,
				14: 21,
				16: 24,
				18: 27,
				20: 30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapFrom(tt.args.s, tt.args.extractKey, tt.args.extractValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
