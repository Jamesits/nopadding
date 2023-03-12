# nopadding

Provides:
- A Golang analyzer to show whether your struct has been automatically padded by the Golang compiler
- A Golang analyzer wrapper which can be invoked from a unit test

This package is very useful during system development (e.g. when marshalling syscall returned structs).

## Usage

Put these code segments in your package's unit test.

### Check for Non-optimized Struct Layout

```go
package main

import (
	"github.com/jamesits/nopadding/analyzer_wrapper"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/fieldalignment"
	"testing"
)

func TestStructOptimizedLayout(t *testing.T) {
	assert.NotZero(t, analyzer_wrapper.Run(
		[]string{},
		[]*analysis.Analyzer{fieldalignment.Analyzer},
	))
}

```

### Check for Padding Existence

```go
package main

import (
	"github.com/jamesits/nopadding/analyzer_wrapper"
	"github.com/jamesits/nopadding/padding"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis"
	"testing"
)

func TestStructPadding(t *testing.T) {
	assert.Zero(t, analyzer_wrapper.Run(
		[]string{},
		[]*analysis.Analyzer{padding.Analyzer},
	))
}
```
