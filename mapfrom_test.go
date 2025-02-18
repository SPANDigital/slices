package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iCallMapFromToConvertTheIntegersToTheirDoubleAndTripleValues(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]int)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	return context.WithValue(ctx, resultKey{}, MapFrom(firstArg, func(i int) int {
		return i * 2
	}, func(i int) int {
		return i * 3
	})), nil
}

// InitializeScenarioForMapFrom initializes the context and steps for the mapfrom feature
func initializeScenarioForMapFrom(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call MapFrom to convert the integers to their double and triple values$`, iCallMapFromToConvertTheIntegersToTheirDoubleAndTripleValues)
}
