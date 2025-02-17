package slices

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"regexp"
)

var naiveStringPattern = regexp.MustCompile(`"([^"]*)"`)

func tableToStringPtrSlice(t *godog.Table) []*string {
	var out []*string
	for _, row := range t.Rows {
		for _, cell := range row.Cells {
			switch value := cell.Value; {
			case value == "nil":
				out = append(out, nil)
			case naiveStringPattern.MatchString(value):
				m := naiveStringPattern.FindStringSubmatch(value)
				out = append(out, &m[1])
			default:
				out = append(out, &value)
			}
		}
	}
	return out
}

func anEmptySlice(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, []*string{}), nil
}

func aStringPtrSliceWithElements(ctx context.Context, elements *godog.Table) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, tableToStringPtrSlice(elements)), nil
}

func iCallAppendNotNilOnTheSliceWithTheFollowingVariableArguments(ctx context.Context, varArgs *godog.Table) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]*string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	result := AppendNotNil(firstArg, tableToStringPtrSlice(varArgs)...)
	return context.WithValue(ctx, resultKey{}, result), nil
}

func iCallAppendNotNilOnTheSlice(ctx context.Context) (context.Context, error) {
	firstArg, ok := ctx.Value(firstArgKey{}).([]*string)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	result := AppendNotNil(firstArg)
	return context.WithValue(ctx, resultKey{}, result), nil
}

func theResultShouldBeASliceWithElements(ctx context.Context, t *godog.Table) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]*string)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	expected := tableToStringPtrSlice(t)
	if len(result) != len(expected) {
		return ctx, errors.New("result length does not match expected length")
	}
	for i := range result {
		if (result[i] == nil && expected[i] != nil) || (result[i] != nil && expected[i] == nil) {
			return ctx, errors.New("nil mismatch")
		}
		if result[i] != nil && *result[i] != *expected[i] {
			return ctx, errors.New("value mismatch")
		}
	}
	return ctx, nil
}

func theResultShouldBeAnEmptySlice(ctx context.Context) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]*string)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if len(result) != 0 {
		return ctx, errors.New("result is not empty")
	}
	return ctx, nil
}

func initializeScenarioForAppendNotNil(ctx *godog.ScenarioContext) {
	ctx.Step(`^an empty slice$`, anEmptySlice)
	ctx.Step(`^a string ptr slice with elements$`, aStringPtrSliceWithElements)
	ctx.Step(`^I call AppendNotNil on the slice with the following variable arguments$`, iCallAppendNotNilOnTheSliceWithTheFollowingVariableArguments)
	ctx.Step(`^I call AppendNotNil on the slice$`, iCallAppendNotNilOnTheSlice)
	ctx.Step(`^the result should be a slice with elements$`, theResultShouldBeASliceWithElements)
	ctx.Step(`^the result should be an empty slice$`, theResultShouldBeAnEmptySlice)
}
