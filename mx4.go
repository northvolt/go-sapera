package sapera

/*
#include <stdlib.h>
#include "mx4.h"
#include "buffer.h"
*/
import "C"
import "unsafe"

type MX4Metadata struct {
	// C.MX4MetadataWrapper
	p unsafe.Pointer
}

func GetMX4MetadataFromBuffer(buf SapBuffer, width, height int) MX4Metadata {
	return MX4Metadata{p: unsafe.Pointer(C.GetMX4MetadataFromBuffer((C.SapBufferWrapper)(buf.p), C.int(width), C.int(height)))}
}

func (mta MX4Metadata) Close() {
	C.MX4Metadata_Close((C.MX4MetadataWrapper)(mta.p))
}

func (mta MX4Metadata) FrameCounter() int {
	return int(((C.MX4MetadataWrapper)(mta.p)).frameCounter)
}

func (mta MX4Metadata) LineCount() int {
	return int(((C.MX4MetadataWrapper)(mta.p)).lineCount)
}
