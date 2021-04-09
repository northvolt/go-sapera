#include "sap.h"
#include "transfer.h"

SapAcqToBufWrapper SapAcqToBuf_New(SapAcquisitionWrapper acq, SapBufferWrapper buf, SapXferContextWrapper context) {
	return new SapAcqToBuf(acq, buf, xfergocallback, context);
}

void SapAcqToBuf_Delete(SapAcqToBufWrapper trs) {
    delete trs;
}

bool SapAcqToBuf_Create(SapAcqToBufWrapper trs) {
    return trs->Create();
}

bool SapAcqToBuf_Destroy(SapAcqToBufWrapper trs) {
    return trs->Destroy();
}

bool SapAcqToBuf_Grab(SapAcqToBufWrapper trs) {
    return trs->Grab();
}

bool SapAcqToBuf_Snap(SapAcqToBufWrapper trs) {
    return trs->Snap(16);
}

bool SapAcqToBuf_IsGrabbing(SapAcqToBufWrapper trs) {
    return trs->IsGrabbing();
}

bool SapAcqToBuf_Freeze(SapAcqToBufWrapper trs) {
    return trs->Freeze();
}

bool SapAcqToBuf_Abort(SapAcqToBufWrapper trs) {
    return trs->Abort();
}

bool SapAcqToBuf_Wait(SapAcqToBufWrapper trs, int ms) {
    return trs->Wait(ms);
}

const char* SapAcqToBuf_GetLastStatus(SapAcqToBufWrapper trs) {
    return trs->GetLastStatus();
}

bool SapAcqToBuf_SetCallbackInfo(SapAcqToBufWrapper trs, SapXferContextWrapper context) {
    return trs->SetCallbackInfo(xferCallback, context);
}

SapXferContextWrapper SapXferContext_New() {
    return new SapXferContext;
}

void SapXferContext_Delete(SapXferContextWrapper ctx) {
    delete ctx;
}

void SapXferContext_SetID(SapXferContextWrapper ctx, int id) {
    ctx->id = id;
    return; 
}

int SapXferContext_GetID(SapXferContextWrapper ctx) {
    return ctx->id;
}

void SapXferContext_SetBuffer(SapXferContextWrapper ctx, SapBufferWrapper buf) {
    ctx->buffer = buf;
    return;
}

SapBufferWrapper SapXferContext_GetBuffer(SapXferContextWrapper ctx) {
    return ctx->buffer;
}

void SapXferContext_SetHeight(SapXferContextWrapper ctx, int height) {
    ctx->height = height;
    return; 
}

int SapXferContext_GetHeight(SapXferContextWrapper ctx) {
    return ctx->height;
}

void SapXferContext_SetWidth(SapXferContextWrapper ctx, int width) {
    ctx->width = width;
    return; 
}

int SapXferContext_GetWidth(SapXferContextWrapper ctx) {
    return ctx->width;
}

void SapXferContext_SetPixelDepth(SapXferContextWrapper ctx, int depth) {
    ctx->pixeldepth = depth;
    return; 
}

int SapXferContext_GetPixelDepth(SapXferContextWrapper ctx) {
    return ctx->pixeldepth;
}

void SapXferContext_SetCounter(SapXferContextWrapper ctx, long count) {
    ctx->counter = count;
    return; 
}

long SapXferContext_GetCounter(SapXferContextWrapper ctx) {
    return ctx->counter;
}

SapXferContextWrapper SapXferCallbackInfo_GetContext(SapXferCallbackInfoWrapper cbinfo) {
    return (SapXferContextWrapper)cbinfo->GetContext();
}

SapXferFrameRateInfoWrapper SapAcqToBuf_GetFrameRateStatistics(SapAcqToBufWrapper trs) {
    return trs->GetFrameRateStatistics();
}

int SapXferFrameRateInfo_GetBufferFrameRate(SapXferFrameRateInfoWrapper inf) {
    return inf->GetBufferFrameRate();
}
