package sapera

/*
#cgo CXXFLAGS:  -g -DPOSIX_HOSTPC -D_REENTRANT -Dx86_64 -Wall -Wno-parentheses -Wno-missing-braces -Wno-unused-but-set-variable -Wno-unknown-pragmas -Wno-cast-qual -Wno-unused-function -Wno-unused-label
#cgo CPPFLAGS:  -g -I/usr/local/include/Sapera/include -I/usr/local/include/Sapera/classes/basic
#cgo LDFLAGS:   -lpthread -lstdc++ -L/usr/local/lib -lSapera++ -lSaperaLT -lCorW32
#include <stdlib.h>
#include "sap.h"

extern void goxferhandler(SapXferCallbackInfoWrapper pinfo);
extern void goacqhandler(SapAcqCallbackInfoWrapper pinfo);
*/
import "C"

import "fmt"

var xferCallbackHandler []func(cb SapXferCallbackInfo)

//export goxferhandler
func goxferhandler(cb C.SapXferCallbackInfoWrapper) {
	cbinfo := SapXferCallbackInfo{p: cb}
	// TODO: use cbinfo context to determine which callback applies to this transfer
	if len(xferCallbackHandler) > 0 && xferCallbackHandler[0] != nil {
		xferCallbackHandler[0](cbinfo)
		return
	}

	ctx := cbinfo.GetContext()
	fmt.Println("no callback handler for frame transfer from ID", ctx.GetID())
}

var acqCallbackHandler []func(cb SapAcqCallbackInfo)

//export goacqhandler
func goacqhandler(cb C.SapAcqCallbackInfoWrapper) {
	cbinfo := SapAcqCallbackInfo{p: cb}
	// TODO: use cbinfo context to determine which callback applies to this acquisition
	if len(acqCallbackHandler) > 0 && acqCallbackHandler[0] != nil {
		acqCallbackHandler[0](cbinfo)
		return
	}

	fmt.Println("no callback handler for acquisition")
}
