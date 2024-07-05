package slices

import (
	"reflect"
	"testing"
)

func TestAppendNotNil(t *testing.T) {
	type args[T any] struct {
		in     []*T
		values []*T
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
				in:     []*string{ptr("a"), ptr("b")},
				values: []*string{ptr("c"), nil, ptr("d")},
			},
			wantOut: []*string{ptr("a"), ptr("b"), ptr("c"), ptr("d")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := AppendNotNil(tt.args.in, tt.args.values...); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppendNotNil() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
