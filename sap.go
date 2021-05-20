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

	// acquisition events
	AcqEventNone                             = 0
	AcqEventShaftEncoderReverseCountOverflow = C.CORACQ_VAL_EVENT_TYPE_SHAFT_ENCODER_REVERSE_COUNT_OVERFLOW
	AcqEventLineTriggerTooFast               = C.CORACQ_VAL_EVENT_TYPE_LINE_TRIGGER_TOO_FAST
	AcqEventPoCL                             = C.CORACQ_VAL_EVENT_TYPE_POCL
	AcqEventLinkError                        = C.CORACQ_VAL_EVENT_TYPE_LINK_ERROR
	AcqEventVerticalTimeout                  = C.CORACQ_VAL_EVENT_TYPE_VERTICAL_TIMEOUT
	AcqEventExternalTrigger2                 = C.CORACQ_VAL_EVENT_TYPE_EXTERNAL_TRIGGER2
	AcqEventVirtualFrame                     = C.CORACQ_VAL_EVENT_TYPE_VIRTUAL_FRAME
	AcqEventUserDefine                       = C.CORACQ_VAL_EVENT_TYPE_USER_DEFINE
	AcqEventExternalTriggerTooSlow           = C.CORACQ_VAL_EVENT_TYPE_EXT_LINE_TRIGGER_TOO_SLOW
	AcqEventExtLineTriggerTooSlow            = C.CORACQ_VAL_EVENT_TYPE_EXT_LINE_TRIGGER_TOO_SLOW
	AcqEventHsyncLock                        = C.CORACQ_VAL_EVENT_TYPE_HSYNC_LOCK
	AcqEventHsyncUnlock                      = C.CORACQ_VAL_EVENT_TYPE_HSYNC_UNLOCK
	AcqEventExternalTriggerIgnored           = C.CORACQ_VAL_EVENT_TYPE_EXTERNAL_TRIGGER_IGNORED
	AcqEventDataOverflow                     = C.CORACQ_VAL_EVENT_TYPE_DATA_OVERFLOW
	AcqEventFrameLost                        = C.CORACQ_VAL_EVENT_TYPE_FRAME_LOST
	AcqEventStartOfField                     = C.CORACQ_VAL_EVENT_TYPE_START_OF_FIELD
	AcqEventStartOfOdd                       = C.CORACQ_VAL_EVENT_TYPE_START_OF_ODD
	AcqEventStartOfEven                      = C.CORACQ_VAL_EVENT_TYPE_START_OF_EVEN
	AcqEventStartOfFrame                     = C.CORACQ_VAL_EVENT_TYPE_START_OF_FRAME
	AcqEventEndOfField                       = C.CORACQ_VAL_EVENT_TYPE_END_OF_FIELD
	AcqEventEndOfOdd                         = C.CORACQ_VAL_EVENT_TYPE_END_OF_ODD
	AcqEventEndOfEven                        = C.CORACQ_VAL_EVENT_TYPE_END_OF_EVEN
	AcqEventEndOfFrame                       = C.CORACQ_VAL_EVENT_TYPE_END_OF_FRAME
	AcqEventExternalTrigger                  = C.CORACQ_VAL_EVENT_TYPE_EXTERNAL_TRIGGER
	AcqEventExternalTriggerEnd               = C.CORACQ_VAL_EVENT_TYPE_EXTERNAL_TRIGGER_END
	AcqEventVerticalSync                     = C.CORACQ_VAL_EVENT_TYPE_VERTICAL_SYNC
	AcqEventEndOfLine                        = C.CORACQ_VAL_EVENT_TYPE_END_OF_LINE
	AcqEventEndOfNLines                      = C.CORACQ_VAL_EVENT_TYPE_END_OF_NLINES
	AcqEventNoHSync                          = C.CORACQ_VAL_EVENT_TYPE_NO_HSYNC
	AcqEventNoVSync                          = C.CORACQ_VAL_EVENT_TYPE_NO_VSYNC
	AcqEventNoPixelClk                       = C.CORACQ_VAL_EVENT_TYPE_NO_PIXEL_CLK
	AcqEventPixelClk                         = C.CORACQ_VAL_EVENT_TYPE_PIXEL_CLK
	AcqEventMask                             = C.CORACQ_VAL_EVENT_TYPE_MASK
	AcqEventLinkLock                         = C.CORACQ_VAL_EVENT_TYPE_LINK_LOCK
	AcqEventLinkUnlock                       = C.CORACQ_VAL_EVENT_TYPE_LINK_UNLOCK
	AcqEventCameraMissedTrigger              = C.CORACQ_VAL_EVENT_TYPE_CAMERA_MISSED_TRIGGER
	AcqEventCameraBufferOverrun              = C.CORACQ_VAL_EVENT_TYPE_CAMERA_BUFFER_OVERRUN
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
