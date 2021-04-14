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
}

int SapManager_GetServerCount() {
    return SapManager::GetServerCount();
}


SapLocationWrapper SapLocation_New(const char* acqServerName, int acqDeviceNumber) {
    return new SapLocation(acqServerName, acqDeviceNumber);
}

void SapLocation_Delete(SapLocationWrapper loc) {
    delete loc;
}

CORSTATUS DisplayStatus(char *functionName, CORSTATUS status)
{
	if (status != CORSTATUS_OK) {
		char szId[128], szInfo[128], szLevel[64], szModule[64];
		CorManGetStatusTextEx(status, szId, sizeof(szId), szInfo, sizeof(szInfo), szLevel, sizeof(szLevel), szModule, sizeof(szModule));
		printf("%s in \"%s\" <%s module> => %s (%s)\n", szLevel, functionName, szModule, szId, szInfo);
		return status;
	}

	return status;
}
