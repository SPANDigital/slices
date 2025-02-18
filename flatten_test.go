package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"strconv"
)

func tableToSliceOfIntegerSlices(t *godog.Table) ([][]int, error) {
	var out [][]int
	for _, row := range t.Rows {
		var inner []int
		for _, cell := range row.Cells {
			if cell.Value != "" {
				i, err := strconv.Atoi(cell.Value)
				if err != nil {
					return nil, err
				}
				inner = append(inner, i)
			}
		}
		out = append(out, inner)
	}
	return out, nil
}

func aSliceOfIntegerSlicesWithElements(ctx context.Context, elements *godog.Table) (context.Context, error) {
	s, err := tableToSliceOfIntegerSlices(elements)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, firstArgKey{}, s), nil
}

func iCallFlattenOnTheSliceOfIntegerSlices(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([][]int)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	result := Flatten(firstArg)
	return context.WithValue(ctx, resultKey{}, result), nil
}

func theResultShouldBeAnIntegerSliceWithElements(ctx context.Context, table *godog.Table) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]int)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	expected, err := tableToIntegerSlice(table)
	if err != nil {
		return ctx, err
	}
	if len(result) != len(expected) {
		return ctx, errors.New("result length does not match expected length")
	}
	for i := range result {
		if result[i] != expected[i] {
			return ctx, errors.New("mismatch")
		}
	}
	return ctx, nil
}

func theResultShouldBeAnEmptyIntegerSlice(ctx context.Context) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]int)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if len(result) != 0 {
		return ctx, errors.New("result is not empty")
	}
	return ctx, nil
}

func anEmptySliceOfIntegerSlices(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, [][]int{}), nil
}

func initializeScenarioForFlatten(ctx *godog.ScenarioContext) {
	ctx.Step(`^a slice of integer slices with elements$`, aSliceOfIntegerSlicesWithElements)
	ctx.Step(`^I call Flatten on the slice of integer slices$`, iCallFlattenOnTheSliceOfIntegerSlices)
	ctx.Step(`^the result should be an integer slice with elements$`, theResultShouldBeAnIntegerSliceWithElements)
	ctx.Step(`^the result should be an empty integer slice$`, theResultShouldBeAnEmptyIntegerSlice)
	ctx.Step(`^an empty slice of integer slices$`, anEmptySliceOfIntegerSlices)
}
