#include "context.h"

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

SapAcqContextWrapper SapAcqContext_New() {
    return new SapAcqContext;
}

void SapAcqContext_Delete(SapAcqContextWrapper ctx) {
    delete ctx;
}

void SapAcqContext_SetID(SapAcqContextWrapper ctx, int id) {
    ctx->id = id;
    return; 
}

int SapAcqContext_GetID(SapAcqContextWrapper ctx) {
    return ctx->id;
}
