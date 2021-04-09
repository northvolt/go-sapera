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
