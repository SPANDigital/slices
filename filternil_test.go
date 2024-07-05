package slices

import (
	"reflect"
	"testing"
)

func TestFilterNil(t *testing.T) {
	type args[T any] struct {
		in []*T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantOut []*T
	}
	tests := []testCase[string]{
		{
			name: "empty",
			args: args[string]{
				in: nil,
			},
			wantOut: make([]*string, 0),
		},
		{
			name: "single-front",
			args: args[string]{
				in: []*string{nil, ptr("a")},
			},
			wantOut: []*string{ptr("a")},
		},
		{
			name: "single-back",
			args: args[string]{
				in: []*string{ptr("a"), nil},
			},
			wantOut: []*string{ptr("a")},
		},
		{
			name: "middle",
			args: args[string]{
				in: []*string{ptr("a"), nil, ptr("b")},
			},
			wantOut: []*string{ptr("a"), ptr("b")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FilterNil(tt.args.in); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("FilterNil() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
