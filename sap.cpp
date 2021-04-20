#include "sap.h"

extern "C"
{
	void xfergocallback(SapXferCallbackInfoWrapper pInfo)
	{
		goxferhandler(pInfo);
	}
}

void SapManager_Init() {
	_CorW32_EnableKernelEventNotification();
	SapManager::SetDisplayStatusMode(SapManager::StatusCustom);
}

int SapManager_GetServerCount() {
    return SapManager::GetServerCount();
}

const char* SapManager_GetLastStatus() {
    return SapManager::GetLastStatus();
}

SapLocationWrapper SapLocation_New(const char* acqServerName, int acqDeviceNumber) {
    return new SapLocation(acqServerName, acqDeviceNumber);
}

void SapLocation_Delete(SapLocationWrapper loc) {
    delete loc;
}
