package sapera

/*
#include <stdlib.h>
#include "context.h"
*/
import "C"

import "unsafe"

type SapXferContext struct {
	// C.SapXferContextWrapper
	p unsafe.Pointer
}

func NewSapXferContext() SapXferContext {
	return SapXferContext{p: unsafe.Pointer(C.SapXferContext_New())}
}

func (ctx *SapXferContext) Delete() {
	C.SapXferContext_Delete((C.SapXferContextWrapper)(ctx.p))
}

func (ctx *SapXferContext) SetID(id int) {
	C.SapXferContext_SetID((C.SapXferContextWrapper)(ctx.p), C.int(id))
}

func (ctx *SapXferContext) GetID() int {
	return int(C.SapXferContext_GetID((C.SapXferContextWrapper)(ctx.p)))
}

func (ctx *SapXferContext) SetBuffer(buf SapBuffer) {
	C.SapXferContext_SetBuffer((C.SapXferContextWrapper)(ctx.p), (C.SapBufferWrapper)(buf.p))
}

func (ctx *SapXferContext) GetBuffer() SapBuffer {
	return SapBuffer{p: unsafe.Pointer(C.SapXferContext_GetBuffer((C.SapXferContextWrapper)(ctx.p)))}
}

func (ctx *SapXferContext) SetHeight(height int) {
	C.SapXferContext_SetHeight((C.SapXferContextWrapper)(ctx.p), C.int(height))
}

func (ctx *SapXferContext) GetHeight() int {
	return int(C.SapXferContext_GetHeight((C.SapXferContextWrapper)(ctx.p)))
}

func (ctx *SapXferContext) SetWidth(width int) {
	C.SapXferContext_SetWidth((C.SapXferContextWrapper)(ctx.p), C.int(width))
}

func (ctx *SapXferContext) GetWidth() int {
	return int(C.SapXferContext_GetWidth((C.SapXferContextWrapper)(ctx.p)))
}

func (ctx *SapXferContext) SetPixelDepth(depth int) {
	C.SapXferContext_SetPixelDepth((C.SapXferContextWrapper)(ctx.p), C.int(depth))
}

func (ctx *SapXferContext) GetPixelDepth() int {
	return int(C.SapXferContext_GetPixelDepth((C.SapXferContextWrapper)(ctx.p)))
}

func (ctx *SapXferContext) SetCounter(count int64) {
	C.SapXferContext_SetCounter((C.SapXferContextWrapper)(ctx.p), C.long(count))
}

func (ctx *SapXferContext) GetCounter() int64 {
	return int64(C.SapXferContext_GetCounter((C.SapXferContextWrapper)(ctx.p)))
}

type SapAcqContext struct {
	// C.SapAcqContextWrapper
	p unsafe.Pointer
}

func NewSapAcqContext() SapAcqContext {
	return SapAcqContext{p: unsafe.Pointer(C.SapAcqContext_New())}
}

func (ctx *SapAcqContext) Delete() {
	C.SapAcqContext_Delete((C.SapAcqContextWrapper)(ctx.p))
}

func (ctx *SapAcqContext) SetID(id int) {
	C.SapAcqContext_SetID((C.SapAcqContextWrapper)(ctx.p), C.int(id))
}

func (ctx *SapAcqContext) GetID() int {
	return int(C.SapAcqContext_GetID((C.SapAcqContextWrapper)(ctx.p)))
}
