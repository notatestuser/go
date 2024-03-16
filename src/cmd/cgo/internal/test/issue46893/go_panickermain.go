// Code generated by gobind. DO NOT EDIT.

// Package main is an autogenerated binder stub for package panicker.
//
//	autogenerated by gobind -lang=go command-line-arguments
package issue46893

/*
#include <stdlib.h>
#include <stdint.h>
#include "seq.h"
#include "panicker.h"

*/
import "C"

// suppress the error if seq ends up unused
var _ = FromRefNum

//export proxypanicker_PanickerImpl_Panic
func proxypanicker_PanickerImpl_Panic(refnum C.int32_t, param_b C.nbyteslice) (C.nbyteslice, C.int32_t) {
	ref := FromRefNum(int32(refnum))
	v := ref.Get().(*PanickerImpl)
	_param_b := toSlice(param_b, false)
	res_0, res_1 := v.Panic(_param_b)
	_res_0 := fromSlice(res_0, true)
	var _res_1 C.int32_t = NullRefNum
	if res_1 != nil {
		_res_1 = C.int32_t(ToRefNum(res_1))
	}
	return _res_0, _res_1
}

//export new_panicker_PanickerImpl
func new_panicker_PanickerImpl() C.int32_t {
	return C.int32_t(ToRefNum(new(PanickerImpl)))
}

//export proxypanicker_Panicker_Panic
func proxypanicker_Panicker_Panic(refnum C.int32_t, param_p0 C.nbyteslice) (C.nbyteslice, C.int32_t) {
	ref := FromRefNum(int32(refnum))
	v := ref.Get().(Panicker)
	_param_p0 := toSlice(param_p0, false)
	res_0, res_1 := v.Panic(_param_p0)
	_res_0 := fromSlice(res_0, true)
	var _res_1 C.int32_t = NullRefNum
	if res_1 != nil {
		_res_1 = C.int32_t(ToRefNum(res_1))
	}
	return _res_0, _res_1
}

type proxypanicker_Panicker Ref

func (p *proxypanicker_Panicker) Bind_proxy_refnum__() int32 {
	return (*Ref)(p).Bind_IncNum()
}

func (p *proxypanicker_Panicker) Panic(param_p0 []byte) ([]byte, error) {
	_param_p0 := fromSlice(param_p0, false)
	res := C.cproxypanicker_Panicker_Panic(C.int32_t(p.Bind_proxy_refnum__()), _param_p0)
	res_0 := toSlice(res.r0, true)
	var res_1 error
	res_1_ref := FromRefNum(int32(res.r1))
	if res_1_ref != nil {
		if res.r1 < 0 { // go object
			res_1 = res_1_ref.Get().(error)
		} else { // foreign object
			res_1 = (*proxy_error)(res_1_ref)
		}
	}
	return res_0, res_1
}

//export proxypanicker__Init
func proxypanicker__Init() {
	Init()
}
