// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "_cgo_export.h"

void issue46893() {
    int32_t panicker = new_panicker_PanickerImpl();
    unsigned char* d = panicker_getCNByteSlice();
    for (int i = 0; i < 10000000; i++) {
        proxypanicker_PanickerImpl_Panic(panicker, d);
    }
}
