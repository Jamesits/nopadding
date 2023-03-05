package analyzer_wrapper

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/fieldalignment"
	"reflect"
	"strings"
	"testing"
)

type SelfPackageReflectionStub struct{}

func selfPackagePath() string {
	// https://stackoverflow.com/a/25263604
	return reflect.TypeOf(SelfPackageReflectionStub{}).PkgPath()
}

func TestStructAlignmentAnalyzerSelf(t *testing.T) {
	// test ourselves (empty string means self package)
	assert.NotZero(t, Run(
		[]string{},
		[]*analysis.Analyzer{fieldalignment.Analyzer},
	))
}

func TestStructAlignmentAnalyzerVictims(t *testing.T) {
	// test victims
	assert.Zero(t, Run(
		[]string{strings.Join([]string{selfPackagePath(), "..", "internal", "testdata", "fieldalignment", "pass"}, "/")},
		[]*analysis.Analyzer{fieldalignment.Analyzer},
	))
	assert.NotZero(t, Run(
		[]string{strings.Join([]string{selfPackagePath(), "..", "internal", "testdata", "fieldalignment", "fail"}, "/")},
		[]*analysis.Analyzer{fieldalignment.Analyzer},
	))
}
