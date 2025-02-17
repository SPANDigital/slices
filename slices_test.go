package slices

import (
	"context"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
	"testing"
)

type firstArgKey struct{}
type resultKey struct{}

func tableToIntegerSlice(table *godog.Table) ([]int, error) {
	var out []int
	for _, row := range table.Rows {
		for _, cell := range row.Cells {
			i, err := strconv.Atoi(cell.Value)
			if err != nil {
				return nil, err
			}
			out = append(out, i)
		}
	}
	return out, nil
}

func anIntegerSliceWithElements(ctx context.Context, elements *godog.Table) (context.Context, error) {
	sl, err := tableToIntegerSlice(elements)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, firstArgKey{}, sl), nil
}

func theResultShouldBeAnIntegerSlice(ctx context.Context, expected *godog.Table) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]int)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	expectedSlice, err := tableToIntegerSlice(expected)
	if err != nil {
		return ctx, err
	}
	if len(result) != len(expectedSlice) {
		return ctx, errors.New("slices have different lengths")
	}
	for i, v := range expectedSlice {
		if result[i] != v {
			return ctx, fmt.Errorf("slices differ at index %d", i)
		}
	}
	return ctx, nil
}

func anEmptyStringPtrSlice(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, []*string{}), nil
}

func aStringPtrSliceWithElements(ctx context.Context, elements *godog.Table) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, tableToStringPtrSlice(elements)), nil
}

func theResultShouldBeAnEmptyStringPtrSlice(ctx context.Context) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]*string)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if len(result) != 0 {
		return ctx, errors.New("result is not empty")
	}
	return ctx, nil
}

func initializeSharedSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^an empty string ptr slice$`, anEmptyStringPtrSlice)
	ctx.Step(`^a string ptr slice with elements$`, aStringPtrSliceWithElements)
	ctx.Step(`^the result should be an empty string ptr slice$`, theResultShouldBeAnEmptyStringPtrSlice)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	initializeSharedSteps(ctx)
	initializeScenarioForAppendNotNil(ctx)
	initializeScenarioForContains(ctx)
	initializeScenarioForFilter(ctx)
	initializeScenarioForFilterNil(ctx)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
