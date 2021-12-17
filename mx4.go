package sapera

/*
#include <stdlib.h>
#include "mx4.h"
#include "buffer.h"
*/
import "C"
import (
	"time"
	"unsafe"
)

// MX4Metadata is the metadata such as the line count and encoder count for each
// captured frame.
type MX4Metadata struct {
	// C.MX4MetadataWrapper
	p unsafe.Pointer
}

// GetMX4MetadataFromBuffer returns the MX4Metadata for a buffer. Pass in the width of the
// image part of the buffer, and the line number (1-indexed) for which you want the metadata.
func GetMX4MetadataFromBuffer(buf SapBuffer, width, line int) MX4Metadata {
	return MX4Metadata{p: unsafe.Pointer(C.GetMX4MetadataFromBuffer((C.SapBufferWrapper)(buf.p), C.int(width), C.int(line)))}
}

// Close cleans up any memory that has been allocated for the MX4Metadata.
func (mta MX4Metadata) Close() {
	C.MX4Metadata_Close((C.MX4MetadataWrapper)(mta.p))
}

func (mta MX4Metadata) ShaftEncoderCount() int {
	return int(((C.MX4MetadataWrapper)(mta.p)).shaftEncoderCount)
}

func (mta MX4Metadata) LineCount() int64 {
	return int64(((C.MX4MetadataWrapper)(mta.p)).lineCount)
}

func (mta MX4Metadata) LineTriggerCount() int64 {
	return int64(((C.MX4MetadataWrapper)(mta.p)).lineTriggerCount)
}

func (mta MX4Metadata) Timestamp() time.Time {
	return time.Unix(0, int64(((C.MX4MetadataWrapper)(mta.p)).timeStamp))
}

func (mta MX4Metadata) FrameCounter() int {
	return int(((C.MX4MetadataWrapper)(mta.p)).frameCounter)
}

func (mta MX4Metadata) GeneralInputs() byte {
	return byte(((C.MX4MetadataWrapper)(mta.p)).generalInputs)
}

func (mta MX4Metadata) GeneralOutputs() byte {
	return byte(((C.MX4MetadataWrapper)(mta.p)).generalOutputs)
}

func (mta MX4Metadata) BiDirectionalIOs() byte {
	return byte(((C.MX4MetadataWrapper)(mta.p)).biDirectionalIOs)
}
