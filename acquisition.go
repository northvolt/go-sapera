package sapera

/*
#include <stdlib.h>
#include "acquisition.h"
*/
import "C"
import "unsafe"

type SapAcquisition struct {
	// C.SapAcquisitionWrapper
	p unsafe.Pointer
}

// SapAcqCallbackInfo is to pass context info to acquisition callbacks.
type SapAcqCallbackInfo struct {
	p C.SapAcqCallbackInfoWrapper
}

func NewSapAcquisition() SapAcquisition {
	return SapAcquisition{p: unsafe.Pointer(C.SapAcquisition_New())}
}

func NewSapAcquisitionForLocation(loc SapLocation, camFilename string) SapAcquisition {
	cName := C.CString(camFilename)
	defer C.free(unsafe.Pointer(cName))

	return SapAcquisition{p: unsafe.Pointer(C.SapAcquisition_NewForLocation(loc.p, cName))}
}

func (acq *SapAcquisition) Delete() {
	C.SapAcquisition_Delete((C.SapAcquisitionWrapper)(acq.p))
}

func (acq *SapAcquisition) Create() bool {
	return bool(C.SapAcquisition_Create((C.SapAcquisitionWrapper)(acq.p)))
}

func (acq *SapAcquisition) Destroy() bool {
	return bool(C.SapAcquisition_Destroy((C.SapAcquisitionWrapper)(acq.p)))
}

func (acq *SapAcquisition) SetParameter(param int, val int) bool {
	return bool(C.SapAcquisition_SetParameter((C.SapAcquisitionWrapper)(acq.p), C.int(param), C.int(val)))
}

func (acq *SapAcquisition) GetParameterInt32(param int) (int32, bool) {
	val := C.int(0)
	result := C.SapAcquisition_GetParameterInt32((C.SapAcquisitionWrapper)(acq.p), C.int(param), &val)
	return int32(val), bool(result)
}

func (acq *SapAcquisition) GetParameterInt64(param int) (int64, bool) {
	val := C.long(0)
	result := C.SapAcquisition_GetParameterInt64((C.SapAcquisitionWrapper)(acq.p), C.int(param), &val)
	return int64(val), bool(result)
}

func (acq *SapAcquisition) GetLastStatus() string {
	return C.GoString(C.SapAcquisition_GetLastStatus((C.SapAcquisitionWrapper)(acq.p)))
}

func (acq *SapAcquisition) ResetTimeStamp() bool {
	return bool(C.SapAcquisition_ResetTimeStamp((C.SapAcquisitionWrapper)(acq.p)))
}

func (acq *SapAcquisition) RegisterCallback(event uint64, hdl func(cb SapAcqCallbackInfo), context SapAcqContext) bool {
	if len(acqCallbackHandler) == 0 {
		acqCallbackHandler = append(acqCallbackHandler, hdl)
	}

	return bool(C.SapAcquisition_RegisterCallback((C.SapAcquisitionWrapper)(acq.p), C.UINT64(event), (C.SapAcqContextWrapper)(context.p)))
}

func (acq *SapAcquisition) UnregisterCallbacks() bool {
	return bool(C.SapAcquisition_UnregisterCallbacks((C.SapAcquisitionWrapper)(acq.p)))
}

func (acq *SapAcquisition) SetCallbackInfo(hdl func(cb SapAcqCallbackInfo), context SapAcqContext) bool {
	if len(acqCallbackHandler) == 0 {
		acqCallbackHandler = append(acqCallbackHandler, hdl)
	}

	return bool(C.SapAcquisition_SetCallbackInfo((C.SapAcquisitionWrapper)(acq.p)))
}

func (acq *SapAcquisition) SetEventType(event uint64) bool {
	return bool(C.SapAcquisition_SetEventType((C.SapAcquisitionWrapper)(acq.p), C.UINT64(event)))
}

func (acqcb *SapAcqCallbackInfo) GetEventType(event uint64) bool {
	return bool(C.SapAcqCallbackInfoWrapper_GetEventType((C.SapAcqCallbackInfoWrapper)(acqcb.p), C.UINT64(event)))
}

// func (cb SapAcqCallbackInfo) GetContext() SapAcqContext {
// 	return SapAcqContext{p: unsafe.Pointer(C.SapAcqCallbackInfo_GetContext(cb.p))}
// }

func (acq *SapAcquisition) SetFlipMode(flipMode int) bool {
	return bool(C.SapAcquisition_SetFlipMode((C.SapAcquisitionWrapper)(acq.p), C.int(flipMode)))
}
