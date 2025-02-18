package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iCallGroupByLenOnTheStringsSliceWithLength(ctx context.Context, length int) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	return context.WithValue(ctx, resultKey{}, GroupByLen(firstArg, length)), nil
}

func initializeScenarioForGroupByLen(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call GroupByLen on the strings slice with length (\d+)$`, iCallGroupByLenOnTheStringsSliceWithLength)
}
