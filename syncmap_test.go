package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"strings"
)

func iCallSyncMapOnTheStringsSliceWithTheFunctionThatReturnsTheLengthOfTheString(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	v, err := SyncMap[[]string, string, int](ctx, firstArg, func(ctx context.Context, s string) (int, error) { return len(s), nil })
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, resultKey{}, v), nil
}

func iCallSyncMapOnTheStringSliceWithAFunctionThatAppendsSuffixToEachElement(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	v, err := SyncMap[[]string, string, string](ctx, firstArg, func(ctx context.Context, s string) (string, error) { return s + "-suffix", nil })
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, resultKey{}, v), nil
}

func iCallSyncMapOnTheStringSliceWithAFunctionThatConvertsEachElementToUppercase(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	v, err := SyncMap[[]string, string, string](ctx, firstArg, func(ctx context.Context, s string) (string, error) { return strings.ToUpper(s), nil })
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, resultKey{}, v), nil
}

func initializeScenarioForSyncMap(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call SyncMap on the string slice with a function that appends "\-suffix" to each element$`, iCallSyncMapOnTheStringSliceWithAFunctionThatAppendsSuffixToEachElement)
	ctx.Step(`^I call SyncMap on the string slice with a function that returns the length of each element$`, iCallSyncMapOnTheStringsSliceWithTheFunctionThatReturnsTheLengthOfTheString)
	ctx.Step(`^I call SyncMap on the string slice with a function that converts each element to uppercase$`, iCallSyncMapOnTheStringSliceWithAFunctionThatConvertsEachElementToUppercase)
}
