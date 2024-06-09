package slices

import (
	"reflect"
	"testing"
)

func TestNumPages(t *testing.T) {
	type args[V any] struct {
		s        []V
		pageSize int
	}
	type testCase[V any] struct {
		name string
		args args[V]
		want int
	}
	tests := []testCase[int]{
		{
			name: "three-pages",
			args: args[int]{
				s:        []int{1, 2, 3, 4, 5, 6},
				pageSize: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumPages(tt.args.s, tt.args.pageSize); got != tt.want {
				t.Errorf("NumPages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage(t *testing.T) {
	type args[V any] struct {
		s         []V
		pageSize  int
		pageIndex int
	}
	type testCase[V any] struct {
		name string
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "two-pages",
			args: args[int]{
				s:         []int{1, 2, 3, 4, 5, 6},
				pageSize:  3,
				pageIndex: 1,
			},
			want: []int{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Page(tt.args.s, tt.args.pageSize, tt.args.pageIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Page() = %v, want %v", got, tt.want)
			}
		})
	}
}
