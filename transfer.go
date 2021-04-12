package sapera

/*
#include <stdlib.h>
#include "transfer.h"
#include "context.h"
*/
import "C"

import "unsafe"

// SapAcqToBuf is used to transfer data from the acquisition device to a provided buffer.
type SapAcqToBuf struct {
	// C.SapAcqToBufWrapper
	p unsafe.Pointer
}

func NewSapAcqToBuf(acq SapAcquisition, buf SapBuffer, context SapXferContext) SapAcqToBuf {
	return SapAcqToBuf{p: unsafe.Pointer(C.SapAcqToBuf_New((C.SapAcquisitionWrapper)(acq.p), (C.SapBufferWrapper)(buf.p), (C.SapXferContextWrapper)(context.p)))}
}

func (trs *SapAcqToBuf) Create() bool {
	return bool(C.SapAcqToBuf_Create((C.SapAcqToBufWrapper)(trs.p)))
}

func (trs *SapAcqToBuf) Delete() {
	C.SapAcqToBuf_Delete((C.SapAcqToBufWrapper)(trs.p))
}

func (trs *SapAcqToBuf) Destroy() bool {
	return bool(C.SapAcqToBuf_Destroy((C.SapAcqToBufWrapper)(trs.p)))
}

func (trs *SapAcqToBuf) Grab() bool {
	return bool(C.SapAcqToBuf_Grab((C.SapAcqToBufWrapper)(trs.p)))
}

func (trs *SapAcqToBuf) Snap() bool {
	return bool(C.SapAcqToBuf_Snap((C.SapAcqToBufWrapper)(trs.p)))
}

func (trs *SapAcqToBuf) IsGrabbing() bool {
	return bool(C.SapAcqToBuf_IsGrabbing((C.SapAcqToBufWrapper)(trs.p)))
}

func (trs *SapAcqToBuf) Freeze() bool {
	return bool(C.SapAcqToBuf_Freeze((C.SapAcqToBufWrapper)(trs.p)))
}

func (trs *SapAcqToBuf) Wait(ms int) bool {
	return bool(C.SapAcqToBuf_Wait((C.SapAcqToBufWrapper)(trs.p), C.int(ms)))
}

func (trs *SapAcqToBuf) Abort() bool {
	return bool(C.SapAcqToBuf_Abort((C.SapAcqToBufWrapper)(trs.p)))
}

func (trs *SapAcqToBuf) GetLastStatus() string {
	return C.GoString(C.SapAcqToBuf_GetLastStatus((C.SapAcqToBufWrapper)(trs.p)))
}

// SapXferCallbackInfo is to pass context info to transfer callbacks.
type SapXferCallbackInfo struct {
	p C.SapXferCallbackInfoWrapper
}

func (trs *SapAcqToBuf) SetCallbackHandler(hdl func(cb SapXferCallbackInfo), context SapXferContext) bool {
	if len(xferCallbackHandler) == 0 {
		xferCallbackHandler = append(xferCallbackHandler, hdl)
	}

	return bool(C.SapAcqToBuf_SetCallbackInfo((C.SapAcqToBufWrapper)(trs.p), (C.SapXferContextWrapper)(context.p)))
}

func (trs *SapAcqToBuf) UnregisterCallback() bool {
	return bool(C.SapAcqToBuf_UnregisterCallback((C.SapAcqToBufWrapper)(trs.p)))
}

// SapXferFrameRateInfo is to retrieve frame rate stats for transfers. Does not seem to work yet.
type SapXferFrameRateInfo struct {
	p C.SapXferFrameRateInfoWrapper
}

func (trs *SapAcqToBuf) GetFrameRateStatistics() SapXferFrameRateInfo {
	return SapXferFrameRateInfo{p: C.SapAcqToBuf_GetFrameRateStatistics((C.SapAcqToBufWrapper)(trs.p))}
}

func (inf *SapXferFrameRateInfo) GetBufferFrameRate() int {
	return int(C.SapXferFrameRateInfo_GetBufferFrameRate(inf.p))
}

func (cb SapXferCallbackInfo) GetContext() SapXferContext {
	return SapXferContext{p: unsafe.Pointer(C.SapXferCallbackInfo_GetContext(cb.p))}
}
