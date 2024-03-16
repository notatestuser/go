// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo

package issue46893

import (
	"testing"
)

/*
#cgo CFLAGS: -x objective-c -fmodules
#cgo LDFLAGS: -framework Foundation

void issue46893(void);
*/
import "C"

func Test(t *testing.T) {
	C.issue46893()
}
