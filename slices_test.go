package slices

import (
	"github.com/cucumber/godog"
	"testing"
)

type firstArgKey struct{}
type resultKey struct{}

func InitializeScenario(ctx *godog.ScenarioContext) {
	initializeScenarioForAppendNotNil(ctx)
	initializeScenarioForContains(ctx)
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
