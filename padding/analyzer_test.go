package padding

import (
	"github.com/jamesits/nopadding/analyzer_wrapper"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis"
	"reflect"
	"strings"
	"testing"
)

type SelfPackageReflectionStub struct{}

func selfPackagePath() string {
	// https://stackoverflow.com/a/25263604
	return reflect.TypeOf(SelfPackageReflectionStub{}).PkgPath()
}

func TestStructPaddingAnalyzerSelf(t *testing.T) {
	// test ourselves
	assert.Zero(t, analyzer_wrapper.Run(
		[]string{},
		[]*analysis.Analyzer{Analyzer},
	))
}

func TestStructPaddingAnalyzerVictims(t *testing.T) {
	// test victims
	assert.Zero(t, analyzer_wrapper.Run(
		[]string{strings.Join([]string{selfPackagePath(), "..", "internal", "testdata", "padding", "pass"}, "/")},
		[]*analysis.Analyzer{Analyzer},
	))
	assert.NotZero(t, analyzer_wrapper.Run(
		[]string{strings.Join([]string{selfPackagePath(), "..", "internal", "testdata", "padding", "fail"}, "/")},
		[]*analysis.Analyzer{Analyzer},
	))
}
