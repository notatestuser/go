// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo

package cgotest

import (
	"testing"

	"cmd/cgo/internal/test/issue46893"
)

func Test46893(t *testing.T) {
	issue46893.Test(t)
}
