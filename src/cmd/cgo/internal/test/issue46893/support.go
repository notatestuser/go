// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package issue46893

import "C"
import (
	"unsafe"
)

// encodeString copies a Go string and returns it as a nstring.
func encodeString(s string) C.nstring {
	n := C.int(len(s))
	if n == 0 {
		return C.nstring{}
	}
	ptr := C.malloc(C.size_t(n))
	if ptr == nil {
		panic("encodeString: malloc failed")
	}
	copy((*[1<<31 - 1]byte)(ptr)[:n], s)
	return C.nstring{ptr: ptr, len: n}
}

// decodeString converts a nstring to a Go string. The
// data in str is freed after use.
func decodeString(str C.nstring) string {
	if str.ptr == nil {
		return ""
	}
	s := C.GoStringN((*C.char)(str.ptr), str.len)
	C.free(str.ptr)
	return s
}

// fromSlice converts a slice to a nbyteslice.
// If cpy is set, a malloc'ed copy of the data is returned.
func fromSlice(s []byte, cpy bool) C.nbyteslice {
	if s == nil || len(s) == 0 {
		return C.nbyteslice{}
	}
	ptr, n := unsafe.Pointer(&s[0]), C.int(len(s))
	if cpy {
		nptr := C.malloc(C.size_t(n))
		if nptr == nil {
			panic("fromSlice: malloc failed")
		}
		copy((*[1<<31 - 1]byte)(nptr)[:n], (*[1<<31 - 1]byte)(ptr)[:n])
		ptr = nptr
	}
	return C.nbyteslice{ptr: ptr, len: n}
}

// toSlice takes a nbyteslice and returns a byte slice with the data. If cpy is
// set, the slice contains a copy of the data. If not, the generated Go code
// calls releaseByteSlice after use.
func toSlice(s C.nbyteslice, cpy bool) []byte {
	if s.ptr == nil || s.len == 0 {
		return nil
	}
	var b []byte
	if cpy {
		b = C.GoBytes(s.ptr, C.int(s.len))
		C.free(s.ptr)
	} else {
		b = (*[1<<31 - 1]byte)(unsafe.Pointer(s.ptr))[:s.len:s.len]
	}
	return b
}
