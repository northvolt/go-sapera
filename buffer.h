#ifndef _SAP_BUFFER_H_
#define _SAP_BUFFER_H_

#include "sap.h"

#ifdef __cplusplus
extern "C" {
#endif

SapBufferWrapper SapBuffer_New();
SapBufferWrapper SapBuffer_NewForAcquisition(int count, SapAcquisitionWrapper src);
SapBufferWrapper SapBuffer_NewWithSize(int count, int width, int height, int format);
void SapBuffer_Delete(SapBufferWrapper buf);
bool SapBuffer_Create(SapBufferWrapper buf);
bool SapBuffer_Destroy(SapBufferWrapper buf);
bool SapBuffer_SetParameter(SapBufferWrapper buf, int param, int val);
bool SapBuffer_GetParameterInt32(SapBufferWrapper buf, int param, int *val);
bool SapBuffer_GetParameterInt64(SapBufferWrapper buf, int param, long *val);
void SapBuffer_SetHeight(SapBufferWrapper buf, int height);
int SapBuffer_GetHeight(SapBufferWrapper buf);
void SapBuffer_SetWidth(SapBufferWrapper buf, int width);
int SapBuffer_GetWidth(SapBufferWrapper buf);
void SapBuffer_SetFormat(SapBufferWrapper buf, int format);
void SapBuffer_SetPixelDepth(SapBufferWrapper buf, int depth);
int SapBuffer_GetPixelDepth(SapBufferWrapper buf);
void SapBuffer_SetSrcNode(SapBufferWrapper buf, SapAcquisitionWrapper acq);
void SapBuffer_ClearSrcNode(SapBufferWrapper buf);
bool SapBuffer_ReadLine(SapBufferWrapper buf, int x1, int y1, int x2, int y2, void* data, int* count);
bool SapBuffer_ReadLineWithIndex(SapBufferWrapper buf, int index, int x1, int y1, int x2, int y2, void* data, int* count);
bool SapBuffer_Copy(SapBufferWrapper buf, SapBufferWrapper pSrc);
bool SapBuffer_CopyWithIndex(SapBufferWrapper buf, SapBufferWrapper pSrc, int srcIndex, int dstIndex);
bool SapBuffer_CopyRect(SapBufferWrapper buf, SapBufferWrapper pSrc, int srcIndex, int xSrc, int ySrc, int width, int height, int dstIndex, int xDest, int yDest);
bool SapBuffer_CopyAll(SapBufferWrapper buf, SapBufferWrapper pSrc);
bool SapBuffer_Save(SapBufferWrapper buf, const char* fileName, const char* options);
bool SapBuffer_SaveForIndex(SapBufferWrapper buf, const char* fileName, const char* options, int index);
const char* SapBuffer_GetLastStatus(SapBufferWrapper buf);
int SapBuffer_GetIndex(SapBufferWrapper buf);
void SapBuffer_SetIndex(SapBufferWrapper buf, int index);
void SapBuffer_Next(SapBufferWrapper buf);

#ifdef __cplusplus
}
#endif

#endif //_SAP_BUFFER_H_
