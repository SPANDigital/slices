package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"strings"
)

func iCallMapOnTheStringsSliceWithTheFunctionThatReturnsTheLengthOfTheString(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	return context.WithValue(ctx, resultKey{}, Map(firstArg, func(s string) int { return len(s) })), nil
}

func iCallMapOnTheStringSliceWithAFunctionThatAppendsSuffixToEachElement(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	return context.WithValue(ctx, resultKey{}, Map(firstArg, func(s string) string { return s + "-suffix" })), nil
}

func iCallMapOnTheStringSliceWithAFunctionThatConvertsEachElementToUppercase(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	return context.WithValue(ctx, resultKey{}, Map(firstArg, func(s string) string { return strings.ToUpper(s) })), nil
}

func initializeScenarioForMap(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call Map on the string slice with a function that appends "\-suffix" to each element$`, iCallMapOnTheStringSliceWithAFunctionThatAppendsSuffixToEachElement)
	ctx.Step(`^I call Map on the string slice with a function that returns the length of each element$`, iCallMapOnTheStringsSliceWithTheFunctionThatReturnsTheLengthOfTheString)
	ctx.Step(`^I call Map on the string slice with a function that converts each element to uppercase$`, iCallMapOnTheStringSliceWithAFunctionThatConvertsEachElementToUppercase)
}
