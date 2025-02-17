package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iCallIntersectionOnTheStringSlices(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([][]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	if len(firstArg) != 2 {
		return ctx, errors.New("first arg must have two elements")
	}
	return context.WithValue(ctx, resultKey{}, Intersection(firstArg...)), nil
}

func initializeScenarioForIntersection(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call Intersection on the slices$`, iCallIntersectionOnTheStringSlices)
}
