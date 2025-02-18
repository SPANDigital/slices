package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

type pageIndexKey struct{}

func aPageIndexOf(ctx context.Context, pageIndex int) (context.Context, error) {
	return context.WithValue(ctx, pageIndexKey{}, pageIndex), nil
}

func iRetrieveThePage(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]int)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	pageSize, ok := ctx.Value(pageSizeKey{}).(int)
	if !ok {
		return ctx, errors.New("page size not found in context")
	}
	pageIndex, ok := ctx.Value(pageIndexKey{}).(int)
	if !ok {
		return ctx, errors.New("page index not found in context")
	}
	return context.WithValue(ctx, resultKey{}, Page(firstArg, pageSize, pageIndex)), nil
}

func initializeScenarioForPaging(ctx *godog.ScenarioContext) {
	ctx.Step(`^a page index of (\d+)$`, aPageIndexOf)
	ctx.Step(`^I retrieve the page$`, iRetrieveThePage)
}
