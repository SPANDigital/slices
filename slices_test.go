package slices

import (
	"context"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
	"testing"
)

type firstArgKey struct{}
type resultKey struct{}

func tableToIntegerSlice(table *godog.Table) ([]int, error) {
	var out []int
	for _, row := range table.Rows {
		for _, cell := range row.Cells {
			i, err := strconv.Atoi(cell.Value)
			if err != nil {
				return nil, err
			}
			out = append(out, i)
		}
	}
	return out, nil
}

func anIntegerSliceWithElements(ctx context.Context, elements *godog.Table) (context.Context, error) {
	sl, err := tableToIntegerSlice(elements)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, firstArgKey{}, sl), nil
}

func theResultShouldBeAnIntegerSlice(ctx context.Context, expected *godog.Table) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]int)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	expectedSlice, err := tableToIntegerSlice(expected)
	if err != nil {
		return ctx, err
	}
	if len(result) != len(expectedSlice) {
		return ctx, errors.New("slices have different lengths")
	}
	for i, v := range expectedSlice {
		if result[i] != v {
			return ctx, fmt.Errorf("slices differ at index %d", i)
		}
	}
	return ctx, nil
}

func theResultShouldBeStringSlice(ctx context.Context, table *godog.Table) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]string)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	expected := tableToStringSlice(table)
	if len(result) != len(expected) {
		return ctx, errors.New("result length does not match expected length")
	}
	for i := range result {
		if result[i] != expected[i] {
			return ctx, fmt.Errorf("value mismatch at index %d expected %s got %s", i, expected[i], result[i])
		}
	}
	return ctx, nil
}

func anEmptyStringPtrSlice(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, []*string{}), nil
}

func aStringPtrSliceWithElements(ctx context.Context, elements *godog.Table) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, tableToStringPtrSlice(elements)), nil
}

func theResultShouldBeAnEmptyStringPtrSlice(ctx context.Context) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]*string)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if len(result) != 0 {
		return ctx, errors.New("result is not empty")
	}
	return ctx, nil
}
func aStringSliceWithElements(ctx context.Context, elements *godog.Table) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, tableToStringSlice(elements)), nil
}

func tableToSliceOfStringSlices(table *godog.Table) ([][]string, error) {
	var out [][]string
	for _, row := range table.Rows {
		var inner []string
		for _, cell := range row.Cells {
			inner = append(inner, cell.Value)
		}
		out = append(out, inner)
	}
	return out, nil
}

func trimRight(strs []string) []string {
	largest := 0
	for i := range len(strs) {
		if strs[i] != "" {
			largest = i + 1
		}
	}
	return strs[:largest]
}

func tableToSliceOfStringSlicesTrimCellsOnRight(table *godog.Table) ([][]string, error) {
	out, err := tableToSliceOfStringSlices(table)
	if err != nil {
		return nil, err
	}
	for i := range out {
		out[i] = trimRight(out[i])
	}
	return out, nil
}

func theResultShouldBeASliceOfStringSlicesWithElements(ctx context.Context, expected *godog.Table) (context.Context, error) {
	e, err := tableToSliceOfStringSlices(expected)
	if err != nil {
		return ctx, err
	}
	result, ok := ctx.Value(resultKey{}).([][]string)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if len(result) != len(e) {
		return ctx, errors.New("result length does not match expected length")
	}
	for i := range result {
		if len(result[i]) != len(e[i]) {
			return ctx, errors.New("inner slice length does not match expected length")
		}
		for j := range result[i] {
			if result[i][j] != e[i][j] {
				return ctx, errors.New("value mismatch")
			}
		}
	}
	return ctx, nil
}

func theResultShouldBeASliceOfStringSlicesWithElementsTrimCellsOnRight(ctx context.Context, expected *godog.Table) (context.Context, error) {
	e, err := tableToSliceOfStringSlicesTrimCellsOnRight(expected)
	if err != nil {
		return ctx, err
	}
	result, ok := ctx.Value(resultKey{}).([][]string)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if len(result) != len(e) {
		return ctx, errors.New("result length does not match expected length")
	}
	for i := range result {
		if len(result[i]) != len(e[i]) {
			return ctx, errors.New("inner slice length does not match expected length")
		}
		for j := range result[i] {
			if result[i][j] != e[i][j] {
				return ctx, errors.New("value mismatch")
			}
		}
	}
	return ctx, nil
}

func theResultShouldBeInt(ctx context.Context, expected int) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).(int)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if result != expected {
		return ctx, fmt.Errorf("expected %d, got %d", expected, result)
	}
	return ctx, nil
}

func theResultShouldBeBool(ctx context.Context, expected string) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).(bool)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if expected == "true" && !result {
		return ctx, errors.New("expected true, got false")
	}
	if expected == "false" && result {
		return ctx, errors.New("expected false, got true")
	}
	return ctx, nil
}

func convertTableToSliceOfVerticalStringSlices(table *godog.Table) ([][]string, error) {
	var out [][]string
	for _, row := range table.Rows {
		for i, cell := range row.Cells {
			if i >= len(out) {
				out = append(out, []string{})
			}
			out[i] = append(out[i], cell.Value)
		}
	}
	return out, nil
}

func theFollowingStringSlices(ctx context.Context, table *godog.Table) (context.Context, error) {
	s, err := convertTableToSliceOfVerticalStringSlices(table)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, firstArgKey{}, s), nil
}

func iHaveEmptyStringSlices(ctx context.Context, quantity int) (context.Context, error) {
	return context.WithValue(ctx, firstArgKey{}, make([][]string, quantity)), nil
}

func theResultShouldBeAnEmptyStringSlice(ctx context.Context) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).([]string)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if len(result) != 0 {
		return ctx, errors.New("result is not empty")
	}
	return ctx, nil
}

func initializeSharedSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^a string slice with elements$`, aStringSliceWithElements)
	ctx.Step(`^a string slice$`, aStringSliceWithElements)
	ctx.Step(`^a slice of strings with elements$`, aStringSliceWithElements)
	ctx.Step(`^an empty string ptr slice$`, anEmptyStringPtrSlice)
	ctx.Step(`^a string ptr slice with elements$`, aStringPtrSliceWithElements)
	ctx.Step(`^the result should be an empty string ptr slice$`, theResultShouldBeAnEmptyStringPtrSlice)
	ctx.Step(`^the result should be a slice of string slices with elements$`, theResultShouldBeASliceOfStringSlicesWithElements)
	ctx.Step(`^the result should be a slice of string slices with elements trim cells on right$`, theResultShouldBeASliceOfStringSlicesWithElementsTrimCellsOnRight)
	ctx.Step(`^the result should be (\d+)$`, theResultShouldBeInt)
	ctx.Step(`^the result should be boolean (\w+)$`, theResultShouldBeBool)
	ctx.Step(`^the following string slices$`, theFollowingStringSlices)
	ctx.Step(`^the result should be string slice$`, theResultShouldBeStringSlice)
	ctx.Step(`^I have (\d+) empty string slices$`, iHaveEmptyStringSlices)
	ctx.Step(`^the result should be an empty string slice$`, theResultShouldBeAnEmptyStringSlice)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	initializeSharedSteps(ctx)
	initializeScenarioForAppendNotNil(ctx)
	initializeScenarioForContains(ctx)
	initializeScenarioForFilter(ctx)
	initializeScenarioForFilterNil(ctx)
	initializeScenarioForFlatten(ctx)
	initializeScenarioForGroupByLen(ctx)
	initializeScenarioForIndex(ctx)
	initializeScenarioForIntersection(ctx)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
