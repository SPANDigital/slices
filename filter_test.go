package slices

import (
	"context"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
)

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

func iCallFilterWithAFunctionThatReturnsTrueForPositiveIntegers(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]int)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	result := Filter(firstArg, func(i int) bool {
		return i > 0
	})
	return context.WithValue(ctx, resultKey{}, result), nil
}

func iCallFilterWithAFunctionThatReturnsTrueForNegativeIntegers(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]int)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	result := Filter(firstArg, func(i int) bool {
		return i < 0
	})
	return context.WithValue(ctx, resultKey{}, result), nil
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

func initializeScenarioForFilter(ctx *godog.ScenarioContext) {
	ctx.Step(`^an integer slice with elements$`, anIntegerSliceWithElements)
	ctx.Step(`^I call Filter with a function that returns true for positive integers$`, iCallFilterWithAFunctionThatReturnsTrueForPositiveIntegers)
	ctx.Step(`^I call Filter with a function that returns true for negative integers$`, iCallFilterWithAFunctionThatReturnsTrueForNegativeIntegers)
	ctx.Step(`^the result should be an integer slice with elements$`, theResultShouldBeAnIntegerSlice)
}
