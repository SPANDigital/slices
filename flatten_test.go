package slices

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	type args[T any] struct {
		s [][]T
	}
	type testCase[T any] struct {
		name          string
		args          args[T]
		wantFlattened []T
	}
	tests := []testCase[int]{
		{
			name: "flatten",
			args: args[int]{
				s: [][]int{{}, {1}, {2, 3}, {4, 5, 6}},
			},
			wantFlattened: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFlattened := Flatten(tt.args.s); !reflect.DeepEqual(gotFlattened, tt.wantFlattened) {
				t.Errorf("Flatten() = %v, want %v", gotFlattened, tt.wantFlattened)
			}
		})
	}
}

func TestFlatten3(t *testing.T) {
	type args[T any] struct {
		s [][][]T
	}
	type testCase[T any] struct {
		name          string
		args          args[T]
		wantFlattened []T
	}
	tests := []testCase[int]{
		{
			name: "flatten3",
			args: args[int]{
				s: [][][]int{{{1}, {2, 3}}, {{4, 5, 6}}},
			},
			wantFlattened: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFlattened := Flatten3(tt.args.s); !reflect.DeepEqual(gotFlattened, tt.wantFlattened) {
				t.Errorf("Flatten3() = %v, want %v", gotFlattened, tt.wantFlattened)
			}
		})
	}
}

func TestFlatten4(t *testing.T) {
	type args[T any] struct {
		s [][][][]T
	}
	type testCase[T any] struct {
		name          string
		args          args[T]
		wantFlattened []T
	}
	tests := []testCase[int]{
		{
			name: "flatten4",
			args: args[int]{
				s: [][][][]int{{{{1}}, {{2}, {3}}}, {{{4, 5, 6}}}},
			},
			wantFlattened: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFlattened := Flatten4(tt.args.s); !reflect.DeepEqual(gotFlattened, tt.wantFlattened) {
				t.Errorf("Flatten4() = %v, want %v", gotFlattened, tt.wantFlattened)
			}
		})
	}
}
