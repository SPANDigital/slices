package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

type pageSizeKey struct{}

func aPageSizeOf(ctx context.Context, pageSize int) (context.Context, error) {
	return context.WithValue(ctx, pageSizeKey{}, pageSize), nil
}

func calculateTheNumberOfPages(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]int)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	pageSize, ok := ctx.Value(pageSizeKey{}).(int)
	if !ok {
		return ctx, errors.New("page size not found in context")
	}
	return context.WithValue(ctx, resultKey{}, NumPages(firstArg, pageSize)), nil
}

func initializeScenarioForNumPages(ctx *godog.ScenarioContext) {
	ctx.Step(`^a page size of (\d+)$`, aPageSizeOf)
	ctx.Step(`^I calculate the number of pages$`, calculateTheNumberOfPages)
}
