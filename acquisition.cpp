#include "sap.h"
#include "acquisition.h"

SapAcquisitionWrapper SapAcquisition_New() {
    return new SapAcquisition();
}

SapAcquisitionWrapper SapAcquisition_NewForLocation(SapLocationWrapper loc, const char* camFilename) {
    return new SapAcquisition(*loc, camFilename);
}

void SapAcquisition_Delete(SapAcquisitionWrapper acq) {
    delete acq;
}

bool SapAcquisition_Create(SapAcquisitionWrapper acq) {
    return acq->Create();
}

bool SapAcquisition_Destroy(SapAcquisitionWrapper acq) {
    return acq->Destroy();
}

bool SapAcquisition_SetParameter(SapAcquisitionWrapper acq, int param, int val) {
    return acq->SetParameter(param, val);
}

bool SapAcquisition_GetParameterInt32(SapAcquisitionWrapper acq, int param, int *val) {
    return acq->GetParameter(param, val);
}

bool SapAcquisition_GetParameterInt64(SapAcquisitionWrapper acq, int param, long *val) {
    return acq->GetParameter(param, val);
}

const char* SapAcquisition_GetLastStatus(SapAcquisitionWrapper acq) {
    return acq->GetLastStatus();
}

bool SapAcquisition_ResetTimeStamp(SapAcquisitionWrapper acq) {
    return acq->ResetTimeStamp();
}
