#include "sap.h"
#include "buffer.h"

// SapBuffer_New creates a new SapBufferWrapper
SapBufferWrapper SapBuffer_New() {
    return new SapBuffer();
}

SapBufferWrapper SapBuffer_NewForAcquisition(int count, SapAcquisitionWrapper src) {
   return new SapBuffer(count, src);
}

SapBufferWrapper SapBuffer_NewWithSize(int count, int width, int height, int format) {
    return new SapBuffer(count, width, height, format, SapBuffer::TypeContiguous);
}

void SapBuffer_Delete(SapBufferWrapper buf){
    delete buf;
}

bool SapBuffer_Create(SapBufferWrapper buf) {
    return buf->Create();
}

bool SapBuffer_Destroy(SapBufferWrapper buf) {
    return buf->Destroy();
}

bool SapBuffer_SetParameter(SapBufferWrapper buf, int param, int val) {
    return buf->SetParameter(param, val);
}

bool SapBuffer_GetParameterInt32(SapBufferWrapper buf, int param, int *val) {
    return buf->GetParameter(param, val);
}

bool SapBuffer_GetParameterInt64(SapBufferWrapper buf, int param, long *val) {
    return buf->GetParameter(param, val);
}

void SapBuffer_SetHeight(SapBufferWrapper buf, int height) {
    buf->SetHeight(height);
}

int SapBuffer_GetHeight(SapBufferWrapper buf) {
    return buf->GetHeight();
}

void SapBuffer_SetWidth(SapBufferWrapper buf, int width) {
    buf->SetWidth(width);
}

int SapBuffer_GetWidth(SapBufferWrapper buf) {
    return buf->GetWidth();
}

void SapBuffer_SetFormat(SapBufferWrapper buf, int format) {
    buf->SetFormat(format);
}

void SapBuffer_SetPixelDepth(SapBufferWrapper buf, int depth) {
    buf->SetPixelDepth(depth);
}

int SapBuffer_GetPixelDepth(SapBufferWrapper buf) {
    return buf->GetPixelDepth();
}

void SapBuffer_SetSrcNode(SapBufferWrapper buf, SapAcquisitionWrapper acq) {
    buf->SetSrcNode(acq);
}

void SapBuffer_ClearSrcNode(SapBufferWrapper buf) {
    buf->SetSrcNode(NULL);
}

bool SapBuffer_ReadLine(SapBufferWrapper buf, int x1, int y1, int x2, int y2, void* data, int* count) {
    return buf->ReadLine(x1, y1, x2, y2, data, count);
}

bool SapBuffer_ReadLineWithIndex(SapBufferWrapper buf, int index, int x1, int y1, int x2, int y2, void* data, int* count) {
    return buf->ReadLine(index, x1, y1, x2, y2, data, count);
}

bool SapBuffer_Copy(SapBufferWrapper buf, SapBufferWrapper pSrc) {
    return buf->Copy(pSrc);
}

bool SapBuffer_CopyWithIndex(SapBufferWrapper buf, SapBufferWrapper pSrc, int srcIndex, int dstIndex) {
    return buf->Copy(pSrc, srcIndex, dstIndex);
}

bool SapBuffer_CopyRect(SapBufferWrapper buf, SapBufferWrapper pSrc, int xSrc, int ySrc, int width, int height, int dstIndex, int xDest, int yDest) {
    return buf->CopyRect(pSrc, xSrc, ySrc, width, height, dstIndex, xDest, yDest);
}

bool SapBuffer_CopyRectWithIndex(SapBufferWrapper buf, SapBufferWrapper pSrc, int srcIndex, int xSrc, int ySrc, int width, int height, int dstIndex, int xDest, int yDest) {
    return buf->CopyRect(pSrc, srcIndex, xSrc, ySrc, width, height, dstIndex, xDest, yDest);
}

bool SapBuffer_CopyAll(SapBufferWrapper buf, SapBufferWrapper pSrc) {
    return buf->CopyAll(pSrc);
}

bool SapBuffer_Save(SapBufferWrapper buf, const char* fileName, const char* options) {
    return buf->Save(fileName, options);
}

bool SapBuffer_SaveForIndex(SapBufferWrapper buf, const char* fileName, const char* options, int index) {
    return buf->Save(fileName, options, index);
}

const char* SapBuffer_GetLastStatus(SapBufferWrapper buf) {
    return buf->GetLastStatus();
}

int SapBuffer_GetIndex(SapBufferWrapper buf) {
    return buf->GetIndex();
}

void SapBuffer_SetIndex(SapBufferWrapper buf, int index) {
    buf->SetIndex(index);
}

void SapBuffer_Next(SapBufferWrapper buf) {
    buf->Next();
}
