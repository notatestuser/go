package issue46893

// Go support functions for generated Go bindings. This file is
// copied into the generated main package, and compiled along
// with the bindings.

// #cgo android CFLAGS: -D__GOBIND_ANDROID__
// #cgo darwin CFLAGS: -D__GOBIND_DARWIN__
// #include <stdlib.h>
// #include "seq.h"
import "C"

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	FinalizeRef = func(ref *Ref) {
		refnum := ref.Bind_Num
		if refnum < 0 {
			panic(fmt.Sprintf("not a foreign ref: %d", refnum))
		}
		C.go_seq_dec_ref(C.int32_t(refnum))
	}
	IncForeignRef = func(refnum int32) {
		if refnum < 0 {
			panic(fmt.Sprintf("not a foreign ref: %d", refnum))
		}
		C.go_seq_inc_ref(C.int32_t(refnum))
	}
	// Workaround for issue #17393.
	signal.Notify(make(chan os.Signal), syscall.SIGPIPE)
}

// IncGoRef is called by foreign code to pin a Go object while its refnum is crossing
// the language barrier
//
//export IncGoRef
func IncGoRef(refnum C.int32_t) {
	Inc(int32(refnum))
}
