#ifndef _SAP_ACQUISITION_H_
#define _SAP_ACQUISITION_H_

#include "sap.h"

#ifdef __cplusplus
extern "C" {
#endif

SapAcquisitionWrapper SapAcquisition_New();
SapAcquisitionWrapper SapAcquisition_NewForLocation(SapLocationWrapper loc, const char* camFilename);
void SapAcquisition_Delete(SapAcquisitionWrapper acq);
bool SapAcquisition_Create(SapAcquisitionWrapper acq);
bool SapAcquisition_Destroy(SapAcquisitionWrapper acq);
bool SapAcquisition_SetParameter(SapAcquisitionWrapper acq, int param, int val);
bool SapAcquisition_GetParameterInt32(SapAcquisitionWrapper acq, int param, int *val);
bool SapAcquisition_GetParameterInt64(SapAcquisitionWrapper acq, int param, long *val);
const char* SapAcquisition_GetLastStatus(SapAcquisitionWrapper acq);
bool SapAcquisition_ResetTimeStamp(SapAcquisitionWrapper acq);

#ifdef __cplusplus
}
#endif

#endif //_SAP_ACQUISITION_H_
