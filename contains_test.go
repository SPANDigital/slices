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
func theContainsFunctionIsCalledWithStringArgument(ctx context.Context, arg string) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	result := Contains(firstArg, arg)
	return context.WithValue(ctx, resultKey{}, result), nil
}

func initializeScenarioForContains(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call Contains with argument "([^"]*)"$`, theContainsFunctionIsCalledWithStringArgument)
}
