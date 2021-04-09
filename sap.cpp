#include "sap.h"

extern "C" {
  void xferCallback(SapXferCallbackInfoWrapper pInfo) {
      printf("xferCallback called\n");
  }
}

extern "C"
{
	void xfergocallback(SapXferCallbackInfoWrapper pInfo)
	{
		goxferhandler(pInfo);
	}

	void SetGoCallbackInfo(SapAcqToBufWrapper xfer)
	{
		xfer->SetCallbackInfo(xfergocallback, NULL);
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
