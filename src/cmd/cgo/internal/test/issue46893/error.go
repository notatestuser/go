// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package issue46893

/*
#include <stdlib.h>
#include <stdint.h>
#include "seq.h"
#include "universe.h"
*/
import "C"

import (
	"cmd/cgo/internal/test/issue46893/seq"
)

// suppress the error if seq ends up unused
var _ = seq.FromRefNum

//export proxy_error_Error
func proxy_error_Error(refnum C.int32_t) C.nstring {
	ref := seq.FromRefNum(int32(refnum))
	v := ref.Get().(error)
	res_0 := v.Error()
	_res_0 := encodeString(res_0)
	return _res_0
}

type proxy_error seq.Ref

func (p *proxy_error) Bind_proxy_refnum__() int32 {
	return (*seq.Ref)(p).Bind_IncNum()
}

func (p *proxy_error) Error() string {
	res := C.cproxy_error_Error(C.int32_t(p.Bind_proxy_refnum__()))
	_res := decodeString(res)
	return _res
}
