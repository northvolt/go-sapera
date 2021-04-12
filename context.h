#ifndef _SAP_CONTEXT_H_
#define _SAP_CONTEXT_H_

#include "sap.h"

#ifdef __cplusplus
extern "C" {
#endif

SapXferContextWrapper SapXferContext_New();
void SapXferContext_Delete(SapXferContextWrapper ctx);
void SapXferContext_SetID(SapXferContextWrapper ctx, int id);
int SapXferContext_GetID(SapXferContextWrapper ctx);
void SapXferContext_SetBuffer(SapXferContextWrapper ctx, SapBufferWrapper buf);
SapBufferWrapper SapXferContext_GetBuffer(SapXferContextWrapper ctx);
void SapXferContext_SetHeight(SapXferContextWrapper ctx, int height);
int SapXferContext_GetHeight(SapXferContextWrapper ctx);
void SapXferContext_SetWidth(SapXferContextWrapper ctx, int width);
int SapXferContext_GetWidth(SapXferContextWrapper ctx);
void SapXferContext_SetPixelDepth(SapXferContextWrapper ctx, int depth);
int SapXferContext_GetPixelDepth(SapXferContextWrapper ctx);
void SapXferContext_SetCounter(SapXferContextWrapper ctx, long count);
long SapXferContext_GetCounter(SapXferContextWrapper ctx);

SapXferContextWrapper SapXferCallbackInfo_GetContext(SapXferCallbackInfoWrapper cbinfo);

#ifdef __cplusplus
}
#endif

#endif //_SAP_CONTEXT_H_
