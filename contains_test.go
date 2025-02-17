package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func tableToStringSlice(table *godog.Table) []string {
	var out []string
	for _, row := range table.Rows {
		for _, cell := range row.Cells {
			out = append(out, cell.Value)
		}
	}
	return out
}

func aStringSliceWithElements(ctx context.Context, elements *godog.Table) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, tableToStringSlice(elements)), nil
}

func theContainsFunctionIsCalledWithStringArgument(ctx context.Context, arg string) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	result := Contains(firstArg, arg)
	return context.WithValue(ctx, resultKey{}, result), nil
}

func theResultShouldBeBool(ctx context.Context, expected string) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).(bool)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if expected == "true" && !result {
		return ctx, errors.New("expected true, got false")
	}
	if expected == "false" && result {
		return ctx, errors.New("expected false, got true")
	}
	return ctx, nil
}

func initializeScenarioForContains(ctx *godog.ScenarioContext) {
	ctx.Step(`^a string slice with elements$`, aStringSliceWithElements)
	ctx.Step(`^I call Contains with argument "([^"]*)"$`, theContainsFunctionIsCalledWithStringArgument)
	ctx.Step(`^the result should be (\w+)$`, theResultShouldBeBool)

}
