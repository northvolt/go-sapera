#ifndef _MX4_H_
#define _MX4_H_

#include "stdio.h"
#include <stdbool.h>

#ifdef __cplusplus
#include "SapClassBasic.h"
extern "C" {
#endif

#include "corapi.h"

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
} MX4Metadeta;

typedef MX4Metadeta* MX4MetadetaWrapper;

#ifdef __cplusplus
}
#endif

#endif //_MX4_H_
