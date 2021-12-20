#include "sap.h"
#include "mx4.h"

// GetMX4MetadataFromBuffer returns a struct of MX4Metadata that has been parsed from the buffer
// of data passed into this function. Also pass in the width of the image area of the buffer, along with the
// line number (1-indexed) of the buffer for which you want the metadata.
MX4MetadataWrapper GetMX4MetadataFromBuffer(SapBufferWrapper buf, int width, int line) {
	int numRead;
	MX4MetadataWrapper metadata = new MX4Metadata;
	buf->ReadLine(width, line-1, width-1+MX4MetadataSize, line-1, metadata, &numRead);
	return metadata;
}

void MX4Metadata_Close(MX4MetadataWrapper mta) {
	delete mta;
}
