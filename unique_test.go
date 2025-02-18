package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	ss "slices"
)

func theUniqueFunctionIsCalledAndResultSortedAlphabetically(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	r := Unique(firstArg)
	ss.Sort(r)
	return context.WithValue(ctx, resultKey{}, r), nil
}

func initializeScenarioForUnique(ctx *godog.ScenarioContext) {
	ctx.Step(`^the Unique function is called and result sorted alphabetically$`, theUniqueFunctionIsCalledAndResultSortedAlphabetically)
}
