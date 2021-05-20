#include "sap.h"

extern "C"
{
	void xfergocallback(SapXferCallbackInfoWrapper pInfo)
	{
		goxferhandler(pInfo);
	}

	void acqgocallback(SapAcqCallbackInfoWrapper pInfo)
	{
		goacqhandler(pInfo);
	}
}

void SapManager_Init() {
	_CorW32_EnableKernelEventNotification();
	SapManager::SetDisplayStatusMode(SapManager::StatusLog);
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
