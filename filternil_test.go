package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func whenICallFilterNilOnTheStringPtrSlice(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]*string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	result := FilterNil(firstArg)
	return context.WithValue(ctx, resultKey{}, result), nil
}

func initializeScenarioForFilterNil(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call FilterNil on the string ptr slice$`, whenICallFilterNilOnTheStringPtrSlice)
}
