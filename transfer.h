#ifndef _SAP_TRANSFER_H_
#define _SAP_TRANSFER_H_

#include "sap.h"

#ifdef __cplusplus
extern "C" {
#endif

SapAcqToBufWrapper SapAcqToBuf_New(SapAcquisitionWrapper acq, SapBufferWrapper buf, SapXferContextWrapper context);
void SapAcqToBuf_Delete(SapAcqToBufWrapper trs);
bool SapAcqToBuf_Create(SapAcqToBufWrapper trs);
bool SapAcqToBuf_Destroy(SapAcqToBufWrapper trs);
bool SapAcqToBuf_Snap(SapAcqToBufWrapper trs);
bool SapAcqToBuf_Grab(SapAcqToBufWrapper trs);
bool SapAcqToBuf_IsGrabbing(SapAcqToBufWrapper trs);
bool SapAcqToBuf_Freeze(SapAcqToBufWrapper trs);
bool SapAcqToBuf_Abort(SapAcqToBufWrapper trs);
bool SapAcqToBuf_Wait(SapAcqToBufWrapper trs, int ms);
const char* SapAcqToBuf_GetLastStatus(SapAcqToBufWrapper trs);
bool SapAcqToBuf_SetCallbackInfo(SapAcqToBufWrapper trs, SapXferContextWrapper context);

SapXferFrameRateInfoWrapper SapAcqToBuf_GetFrameRateStatistics(SapAcqToBufWrapper trs);
int SapXferFrameRateInfo_GetBufferFrameRate(SapXferFrameRateInfoWrapper inf);

#ifdef __cplusplus
}
#endif

#endif //_SAP_TRANSFER_H_
