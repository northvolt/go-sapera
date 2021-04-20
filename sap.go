package sapera

/*
#include <stdlib.h>
#include "sap.h"
#include "mx4.h"
*/
import "C"
import "unsafe"

const (
	CORACQ_PRM_META_DATA           = C.CORACQ_PRM_META_DATA
	CORACQ_PRM_OUTPUT_FORMAT       = C.CORACQ_PRM_OUTPUT_FORMAT
	CORACQ_PRM_PIXEL_DEPTH         = C.CORACQ_PRM_PIXEL_DEPTH
	CORACQ_PRM_CROP_HEIGHT         = C.CORACQ_PRM_CROP_HEIGHT
	CORACQ_PRM_CROP_WIDTH          = C.CORACQ_PRM_CROP_WIDTH
	CORACQ_PRM_LINE_TRIGGER_ENABLE = C.CORACQ_PRM_LINE_TRIGGER_ENABLE

	CORACQ_VAL_META_DATA_PER_LINE_RIGHT = C.CORACQ_VAL_META_DATA_PER_LINE_RIGHT

	CORBUFFER_PRM_ADDRESS = C.CORBUFFER_PRM_ADDRESS
)

// Init initializes the SaperaLT SDK signaling. This function must be called from the main thread of the application,
// and it also must be called before any other functions of the SaperaLT SDK are used.
func Init() {
	C.SapManager_Init()
}

type SapLocation struct {
	p C.SapLocationWrapper
}

func NewSapLocation(acqServerName string, acqDeviceNumber uint32) SapLocation {
	cName := C.CString(acqServerName)
	defer C.free(unsafe.Pointer(cName))

	return SapLocation{p: C.SapLocation_New(cName, C.int(acqDeviceNumber))}
}

func (loc *SapLocation) Delete() {
	C.SapLocation_Delete(loc.p)
}

func GetServerCount() int {
	return int(C.SapManager_GetServerCount())
}

func GetLastStatus() string {
	return C.GoString(C.SapManager_GetLastStatus())
}
