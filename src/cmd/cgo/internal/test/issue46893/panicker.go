// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package issue46893

import "runtime/debug"

type Panicker interface {
	Panic([]byte) ([]byte, error)
}

type PanickerImpl struct {
}

func (p *PanickerImpl) Panic(b []byte) ([]byte, error) {
	return make([]byte, len(b)), nil
}

func Init() {
	debug.SetTraceback("crash")
}
