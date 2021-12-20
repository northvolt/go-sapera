#ifndef _MX4_H_
#define _MX4_H_

#include "sap.h"
#include "buffer.h"

#ifdef __cplusplus
extern "C" {
#endif

#define MX4MetadataSize 64

typedef struct
{
  ULONGLONG shaftEncoderCount;
  ULONGLONG lineCount;
  ULONGLONG lineTriggerCount;
  ULONGLONG timeStamp;
  ULONG frameCounter;
  UCHAR generalInputs;
  UCHAR generalOutputs;
  UCHAR biDirectionalIOs;
  UCHAR reserved[25];
} MX4Metadata;

typedef MX4Metadata* MX4MetadataWrapper;

MX4MetadataWrapper GetMX4MetadataFromBuffer(SapBufferWrapper buf, int width, int line);
void MX4Metadata_Close(MX4MetadataWrapper mta);

#ifdef __cplusplus
}
#endif

#endif //_MX4_H_
