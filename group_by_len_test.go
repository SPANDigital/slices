package slices

import (
	"reflect"
	"testing"
)

func TestGroupByLen(t *testing.T) {
	type args[S interface{ ~[]V }, V any] struct {
		input  S
		length int
	}
	type testCase[S interface{ ~[]V }, V any] struct {
		name       string
		args       args[S, V]
		wantOutput [][]V
	}
	tests := []testCase[[]string, string]{
		{
			name: "pairs",
			args: args[[]string, string]{
				input:  []string{"a", "b", "c", "d"},
				length: 2,
			},
			wantOutput: [][]string{{"a", "b"}, {"c", "d"}},
		},
		{
			name: "triples",
			args: args[[]string, string]{
				input:  []string{"a", "b", "c", "d", "e", "f"},
				length: 3,
			},
			wantOutput: [][]string{{"a", "b", "c"}, {"d", "e", "f"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := GroupByLen(tt.args.input, tt.args.length); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("GroupByLen() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
