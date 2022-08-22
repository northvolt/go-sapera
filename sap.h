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
typedef SapAcqCallbackInfo* SapAcqCallbackInfoWrapper;
extern "C" {
  extern void xferCallback(SapXferCallbackInfoWrapper pInfo);
  extern void acqCallback(SapAcqCallbackInfoWrapper pInfo);
}
typedef SapManVersionInfo* SapVersionInfoWrapper;
#else
typedef void* SapBufferWrapper;
typedef void* SapAcquisitionWrapper;
typedef void* SapLocationWrapper;
typedef void* SapAcqToBufWrapper;
typedef void* SapXferCallbackInfoWrapper;
typedef void* SapXferFrameRateInfoWrapper;
typedef void* SapAcqCallbackInfoWrapper;
extern  void xferCallback(SapXferCallbackInfoWrapper pInfo);
extern  void acqCallback(SapAcqCallbackInfoWrapper pInfo);
typedef void* SapVersionInfoWrapper;
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

typedef struct
{
	int                 id;
} SapAcqContext;
typedef SapAcqContext* SapAcqContextWrapper;

extern void goxferhandler(SapXferCallbackInfoWrapper pInfo);
extern void goacqhandler(SapAcqCallbackInfoWrapper pInfo);

extern void _CorW32_EnableKernelEventNotification(void); // New entrypoint that will be part of next SaperaLT SDK release.

void SapManager_Init();
int SapManager_GetServerCount();
const char* SapManager_GetLastStatus();
SapLocationWrapper SapLocation_New(const char* acqServerName, int acqDeviceNumber);
void SapLocation_Delete(SapLocationWrapper loc);

SapVersionInfoWrapper SapManager_GetVersionInfo();
void SapVersionInfo_Delete(SapVersionInfoWrapper v);
int SapVersionInfo_GetMajor(SapVersionInfoWrapper v);
int SapVersionInfo_GetMinor(SapVersionInfoWrapper v);
int SapVersionInfo_GetRevision(SapVersionInfoWrapper v);
int SapVersionInfo_GetBuild(SapVersionInfoWrapper v);

#ifdef __cplusplus
}
#endif

#endif //_SAP_H_
