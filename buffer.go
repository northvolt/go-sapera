package sapera

/*
#include <stdlib.h>
#include "buffer.h"
*/
import "C"
import "unsafe"

type SapBuffer struct {
	// C.SapBufferWrapper
	p unsafe.Pointer
}

func NewSapBuffer() SapBuffer {
	return SapBuffer{p: unsafe.Pointer(C.SapBuffer_New())}
}

func NewSapBufferForAcquisition(count int, src SapAcquisition) SapBuffer {
	return SapBuffer{p: unsafe.Pointer(C.SapBuffer_NewForAcquisition(C.int(count), (C.SapAcquisitionWrapper)(src.p)))}
}

func NewSapBufferWithSize(count, width, height, format int) SapBuffer {
	return SapBuffer{p: unsafe.Pointer(C.SapBuffer_NewWithSize(C.int(count), C.int(width), C.int(height), C.int(format)))}
}

func (buf *SapBuffer) Create() bool {
	return bool(C.SapBuffer_Create((C.SapBufferWrapper)(buf.p)))
}

func (buf *SapBuffer) Delete() {
	C.SapBuffer_Delete((C.SapBufferWrapper)(buf.p))
}

func (buf *SapBuffer) Destroy() bool {
	return bool(C.SapBuffer_Destroy((C.SapBufferWrapper)(buf.p)))
}

func (buf *SapBuffer) ReadLine(x1, y1, x2, y2 int, data []byte) (int, bool) {
	cCount := C.int(0)

	// line must be in a single row
	if y1 != y2 {
		return 0, false
	}

	// data is not large enough for desired bytes to be read
	if len(data) < x2-x1 {
		return 0, false
	}

	result := C.SapBuffer_ReadLine((C.SapBufferWrapper)(buf.p), C.int(x1), C.int(y1), C.int(x2), C.int(y2), (unsafe.Pointer(&data[0])), &cCount)
	return int(cCount), bool(result)
}

func (buf *SapBuffer) Copy(src SapBuffer, srcIndex, dstIndex int) bool {
	return bool(C.SapBuffer_Copy((C.SapBufferWrapper)(buf.p), (C.SapBufferWrapper)(src.p), C.int(srcIndex), C.int(dstIndex)))
}

func (buf *SapBuffer) CopyRect(src SapBuffer, srcIndex, xSrc, ySrc, width, height, dstIndex, xDest, yDest int) bool {
	return bool(C.SapBuffer_CopyRect((C.SapBufferWrapper)(buf.p), (C.SapBufferWrapper)(src.p), C.int(srcIndex), C.int(xSrc), C.int(ySrc), C.int(width), C.int(height), C.int(dstIndex), C.int(xDest), C.int(yDest)))
}

func (buf *SapBuffer) CopyAll(src SapBuffer) bool {
	return bool(C.SapBuffer_CopyAll((C.SapBufferWrapper)(buf.p), (C.SapBufferWrapper)(src.p)))
}

func (buf *SapBuffer) Save(fileName, options string) bool {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))

	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))

	return bool(C.SapBuffer_Save((C.SapBufferWrapper)(buf.p), cfileName, coptions))
}

func (buf *SapBuffer) SaveForIndex(fileName, options string, index int) bool {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))

	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))

	return bool(C.SapBuffer_SaveForIndex((C.SapBufferWrapper)(buf.p), cfileName, coptions, C.int(index)))
}

func (buf *SapBuffer) SetParameter(param int, val int) bool {
	return bool(C.SapBuffer_SetParameter((C.SapBufferWrapper)(buf.p), C.int(param), C.int(val)))
}

func (buf *SapBuffer) GetParameterInt32(param int) (int32, bool) {
	val := C.int(0)
	result := C.SapBuffer_GetParameterInt32((C.SapBufferWrapper)(buf.p), C.int(param), &val)
	return int32(val), bool(result)
}

func (buf *SapBuffer) GetParameterInt64(param int) (int64, bool) {
	val := C.long(0)
	result := C.SapBuffer_GetParameterInt64((C.SapBufferWrapper)(buf.p), C.int(param), &val)
	return int64(val), bool(result)
}

func (buf *SapBuffer) SetSrcNode(acq SapAcquisition) bool {
	C.SapBuffer_SetSrcNode((C.SapBufferWrapper)(buf.p), (C.SapAcquisitionWrapper)(acq.p))
	return true
}

func (buf *SapBuffer) ClearSrcNode() bool {
	C.SapBuffer_ClearSrcNode((C.SapBufferWrapper)(buf.p))
	return true
}

func (buf *SapBuffer) SetWidth(width int) bool {
	C.SapBuffer_SetWidth((C.SapBufferWrapper)(buf.p), C.int(width))
	return true
}

func (buf *SapBuffer) GetWidth() int {
	return int(C.SapBuffer_GetWidth((C.SapBufferWrapper)(buf.p)))
}

func (buf *SapBuffer) SetHeight(height int) bool {
	C.SapBuffer_SetHeight((C.SapBufferWrapper)(buf.p), C.int(height))
	return true
}

func (buf *SapBuffer) GetHeight() int {
	return int(C.SapBuffer_GetHeight((C.SapBufferWrapper)(buf.p)))
}

func (buf *SapBuffer) SetFormat(format int) bool {
	C.SapBuffer_SetFormat((C.SapBufferWrapper)(buf.p), C.int(format))
	return true
}

func (buf *SapBuffer) GetLastStatus() string {
	return C.GoString(C.SapBuffer_GetLastStatus((C.SapBufferWrapper)(buf.p)))
}

func (buf *SapBuffer) GetIndex() int {
	return int(C.SapBuffer_GetIndex((C.SapBufferWrapper)(buf.p)))
}

func (buf *SapBuffer) SetIndex(index int) bool {
	C.SapBuffer_SetIndex((C.SapBufferWrapper)(buf.p), C.int(index))
	return true
}

func (buf *SapBuffer) Next() {
	C.SapBuffer_Next((C.SapBufferWrapper)(buf.p))
}
