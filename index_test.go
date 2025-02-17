package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iCallIndexOnTheStringSliceWithArgument(ctx context.Context, argument string) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	return context.WithValue(ctx, resultKey{}, Index(firstArg, argument)), nil
}

func initializeScenarioForIndex(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call Index on the string slice with argument "([^"]*)"$`, iCallIndexOnTheStringSliceWithArgument)
}
