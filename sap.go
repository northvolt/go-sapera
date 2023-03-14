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
	CORACQ_PRM_META_DATA_CLEAR     = C.CORACQ_PRM_META_DATA_CLEAR
	CORACQ_PRM_OUTPUT_FORMAT       = C.CORACQ_PRM_OUTPUT_FORMAT
	CORACQ_PRM_PIXEL_DEPTH         = C.CORACQ_PRM_PIXEL_DEPTH
	CORACQ_PRM_CROP_HEIGHT         = C.CORACQ_PRM_CROP_HEIGHT
	CORACQ_PRM_CROP_WIDTH          = C.CORACQ_PRM_CROP_WIDTH
	CORACQ_PRM_LINE_TRIGGER_ENABLE = C.CORACQ_PRM_LINE_TRIGGER_ENABLE

	CORACQ_VAL_META_DATA_PER_LINE_RIGHT = C.CORACQ_VAL_META_DATA_PER_LINE_RIGHT

	CORBUFFER_PRM_ADDRESS = C.CORBUFFER_PRM_ADDRESS

	// registers related to horizontal flip
	CORACQ_PRM_FLIP = C.CORACQ_PRM_FLIP
	CORACQ_CAP_FLIP = C.CORACQ_CAP_FLIP

	// registers for line/free mode
	CORACQ_PRM_SHAFT_ENCODER_ENABLE    = C.CORACQ_PRM_SHAFT_ENCODER_ENABLE
	CORACQ_PRM_SHAFT_ENCODER_COUNT     = C.CORACQ_PRM_SHAFT_ENCODER_COUNT
	CORACQ_PRM_INT_LINE_TRIGGER_ENABLE = C.CORACQ_PRM_INT_LINE_TRIGGER_ENABLE

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

	// Monochrome data formats
	SapFormatMono8  = C.CORDATA_FORMAT_MONO8
	SapFormatInt8   = C.CORDATA_FORMAT_INT8
	SapFormatUint8  = C.CORDATA_FORMAT_UINT8
	SapFormatMono16 = C.CORDATA_FORMAT_MONO16
	SapFormatInt16  = C.CORDATA_FORMAT_INT16
	SapFormatUint16 = C.CORDATA_FORMAT_UINT16
	SapFormatMono24 = C.CORDATA_FORMAT_MONO24
	SapFormatInt24  = C.CORDATA_FORMAT_INT24
	SapFormatUint24 = C.CORDATA_FORMAT_UINT24
	SapFormatMono32 = C.CORDATA_FORMAT_MONO32
	SapFormatInt32  = C.CORDATA_FORMAT_INT32
	SapFormatUint32 = C.CORDATA_FORMAT_UINT32
	SapFormatMono64 = C.CORDATA_FORMAT_MONO64
	SapFormatInt64  = C.CORDATA_FORMAT_INT64
	SapFormatUint64 = C.CORDATA_FORMAT_UINT64

	// Color RGB data formats
	SapFormatRGB5551     = C.CORDATA_FORMAT_RGB5551     // 16-bit
	SapFormatRGB565      = C.CORDATA_FORMAT_RGB565      // 16-bit
	SapFormatRGB888      = C.CORDATA_FORMAT_RGB888      // 24-bit
	SapFormatRGBR888     = C.CORDATA_FORMAT_RGBR888     // 24-bit
	SapFormatRGB8888     = C.CORDATA_FORMAT_RGB8888     // 32-bit
	SapFormatRGB101010   = C.CORDATA_FORMAT_RGB101010   // 32-bit
	SapFormatRGB161616   = C.CORDATA_FORMAT_RGB161616   // 48-bit
	SapFormatRGB16161616 = C.CORDATA_FORMAT_RGB16161616 // 64-bit

	// Flip mode
	SapFlipOff        = C.CORACQ_VAL_FLIP_OFF
	SapFlipHorizontal = C.CORACQ_VAL_FLIP_HORZ
	SapFlipVertical   = C.CORACQ_VAL_FLIP_VERT
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

type SapVersionInfo struct {
	p C.SapVersionInfoWrapper
}

func GetVersionInfo() SapVersionInfo {
	return SapVersionInfo{p: C.SapManager_GetVersionInfo()}
}

func (v *SapVersionInfo) Delete() {
	C.SapVersionInfo_Delete(v.p)
}

func (v *SapVersionInfo) GetMajor() int {
	return int(C.SapVersionInfo_GetMajor(v.p))
}

func (v *SapVersionInfo) GetMinor() int {
	return int(C.SapVersionInfo_GetMinor(v.p))
}

func (v *SapVersionInfo) GetRevision() int {
	return int(C.SapVersionInfo_GetRevision(v.p))
}

func (v *SapVersionInfo) GetBuild() int {
	return int(C.SapVersionInfo_GetBuild(v.p))
}
