package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

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

func initializeScenarioForFilter(ctx *godog.ScenarioContext) {
	ctx.Step(`^an integer slice with elements$`, anIntegerSliceWithElements)
	ctx.Step(`^I call Filter with a function that returns true for positive integers$`, iCallFilterWithAFunctionThatReturnsTrueForPositiveIntegers)
	ctx.Step(`^I call Filter with a function that returns true for negative integers$`, iCallFilterWithAFunctionThatReturnsTrueForNegativeIntegers)
	ctx.Step(`^the result should be an integer slice with elements$`, theResultShouldBeAnIntegerSlice)
}
