// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assign_test

import (
	"testing"

	"go.coder.com/go-tools/go/analysis/analysistest"
	"go.coder.com/go-tools/go/analysis/passes/assign"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, assign.Analyzer, "a")
}
