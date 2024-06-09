package slices

import (
	"reflect"
	"testing"
)

type testStruct struct{}

func TestRemoveNil(t *testing.T) {
	type args[T any] struct {
		in []T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantOut []T
	}
	var na, nb, nc *testStruct
	a, b, c := &testStruct{}, &testStruct{}, &testStruct{}
	tests := []testCase[*testStruct]{
		{
			name: "scattered",
			args: args[*testStruct]{
				in: []*testStruct{a, na, b, nb, c, nc},
			},
			wantOut: []*testStruct{a, b, c},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := RemoveNil(tt.args.in); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("RemoveNil() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
