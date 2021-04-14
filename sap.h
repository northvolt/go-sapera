#ifndef _SAP_H_
#define _SAP_H_

#include "stdio.h"
#include <stdbool.h>

#ifdef __cplusplus
#include "SapClassBasic.h"
extern "C" {
#endif

#include "corapi.h"

#ifdef __cplusplus
typedef SapBuffer* SapBufferWrapper;
typedef SapAcquisition* SapAcquisitionWrapper;
typedef SapLocation* SapLocationWrapper;
typedef SapTransfer* SapAcqToBufWrapper;
typedef SapXferCallbackInfo* SapXferCallbackInfoWrapper;
typedef SapXferFrameRateInfo* SapXferFrameRateInfoWrapper;
extern "C" {
  extern void xferCallback(SapXferCallbackInfoWrapper pInfo);
}
#else
typedef void* SapBufferWrapper;
typedef void* SapAcquisitionWrapper;
typedef void* SapLocationWrapper;
typedef void* SapAcqToBufWrapper;
typedef void* SapXferCallbackInfoWrapper;
typedef void* SapXferFrameRateInfoWrapper;
extern  void xferCallback(SapXferCallbackInfoWrapper pInfo);
#endif

typedef struct
{
	int                 id;
	SapBufferWrapper    buffer;
	int                 height;
	int                 width;
	int                 pixeldepth;
	long                counter;
} SapXferContext;
typedef SapXferContext* SapXferContextWrapper;

extern void goxferhandler(SapXferCallbackInfoWrapper pInfo);

extern void _CorW32_EnableKernelEventNotification(void); // New entrypoint that will be part of next SaperaLT SDK release.

void SapManager_Init();
int SapManager_GetServerCount();

SapLocationWrapper SapLocation_New(const char* acqServerName, int acqDeviceNumber);
void SapLocation_Delete(SapLocationWrapper loc);

#ifdef __cplusplus
}
#endif

#endif //_SAP_H_
