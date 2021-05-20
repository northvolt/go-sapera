#include "sap.h"
#include "acquisition.h"

SapAcquisitionWrapper SapAcquisition_New() {
    return new SapAcquisition();
}

SapAcquisitionWrapper SapAcquisition_NewForLocation(SapLocationWrapper loc, const char* camFilename) {
    return new SapAcquisition(*loc, camFilename, EventNone, goacqhandler);
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

bool SapAcquisition_RegisterCallback(SapAcquisitionWrapper acq, UINT64 eventType, SapAcqContextWrapper context) {
    return acq->RegisterCallback(eventType, goacqhandler, NULL);
}

bool SapAcquisition_UnregisterCallbacks(SapAcquisitionWrapper acq) {
    // UINT32 eventCount, eventIndex;
    // acq->GetEventCount(&eventCount);
    // for (eventIndex = 0; eventIndex < eventCount; eventIndex++) {
    //     BOOL isRegistered;
    //     acq.IsCallbackRegistered(eventIndex, &isRegistered);
    //     if (isRegistered) {
    //         acq->UnregisterCallback(eventIndex);
    //     }
    // }

    return true;
}

bool SapAcquisition_SetCallbackInfo(SapAcquisitionWrapper acq) {
    return acq->SetCallbackInfo(goacqhandler);
}

bool SapAcquisition_SetEventType(SapAcquisitionWrapper acq, UINT64 eventType) {
    return acq->SetEventType(eventType);
}
